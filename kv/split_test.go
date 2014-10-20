// Copyright 2014 The Cockroach Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License. See the AUTHORS file
// for names of contributors.
//
// Author: Spencer Kimball (spencer.kimball@gmail.com)

package kv

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/cockroachdb/cockroach/proto"
	"github.com/cockroachdb/cockroach/storage"
	"github.com/cockroachdb/cockroach/storage/engine"
	"github.com/cockroachdb/cockroach/util"
	"github.com/cockroachdb/cockroach/util/log"
)

// startTestWriter creates a writer which intiates a sequence of
// transactions, each which writes up to 10 times to random keys with
// random values. If not nil, txnChannel is written to every time a
// new transaction starts.
func startTestWriter(db storage.DB, i int64, valBytes int32, wg *sync.WaitGroup, retries *int32,
	txnChannel chan struct{}, done <-chan struct{}, t *testing.T) {
	src := rand.New(rand.NewSource(i))
	for {
		select {
		case <-done:
			if wg != nil {
				wg.Done()
			}
			return
		default:
			txnOpts := &storage.TransactionOptions{
				Name: fmt.Sprintf("concurrent test %d", i),
				Retry: &util.RetryOptions{
					Backoff:    1 * time.Millisecond,
					MaxBackoff: 10 * time.Millisecond,
					Constant:   2,
				},
			}
			first := true
			err := db.RunTransaction(txnOpts, func(txn storage.DB) error {
				if first && txnChannel != nil {
					txnChannel <- struct{}{}
				} else if !first && retries != nil {
					atomic.AddInt32(retries, 1)
				}
				first = false
				for j := 0; j <= int(src.Int31n(10)); j++ {
					key := []byte(util.RandString(src, 10))
					val := []byte(util.RandString(src, int(src.Int31n(valBytes))))
					putR := <-txn.Put(&proto.PutRequest{RequestHeader: proto.RequestHeader{Key: key}, Value: proto.Value{Bytes: val}})
					if putR.GoError() != nil {
						log.Infof("experienced an error in routine %d: %s", i, putR.GoError())
						return putR.GoError()
					}
				}
				return nil
			})
			if err != nil {
				t.Error(err)
			} else {
				time.Sleep(1 * time.Millisecond)
			}
		}
	}
}

// TestRangeSplitsWithConcurrentTxns does 5 consecutive splits while
// 10 concurrent goroutines, each running successive transactions
// composed of a random mix of puts.
func TestRangeSplitsWithConcurrentTxns(t *testing.T) {
	db, _, _ := createTestDB(t)
	defer db.Close()

	// This channel shuts the whole apparatus down.
	done := make(chan struct{})
	txnChannel := make(chan struct{}, 1000)

	// Set five split keys, about evenly spaced along the range of random keys.
	splitKeys := []engine.Key{engine.Key("G"), engine.Key("R"), engine.Key("a"), engine.Key("l"), engine.Key("s")}

	// Start up the concurrent goroutines which run transactions.
	const concurrency = 10
	var retries int32
	var wg sync.WaitGroup
	wg.Add(concurrency)
	for i := 0; i < concurrency; i++ {
		go startTestWriter(db, int64(i), 1<<7, &wg, &retries, txnChannel, done, t)
	}

	// Execute the consecutive splits.
	for _, splitKey := range splitKeys {
		// Allow txns to start before initiating split.
		for i := 0; i < concurrency; i++ {
			<-txnChannel
		}
		log.Infof("starting split at key %q..", splitKey)
		splitR := <-db.AdminSplit(&proto.AdminSplitRequest{RequestHeader: proto.RequestHeader{Key: splitKey}, SplitKey: splitKey})
		if splitR.GoError() != nil {
			t.Fatal(splitR.GoError())
		}
		log.Infof("split at key %q complete", splitKey)
	}

	close(done)
	wg.Wait()

	if retries != 0 {
		t.Errorf("expected no retries splitting a range with concurrent writes, "+
			"as range splits do not cause conflicts; got %d", retries)
	}
}

// TestRangeSplitsWithWritePressure sets the zone config max bytes for
// a range to 256K and writes data until there are five ranges.
func TestRangeSplitsWithWritePressure(t *testing.T) {
	db, _, _ := createTestDB(t)
	defer db.Close()
	txnOpts := &storage.TransactionOptions{
		Name: "scan meta2 records",
		Retry: &util.RetryOptions{
			Backoff:    1 * time.Millisecond,
			MaxBackoff: 10 * time.Millisecond,
			Constant:   2,
		},
	}

	// Rewrite a zone config with low max bytes.
	zoneConfig := &proto.ZoneConfig{
		ReplicaAttrs: []proto.Attributes{
			proto.Attributes{},
			proto.Attributes{},
			proto.Attributes{},
		},
		RangeMinBytes: 1 << 8,
		RangeMaxBytes: 1 << 18,
	}
	if err := storage.PutProto(db, engine.MakeKey(engine.KeyConfigZonePrefix, engine.KeyMin), zoneConfig); err != nil {
		t.Fatal(err)
	}

	// Start test writer write about a 32K/key so there aren't too many writes necessary to split 64K range.
	done := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(1)
	go startTestWriter(db, int64(0), 1<<15, &wg, nil, nil, done, t)

	// Check that we split 5 times in allotted time.
	if err := util.IsTrueWithin(func() bool {
		// Scan the txn records (in a txn due to possible retries) to see number of ranges.
		var kvs []proto.KeyValue
		if err := db.RunTransaction(txnOpts, func(txn storage.DB) error {
			scanR := <-txn.Scan(&proto.ScanRequest{
				RequestHeader: proto.RequestHeader{
					Key:    engine.KeyMeta2Prefix,
					EndKey: engine.KeyMetaMax,
				},
			})
			if scanR.GoError() != nil {
				return scanR.GoError()
			}
			kvs = scanR.Rows
			return nil
		}); err != nil {
			t.Fatalf("failed to scan meta1 keys: %s", err)
		}
		return len(kvs) >= 5
	}, 2*time.Second); err != nil {
		t.Errorf("failed to split 5 times: %s", err)
	}
	close(done)
	wg.Wait()
}
