// Copyright 2021 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package streaming

import (
	"github.com/cockroachdb/cockroach/pkg/kv"
	"github.com/cockroachdb/cockroach/pkg/sql/sem/tree"
	"github.com/cockroachdb/cockroach/pkg/util/hlc"
	"github.com/cockroachdb/errors"
)

// StreamID is the ID of a replication stream.
type StreamID int64

// InvalidStreamID is the zero value for StreamID corresponding to no stream.
const InvalidStreamID StreamID = 0

// GetReplicationStreamManagerHook is the hook to get a collection of APIs that streaming replication supports.
var GetReplicationStreamManagerHook func() (ReplicationStreamManager, error)

// ReplicationStreamManager represents a collection of APIs that streaming replication supports.
type ReplicationStreamManager interface {
	// CompleteStreamIngestion signals a running stream ingestion job to complete.
	CompleteStreamIngestion(evalCtx *tree.EvalContext, txn *kv.Txn, jobID int, cutoverTimestamp hlc.Timestamp) error

	// StartReplicationStream starts a stream replication job for the specified tenant on the producer side.
	StartReplicationStream(evalCtx *tree.EvalContext, txn *kv.Txn, tenantID uint64) (StreamID, error)
}

// GetReplicationStreamManager returns a ReplicationStreamManager if a CCL binary is loaded.
func GetReplicationStreamManager() (ReplicationStreamManager, error) {
	if GetReplicationStreamManagerHook == nil {
		return nil, errors.New("replication streaming requires a CCL binary")
	}
	return GetReplicationStreamManagerHook()
}
