// Copyright 2019 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package colexec

import (
	"context"

	"github.com/cockroachdb/cockroach/pkg/col/coldata"
	"github.com/cockroachdb/cockroach/pkg/sql/colexec/colexecargs"
	"github.com/cockroachdb/cockroach/pkg/sql/colexecop"
	"github.com/cockroachdb/cockroach/pkg/sql/execinfra"
	"github.com/cockroachdb/cockroach/pkg/sql/execinfrapb"
	"github.com/cockroachdb/cockroach/pkg/util/tracing"
)

// SerialUnorderedSynchronizer is an Operator that combines multiple Operator
// streams into one. It reads its inputs one by one until each one is exhausted,
// at which point it moves to the next input. See ParallelUnorderedSynchronizer
// for a parallel implementation. The serial one is used when concurrency is
// undesirable - for example when the whole query is planned on the gateway and
// we want to run it in the RootTxn.
type SerialUnorderedSynchronizer struct {
	ctx  context.Context
	span *tracing.Span

	inputs []colexecargs.OpWithMetaInfo
	// curSerialInputIdx indicates the index of the current input being consumed.
	curSerialInputIdx int
}

var (
	_ colexecop.Operator = &SerialUnorderedSynchronizer{}
	_ execinfra.OpNode   = &SerialUnorderedSynchronizer{}
	_ colexecop.Closer   = &SerialUnorderedSynchronizer{}
)

// ChildCount implements the execinfra.OpNode interface.
func (s *SerialUnorderedSynchronizer) ChildCount(verbose bool) int {
	return len(s.inputs)
}

// Child implements the execinfra.OpNode interface.
func (s *SerialUnorderedSynchronizer) Child(nth int, verbose bool) execinfra.OpNode {
	return s.inputs[nth].Root
}

// NewSerialUnorderedSynchronizer creates a new SerialUnorderedSynchronizer.
func NewSerialUnorderedSynchronizer(
	inputs []colexecargs.OpWithMetaInfo,
) *SerialUnorderedSynchronizer {
	return &SerialUnorderedSynchronizer{
		inputs: inputs,
	}
}

// Init is part of the Operator interface.
func (s *SerialUnorderedSynchronizer) Init() {
	for _, input := range s.inputs {
		input.Root.Init()
	}
}

// maybeStartTracingSpan stores the context and possibly starts a tracing span
// on its first call and is a noop on all consequent calls.
// TODO(yuzefovich): remove this once ctx is passed in Init.
func (s *SerialUnorderedSynchronizer) maybeStartTracingSpan(ctx context.Context) {
	if s.ctx == nil {
		// It is the very first call to maybeStartTracingSpan.
		s.ctx, s.span = execinfra.ProcessorSpan(ctx, "serial unordered sync")
	}
}

// Next is part of the Operator interface.
func (s *SerialUnorderedSynchronizer) Next(ctx context.Context) coldata.Batch {
	s.maybeStartTracingSpan(ctx)
	for {
		if s.curSerialInputIdx == len(s.inputs) {
			return coldata.ZeroBatch
		}
		b := s.inputs[s.curSerialInputIdx].Root.Next(s.ctx)
		if b.Length() == 0 {
			s.curSerialInputIdx++
		} else {
			return b
		}
	}
}

// DrainMeta is part of the colexecop.MetadataSource interface.
func (s *SerialUnorderedSynchronizer) DrainMeta(
	ctx context.Context,
) []execinfrapb.ProducerMetadata {
	var bufferedMeta []execinfrapb.ProducerMetadata
	// It is possible that Next was never called, yet the tracing is enabled.
	s.maybeStartTracingSpan(ctx)
	if s.span != nil {
		for i := range s.inputs {
			for _, stats := range s.inputs[i].StatsCollectors {
				s.span.RecordStructured(stats.GetStats())
			}
		}
		if meta := execinfra.GetTraceDataAsMetadata(s.span); meta != nil {
			bufferedMeta = append(bufferedMeta, *meta)
		}
	}
	for _, input := range s.inputs {
		bufferedMeta = append(bufferedMeta, input.MetadataSources.DrainMeta(s.ctx)...)
	}
	return bufferedMeta
}

// Close is part of the Closer interface.
func (s *SerialUnorderedSynchronizer) Close(ctx context.Context) error {
	for _, input := range s.inputs {
		input.ToClose.CloseAndLogOnErr(ctx, "serial unordered synchronizer")
	}
	if s.span != nil {
		s.span.Finish()
	}
	return nil
}
