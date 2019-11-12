// Code generated by gen_batch.go; DO NOT EDIT.
// GENERATED FILE DO NOT EDIT

package roachpb

import (
	"fmt"
	"strconv"
	"strings"
)

// GetInner returns the error contained in the union.
func (ru ErrorDetail) GetInner() error {
	switch t := ru.GetValue().(type) {
	case *ErrorDetail_NotLeaseHolder:
		return t.NotLeaseHolder
	case *ErrorDetail_RangeNotFound:
		return t.RangeNotFound
	case *ErrorDetail_RangeKeyMismatch:
		return t.RangeKeyMismatch
	case *ErrorDetail_ReadWithinUncertaintyInterval:
		return t.ReadWithinUncertaintyInterval
	case *ErrorDetail_TransactionAborted:
		return t.TransactionAborted
	case *ErrorDetail_TransactionPush:
		return t.TransactionPush
	case *ErrorDetail_TransactionRetry:
		return t.TransactionRetry
	case *ErrorDetail_TransactionStatus:
		return t.TransactionStatus
	case *ErrorDetail_WriteIntent:
		return t.WriteIntent
	case *ErrorDetail_WriteTooOld:
		return t.WriteTooOld
	case *ErrorDetail_OpRequiresTxn:
		return t.OpRequiresTxn
	case *ErrorDetail_ConditionFailed:
		return t.ConditionFailed
	case *ErrorDetail_LeaseRejected:
		return t.LeaseRejected
	case *ErrorDetail_NodeUnavailable:
		return t.NodeUnavailable
	case *ErrorDetail_Send:
		return t.Send
	case *ErrorDetail_RaftGroupDeleted:
		return t.RaftGroupDeleted
	case *ErrorDetail_ReplicaCorruption:
		return t.ReplicaCorruption
	case *ErrorDetail_ReplicaTooOld:
		return t.ReplicaTooOld
	case *ErrorDetail_AmbiguousResult:
		return t.AmbiguousResult
	case *ErrorDetail_StoreNotFound:
		return t.StoreNotFound
	case *ErrorDetail_TransactionRetryWithProtoRefresh:
		return t.TransactionRetryWithProtoRefresh
	case *ErrorDetail_IntegerOverflow:
		return t.IntegerOverflow
	case *ErrorDetail_UnsupportedRequest:
		return t.UnsupportedRequest
	case *ErrorDetail_TimestampBefore:
		return t.TimestampBefore
	case *ErrorDetail_TxnAlreadyEncounteredError:
		return t.TxnAlreadyEncounteredError
	case *ErrorDetail_IntentMissing:
		return t.IntentMissing
	case *ErrorDetail_MergeInProgress:
		return t.MergeInProgress
	case *ErrorDetail_RangefeedRetry:
		return t.RangefeedRetry
	case *ErrorDetail_IndeterminateCommit:
		return t.IndeterminateCommit
	default:
		return nil
	}
}

// GetInner returns the Request contained in the union.
func (ru RequestUnion) GetInner() Request {
	switch t := ru.GetValue().(type) {
	case *RequestUnion_Get:
		return t.Get
	case *RequestUnion_Put:
		return t.Put
	case *RequestUnion_ConditionalPut:
		return t.ConditionalPut
	case *RequestUnion_Increment:
		return t.Increment
	case *RequestUnion_Delete:
		return t.Delete
	case *RequestUnion_DeleteRange:
		return t.DeleteRange
	case *RequestUnion_ClearRange:
		return t.ClearRange
	case *RequestUnion_RevertRange:
		return t.RevertRange
	case *RequestUnion_Scan:
		return t.Scan
	case *RequestUnion_BeginTransaction:
		return t.BeginTransaction
	case *RequestUnion_EndTransaction:
		return t.EndTransaction
	case *RequestUnion_AdminSplit:
		return t.AdminSplit
	case *RequestUnion_AdminUnsplit:
		return t.AdminUnsplit
	case *RequestUnion_AdminMerge:
		return t.AdminMerge
	case *RequestUnion_AdminTransferLease:
		return t.AdminTransferLease
	case *RequestUnion_AdminChangeReplicas:
		return t.AdminChangeReplicas
	case *RequestUnion_AdminRelocateRange:
		return t.AdminRelocateRange
	case *RequestUnion_HeartbeatTxn:
		return t.HeartbeatTxn
	case *RequestUnion_Gc:
		return t.Gc
	case *RequestUnion_PushTxn:
		return t.PushTxn
	case *RequestUnion_RecoverTxn:
		return t.RecoverTxn
	case *RequestUnion_ResolveIntent:
		return t.ResolveIntent
	case *RequestUnion_ResolveIntentRange:
		return t.ResolveIntentRange
	case *RequestUnion_Merge:
		return t.Merge
	case *RequestUnion_TruncateLog:
		return t.TruncateLog
	case *RequestUnion_RequestLease:
		return t.RequestLease
	case *RequestUnion_ReverseScan:
		return t.ReverseScan
	case *RequestUnion_ComputeChecksum:
		return t.ComputeChecksum
	case *RequestUnion_CheckConsistency:
		return t.CheckConsistency
	case *RequestUnion_InitPut:
		return t.InitPut
	case *RequestUnion_TransferLease:
		return t.TransferLease
	case *RequestUnion_LeaseInfo:
		return t.LeaseInfo
	case *RequestUnion_WriteBatch:
		return t.WriteBatch
	case *RequestUnion_Export:
		return t.Export
	case *RequestUnion_Import:
		return t.Import
	case *RequestUnion_QueryTxn:
		return t.QueryTxn
	case *RequestUnion_QueryIntent:
		return t.QueryIntent
	case *RequestUnion_AdminScatter:
		return t.AdminScatter
	case *RequestUnion_AddSstable:
		return t.AddSstable
	case *RequestUnion_RecomputeStats:
		return t.RecomputeStats
	case *RequestUnion_Refresh:
		return t.Refresh
	case *RequestUnion_RefreshRange:
		return t.RefreshRange
	case *RequestUnion_Subsume:
		return t.Subsume
	case *RequestUnion_RangeStats:
		return t.RangeStats
	default:
		return nil
	}
}

// GetInner returns the Response contained in the union.
func (ru ResponseUnion) GetInner() Response {
	switch t := ru.GetValue().(type) {
	case *ResponseUnion_Get:
		return t.Get
	case *ResponseUnion_Put:
		return t.Put
	case *ResponseUnion_ConditionalPut:
		return t.ConditionalPut
	case *ResponseUnion_Increment:
		return t.Increment
	case *ResponseUnion_Delete:
		return t.Delete
	case *ResponseUnion_DeleteRange:
		return t.DeleteRange
	case *ResponseUnion_ClearRange:
		return t.ClearRange
	case *ResponseUnion_RevertRange:
		return t.RevertRange
	case *ResponseUnion_Scan:
		return t.Scan
	case *ResponseUnion_BeginTransaction:
		return t.BeginTransaction
	case *ResponseUnion_EndTransaction:
		return t.EndTransaction
	case *ResponseUnion_AdminSplit:
		return t.AdminSplit
	case *ResponseUnion_AdminUnsplit:
		return t.AdminUnsplit
	case *ResponseUnion_AdminMerge:
		return t.AdminMerge
	case *ResponseUnion_AdminTransferLease:
		return t.AdminTransferLease
	case *ResponseUnion_AdminChangeReplicas:
		return t.AdminChangeReplicas
	case *ResponseUnion_AdminRelocateRange:
		return t.AdminRelocateRange
	case *ResponseUnion_HeartbeatTxn:
		return t.HeartbeatTxn
	case *ResponseUnion_Gc:
		return t.Gc
	case *ResponseUnion_PushTxn:
		return t.PushTxn
	case *ResponseUnion_RecoverTxn:
		return t.RecoverTxn
	case *ResponseUnion_ResolveIntent:
		return t.ResolveIntent
	case *ResponseUnion_ResolveIntentRange:
		return t.ResolveIntentRange
	case *ResponseUnion_Merge:
		return t.Merge
	case *ResponseUnion_TruncateLog:
		return t.TruncateLog
	case *ResponseUnion_RequestLease:
		return t.RequestLease
	case *ResponseUnion_ReverseScan:
		return t.ReverseScan
	case *ResponseUnion_ComputeChecksum:
		return t.ComputeChecksum
	case *ResponseUnion_CheckConsistency:
		return t.CheckConsistency
	case *ResponseUnion_InitPut:
		return t.InitPut
	case *ResponseUnion_LeaseInfo:
		return t.LeaseInfo
	case *ResponseUnion_WriteBatch:
		return t.WriteBatch
	case *ResponseUnion_Export:
		return t.Export
	case *ResponseUnion_Import:
		return t.Import
	case *ResponseUnion_QueryTxn:
		return t.QueryTxn
	case *ResponseUnion_QueryIntent:
		return t.QueryIntent
	case *ResponseUnion_AdminScatter:
		return t.AdminScatter
	case *ResponseUnion_AddSstable:
		return t.AddSstable
	case *ResponseUnion_RecomputeStats:
		return t.RecomputeStats
	case *ResponseUnion_Refresh:
		return t.Refresh
	case *ResponseUnion_RefreshRange:
		return t.RefreshRange
	case *ResponseUnion_Subsume:
		return t.Subsume
	case *ResponseUnion_RangeStats:
		return t.RangeStats
	default:
		return nil
	}
}

// SetInner sets the error in the union.
func (ru *ErrorDetail) SetInner(r error) bool {
	var union isErrorDetail_Value
	switch t := r.(type) {
	case *NotLeaseHolderError:
		union = &ErrorDetail_NotLeaseHolder{t}
	case *RangeNotFoundError:
		union = &ErrorDetail_RangeNotFound{t}
	case *RangeKeyMismatchError:
		union = &ErrorDetail_RangeKeyMismatch{t}
	case *ReadWithinUncertaintyIntervalError:
		union = &ErrorDetail_ReadWithinUncertaintyInterval{t}
	case *TransactionAbortedError:
		union = &ErrorDetail_TransactionAborted{t}
	case *TransactionPushError:
		union = &ErrorDetail_TransactionPush{t}
	case *TransactionRetryError:
		union = &ErrorDetail_TransactionRetry{t}
	case *TransactionStatusError:
		union = &ErrorDetail_TransactionStatus{t}
	case *WriteIntentError:
		union = &ErrorDetail_WriteIntent{t}
	case *WriteTooOldError:
		union = &ErrorDetail_WriteTooOld{t}
	case *OpRequiresTxnError:
		union = &ErrorDetail_OpRequiresTxn{t}
	case *ConditionFailedError:
		union = &ErrorDetail_ConditionFailed{t}
	case *LeaseRejectedError:
		union = &ErrorDetail_LeaseRejected{t}
	case *NodeUnavailableError:
		union = &ErrorDetail_NodeUnavailable{t}
	case *SendError:
		union = &ErrorDetail_Send{t}
	case *RaftGroupDeletedError:
		union = &ErrorDetail_RaftGroupDeleted{t}
	case *ReplicaCorruptionError:
		union = &ErrorDetail_ReplicaCorruption{t}
	case *ReplicaTooOldError:
		union = &ErrorDetail_ReplicaTooOld{t}
	case *AmbiguousResultError:
		union = &ErrorDetail_AmbiguousResult{t}
	case *StoreNotFoundError:
		union = &ErrorDetail_StoreNotFound{t}
	case *TransactionRetryWithProtoRefreshError:
		union = &ErrorDetail_TransactionRetryWithProtoRefresh{t}
	case *IntegerOverflowError:
		union = &ErrorDetail_IntegerOverflow{t}
	case *UnsupportedRequestError:
		union = &ErrorDetail_UnsupportedRequest{t}
	case *BatchTimestampBeforeGCError:
		union = &ErrorDetail_TimestampBefore{t}
	case *TxnAlreadyEncounteredErrorError:
		union = &ErrorDetail_TxnAlreadyEncounteredError{t}
	case *IntentMissingError:
		union = &ErrorDetail_IntentMissing{t}
	case *MergeInProgressError:
		union = &ErrorDetail_MergeInProgress{t}
	case *RangeFeedRetryError:
		union = &ErrorDetail_RangefeedRetry{t}
	case *IndeterminateCommitError:
		union = &ErrorDetail_IndeterminateCommit{t}
	default:
		return false
	}
	ru.Value = union
	return true
}

// SetInner sets the Request in the union.
func (ru *RequestUnion) SetInner(r Request) bool {
	var union isRequestUnion_Value
	switch t := r.(type) {
	case *GetRequest:
		union = &RequestUnion_Get{t}
	case *PutRequest:
		union = &RequestUnion_Put{t}
	case *ConditionalPutRequest:
		union = &RequestUnion_ConditionalPut{t}
	case *IncrementRequest:
		union = &RequestUnion_Increment{t}
	case *DeleteRequest:
		union = &RequestUnion_Delete{t}
	case *DeleteRangeRequest:
		union = &RequestUnion_DeleteRange{t}
	case *ClearRangeRequest:
		union = &RequestUnion_ClearRange{t}
	case *RevertRangeRequest:
		union = &RequestUnion_RevertRange{t}
	case *ScanRequest:
		union = &RequestUnion_Scan{t}
	case *BeginTransactionRequest:
		union = &RequestUnion_BeginTransaction{t}
	case *EndTransactionRequest:
		union = &RequestUnion_EndTransaction{t}
	case *AdminSplitRequest:
		union = &RequestUnion_AdminSplit{t}
	case *AdminUnsplitRequest:
		union = &RequestUnion_AdminUnsplit{t}
	case *AdminMergeRequest:
		union = &RequestUnion_AdminMerge{t}
	case *AdminTransferLeaseRequest:
		union = &RequestUnion_AdminTransferLease{t}
	case *AdminChangeReplicasRequest:
		union = &RequestUnion_AdminChangeReplicas{t}
	case *AdminRelocateRangeRequest:
		union = &RequestUnion_AdminRelocateRange{t}
	case *HeartbeatTxnRequest:
		union = &RequestUnion_HeartbeatTxn{t}
	case *GCRequest:
		union = &RequestUnion_Gc{t}
	case *PushTxnRequest:
		union = &RequestUnion_PushTxn{t}
	case *RecoverTxnRequest:
		union = &RequestUnion_RecoverTxn{t}
	case *ResolveIntentRequest:
		union = &RequestUnion_ResolveIntent{t}
	case *ResolveIntentRangeRequest:
		union = &RequestUnion_ResolveIntentRange{t}
	case *MergeRequest:
		union = &RequestUnion_Merge{t}
	case *TruncateLogRequest:
		union = &RequestUnion_TruncateLog{t}
	case *RequestLeaseRequest:
		union = &RequestUnion_RequestLease{t}
	case *ReverseScanRequest:
		union = &RequestUnion_ReverseScan{t}
	case *ComputeChecksumRequest:
		union = &RequestUnion_ComputeChecksum{t}
	case *CheckConsistencyRequest:
		union = &RequestUnion_CheckConsistency{t}
	case *InitPutRequest:
		union = &RequestUnion_InitPut{t}
	case *TransferLeaseRequest:
		union = &RequestUnion_TransferLease{t}
	case *LeaseInfoRequest:
		union = &RequestUnion_LeaseInfo{t}
	case *WriteBatchRequest:
		union = &RequestUnion_WriteBatch{t}
	case *ExportRequest:
		union = &RequestUnion_Export{t}
	case *ImportRequest:
		union = &RequestUnion_Import{t}
	case *QueryTxnRequest:
		union = &RequestUnion_QueryTxn{t}
	case *QueryIntentRequest:
		union = &RequestUnion_QueryIntent{t}
	case *AdminScatterRequest:
		union = &RequestUnion_AdminScatter{t}
	case *AddSSTableRequest:
		union = &RequestUnion_AddSstable{t}
	case *RecomputeStatsRequest:
		union = &RequestUnion_RecomputeStats{t}
	case *RefreshRequest:
		union = &RequestUnion_Refresh{t}
	case *RefreshRangeRequest:
		union = &RequestUnion_RefreshRange{t}
	case *SubsumeRequest:
		union = &RequestUnion_Subsume{t}
	case *RangeStatsRequest:
		union = &RequestUnion_RangeStats{t}
	default:
		return false
	}
	ru.Value = union
	return true
}

// SetInner sets the Response in the union.
func (ru *ResponseUnion) SetInner(r Response) bool {
	var union isResponseUnion_Value
	switch t := r.(type) {
	case *GetResponse:
		union = &ResponseUnion_Get{t}
	case *PutResponse:
		union = &ResponseUnion_Put{t}
	case *ConditionalPutResponse:
		union = &ResponseUnion_ConditionalPut{t}
	case *IncrementResponse:
		union = &ResponseUnion_Increment{t}
	case *DeleteResponse:
		union = &ResponseUnion_Delete{t}
	case *DeleteRangeResponse:
		union = &ResponseUnion_DeleteRange{t}
	case *ClearRangeResponse:
		union = &ResponseUnion_ClearRange{t}
	case *RevertRangeResponse:
		union = &ResponseUnion_RevertRange{t}
	case *ScanResponse:
		union = &ResponseUnion_Scan{t}
	case *BeginTransactionResponse:
		union = &ResponseUnion_BeginTransaction{t}
	case *EndTransactionResponse:
		union = &ResponseUnion_EndTransaction{t}
	case *AdminSplitResponse:
		union = &ResponseUnion_AdminSplit{t}
	case *AdminUnsplitResponse:
		union = &ResponseUnion_AdminUnsplit{t}
	case *AdminMergeResponse:
		union = &ResponseUnion_AdminMerge{t}
	case *AdminTransferLeaseResponse:
		union = &ResponseUnion_AdminTransferLease{t}
	case *AdminChangeReplicasResponse:
		union = &ResponseUnion_AdminChangeReplicas{t}
	case *AdminRelocateRangeResponse:
		union = &ResponseUnion_AdminRelocateRange{t}
	case *HeartbeatTxnResponse:
		union = &ResponseUnion_HeartbeatTxn{t}
	case *GCResponse:
		union = &ResponseUnion_Gc{t}
	case *PushTxnResponse:
		union = &ResponseUnion_PushTxn{t}
	case *RecoverTxnResponse:
		union = &ResponseUnion_RecoverTxn{t}
	case *ResolveIntentResponse:
		union = &ResponseUnion_ResolveIntent{t}
	case *ResolveIntentRangeResponse:
		union = &ResponseUnion_ResolveIntentRange{t}
	case *MergeResponse:
		union = &ResponseUnion_Merge{t}
	case *TruncateLogResponse:
		union = &ResponseUnion_TruncateLog{t}
	case *RequestLeaseResponse:
		union = &ResponseUnion_RequestLease{t}
	case *ReverseScanResponse:
		union = &ResponseUnion_ReverseScan{t}
	case *ComputeChecksumResponse:
		union = &ResponseUnion_ComputeChecksum{t}
	case *CheckConsistencyResponse:
		union = &ResponseUnion_CheckConsistency{t}
	case *InitPutResponse:
		union = &ResponseUnion_InitPut{t}
	case *LeaseInfoResponse:
		union = &ResponseUnion_LeaseInfo{t}
	case *WriteBatchResponse:
		union = &ResponseUnion_WriteBatch{t}
	case *ExportResponse:
		union = &ResponseUnion_Export{t}
	case *ImportResponse:
		union = &ResponseUnion_Import{t}
	case *QueryTxnResponse:
		union = &ResponseUnion_QueryTxn{t}
	case *QueryIntentResponse:
		union = &ResponseUnion_QueryIntent{t}
	case *AdminScatterResponse:
		union = &ResponseUnion_AdminScatter{t}
	case *AddSSTableResponse:
		union = &ResponseUnion_AddSstable{t}
	case *RecomputeStatsResponse:
		union = &ResponseUnion_RecomputeStats{t}
	case *RefreshResponse:
		union = &ResponseUnion_Refresh{t}
	case *RefreshRangeResponse:
		union = &ResponseUnion_RefreshRange{t}
	case *SubsumeResponse:
		union = &ResponseUnion_Subsume{t}
	case *RangeStatsResponse:
		union = &ResponseUnion_RangeStats{t}
	default:
		return false
	}
	ru.Value = union
	return true
}

type reqCounts [44]int32

// getReqCounts returns the number of times each
// request type appears in the batch.
func (ba *BatchRequest) getReqCounts() reqCounts {
	var counts reqCounts
	for _, ru := range ba.Requests {
		switch ru.GetValue().(type) {
		case *RequestUnion_Get:
			counts[0]++
		case *RequestUnion_Put:
			counts[1]++
		case *RequestUnion_ConditionalPut:
			counts[2]++
		case *RequestUnion_Increment:
			counts[3]++
		case *RequestUnion_Delete:
			counts[4]++
		case *RequestUnion_DeleteRange:
			counts[5]++
		case *RequestUnion_ClearRange:
			counts[6]++
		case *RequestUnion_RevertRange:
			counts[7]++
		case *RequestUnion_Scan:
			counts[8]++
		case *RequestUnion_BeginTransaction:
			counts[9]++
		case *RequestUnion_EndTransaction:
			counts[10]++
		case *RequestUnion_AdminSplit:
			counts[11]++
		case *RequestUnion_AdminUnsplit:
			counts[12]++
		case *RequestUnion_AdminMerge:
			counts[13]++
		case *RequestUnion_AdminTransferLease:
			counts[14]++
		case *RequestUnion_AdminChangeReplicas:
			counts[15]++
		case *RequestUnion_AdminRelocateRange:
			counts[16]++
		case *RequestUnion_HeartbeatTxn:
			counts[17]++
		case *RequestUnion_Gc:
			counts[18]++
		case *RequestUnion_PushTxn:
			counts[19]++
		case *RequestUnion_RecoverTxn:
			counts[20]++
		case *RequestUnion_ResolveIntent:
			counts[21]++
		case *RequestUnion_ResolveIntentRange:
			counts[22]++
		case *RequestUnion_Merge:
			counts[23]++
		case *RequestUnion_TruncateLog:
			counts[24]++
		case *RequestUnion_RequestLease:
			counts[25]++
		case *RequestUnion_ReverseScan:
			counts[26]++
		case *RequestUnion_ComputeChecksum:
			counts[27]++
		case *RequestUnion_CheckConsistency:
			counts[28]++
		case *RequestUnion_InitPut:
			counts[29]++
		case *RequestUnion_TransferLease:
			counts[30]++
		case *RequestUnion_LeaseInfo:
			counts[31]++
		case *RequestUnion_WriteBatch:
			counts[32]++
		case *RequestUnion_Export:
			counts[33]++
		case *RequestUnion_Import:
			counts[34]++
		case *RequestUnion_QueryTxn:
			counts[35]++
		case *RequestUnion_QueryIntent:
			counts[36]++
		case *RequestUnion_AdminScatter:
			counts[37]++
		case *RequestUnion_AddSstable:
			counts[38]++
		case *RequestUnion_RecomputeStats:
			counts[39]++
		case *RequestUnion_Refresh:
			counts[40]++
		case *RequestUnion_RefreshRange:
			counts[41]++
		case *RequestUnion_Subsume:
			counts[42]++
		case *RequestUnion_RangeStats:
			counts[43]++
		default:
			panic(fmt.Sprintf("unsupported request: %+v", ru))
		}
	}
	return counts
}

var requestNames = []string{
	"Get",
	"Put",
	"CPut",
	"Inc",
	"Del",
	"DelRng",
	"ClearRng",
	"RevertRng",
	"Scan",
	"BeginTxn",
	"EndTxn",
	"AdmSplit",
	"AdmUnsplit",
	"AdmMerge",
	"AdmTransferLease",
	"AdmChangeReplicas",
	"AdmRelocateRng",
	"HeartbeatTxn",
	"Gc",
	"PushTxn",
	"RecoverTxn",
	"ResolveIntent",
	"ResolveIntentRng",
	"Merge",
	"TruncLog",
	"RequestLease",
	"RevScan",
	"ComputeChksum",
	"ChkConsistency",
	"InitPut",
	"TransferLease",
	"LeaseInfo",
	"WriteBatch",
	"Export",
	"Import",
	"QueryTxn",
	"QueryIntent",
	"AdmScatter",
	"AddSstable",
	"RecomputeStats",
	"Refresh",
	"RefreshRng",
	"Subsume",
	"RngStats",
}

// Summary prints a short summary of the requests in a batch.
func (ba *BatchRequest) Summary() string {
	var b strings.Builder
	ba.WriteSummary(&b)
	return b.String()
}

// WriteSummary writes a short summary of the requests in a batch
// to the provided builder.
func (ba *BatchRequest) WriteSummary(b *strings.Builder) {
	if len(ba.Requests) == 0 {
		b.WriteString("empty batch")
		return
	}
	counts := ba.getReqCounts()
	var tmp [10]byte
	var comma bool
	for i, v := range counts {
		if v != 0 {
			if comma {
				b.WriteString(", ")
			}
			comma = true

			b.Write(strconv.AppendInt(tmp[:0], int64(v), 10))
			b.WriteString(" ")
			b.WriteString(requestNames[i])
		}
	}
}

// The following types are used to group the allocations of Responses
// and their corresponding isResponseUnion_Value union wrappers together.
type getResponseAlloc struct {
	union ResponseUnion_Get
	resp  GetResponse
}
type putResponseAlloc struct {
	union ResponseUnion_Put
	resp  PutResponse
}
type conditionalPutResponseAlloc struct {
	union ResponseUnion_ConditionalPut
	resp  ConditionalPutResponse
}
type incrementResponseAlloc struct {
	union ResponseUnion_Increment
	resp  IncrementResponse
}
type deleteResponseAlloc struct {
	union ResponseUnion_Delete
	resp  DeleteResponse
}
type deleteRangeResponseAlloc struct {
	union ResponseUnion_DeleteRange
	resp  DeleteRangeResponse
}
type clearRangeResponseAlloc struct {
	union ResponseUnion_ClearRange
	resp  ClearRangeResponse
}
type revertRangeResponseAlloc struct {
	union ResponseUnion_RevertRange
	resp  RevertRangeResponse
}
type scanResponseAlloc struct {
	union ResponseUnion_Scan
	resp  ScanResponse
}
type beginTransactionResponseAlloc struct {
	union ResponseUnion_BeginTransaction
	resp  BeginTransactionResponse
}
type endTransactionResponseAlloc struct {
	union ResponseUnion_EndTransaction
	resp  EndTransactionResponse
}
type adminSplitResponseAlloc struct {
	union ResponseUnion_AdminSplit
	resp  AdminSplitResponse
}
type adminUnsplitResponseAlloc struct {
	union ResponseUnion_AdminUnsplit
	resp  AdminUnsplitResponse
}
type adminMergeResponseAlloc struct {
	union ResponseUnion_AdminMerge
	resp  AdminMergeResponse
}
type adminTransferLeaseResponseAlloc struct {
	union ResponseUnion_AdminTransferLease
	resp  AdminTransferLeaseResponse
}
type adminChangeReplicasResponseAlloc struct {
	union ResponseUnion_AdminChangeReplicas
	resp  AdminChangeReplicasResponse
}
type adminRelocateRangeResponseAlloc struct {
	union ResponseUnion_AdminRelocateRange
	resp  AdminRelocateRangeResponse
}
type heartbeatTxnResponseAlloc struct {
	union ResponseUnion_HeartbeatTxn
	resp  HeartbeatTxnResponse
}
type gCResponseAlloc struct {
	union ResponseUnion_Gc
	resp  GCResponse
}
type pushTxnResponseAlloc struct {
	union ResponseUnion_PushTxn
	resp  PushTxnResponse
}
type recoverTxnResponseAlloc struct {
	union ResponseUnion_RecoverTxn
	resp  RecoverTxnResponse
}
type resolveIntentResponseAlloc struct {
	union ResponseUnion_ResolveIntent
	resp  ResolveIntentResponse
}
type resolveIntentRangeResponseAlloc struct {
	union ResponseUnion_ResolveIntentRange
	resp  ResolveIntentRangeResponse
}
type mergeResponseAlloc struct {
	union ResponseUnion_Merge
	resp  MergeResponse
}
type truncateLogResponseAlloc struct {
	union ResponseUnion_TruncateLog
	resp  TruncateLogResponse
}
type requestLeaseResponseAlloc struct {
	union ResponseUnion_RequestLease
	resp  RequestLeaseResponse
}
type reverseScanResponseAlloc struct {
	union ResponseUnion_ReverseScan
	resp  ReverseScanResponse
}
type computeChecksumResponseAlloc struct {
	union ResponseUnion_ComputeChecksum
	resp  ComputeChecksumResponse
}
type checkConsistencyResponseAlloc struct {
	union ResponseUnion_CheckConsistency
	resp  CheckConsistencyResponse
}
type initPutResponseAlloc struct {
	union ResponseUnion_InitPut
	resp  InitPutResponse
}
type leaseInfoResponseAlloc struct {
	union ResponseUnion_LeaseInfo
	resp  LeaseInfoResponse
}
type writeBatchResponseAlloc struct {
	union ResponseUnion_WriteBatch
	resp  WriteBatchResponse
}
type exportResponseAlloc struct {
	union ResponseUnion_Export
	resp  ExportResponse
}
type importResponseAlloc struct {
	union ResponseUnion_Import
	resp  ImportResponse
}
type queryTxnResponseAlloc struct {
	union ResponseUnion_QueryTxn
	resp  QueryTxnResponse
}
type queryIntentResponseAlloc struct {
	union ResponseUnion_QueryIntent
	resp  QueryIntentResponse
}
type adminScatterResponseAlloc struct {
	union ResponseUnion_AdminScatter
	resp  AdminScatterResponse
}
type addSSTableResponseAlloc struct {
	union ResponseUnion_AddSstable
	resp  AddSSTableResponse
}
type recomputeStatsResponseAlloc struct {
	union ResponseUnion_RecomputeStats
	resp  RecomputeStatsResponse
}
type refreshResponseAlloc struct {
	union ResponseUnion_Refresh
	resp  RefreshResponse
}
type refreshRangeResponseAlloc struct {
	union ResponseUnion_RefreshRange
	resp  RefreshRangeResponse
}
type subsumeResponseAlloc struct {
	union ResponseUnion_Subsume
	resp  SubsumeResponse
}
type rangeStatsResponseAlloc struct {
	union ResponseUnion_RangeStats
	resp  RangeStatsResponse
}

// CreateReply creates replies for each of the contained requests, wrapped in a
// BatchResponse. The response objects are batch allocated to minimize
// allocation overhead.
func (ba *BatchRequest) CreateReply() *BatchResponse {
	br := &BatchResponse{}
	br.Responses = make([]ResponseUnion, len(ba.Requests))

	counts := ba.getReqCounts()

	var buf0 []getResponseAlloc
	var buf1 []putResponseAlloc
	var buf2 []conditionalPutResponseAlloc
	var buf3 []incrementResponseAlloc
	var buf4 []deleteResponseAlloc
	var buf5 []deleteRangeResponseAlloc
	var buf6 []clearRangeResponseAlloc
	var buf7 []revertRangeResponseAlloc
	var buf8 []scanResponseAlloc
	var buf9 []beginTransactionResponseAlloc
	var buf10 []endTransactionResponseAlloc
	var buf11 []adminSplitResponseAlloc
	var buf12 []adminUnsplitResponseAlloc
	var buf13 []adminMergeResponseAlloc
	var buf14 []adminTransferLeaseResponseAlloc
	var buf15 []adminChangeReplicasResponseAlloc
	var buf16 []adminRelocateRangeResponseAlloc
	var buf17 []heartbeatTxnResponseAlloc
	var buf18 []gCResponseAlloc
	var buf19 []pushTxnResponseAlloc
	var buf20 []recoverTxnResponseAlloc
	var buf21 []resolveIntentResponseAlloc
	var buf22 []resolveIntentRangeResponseAlloc
	var buf23 []mergeResponseAlloc
	var buf24 []truncateLogResponseAlloc
	var buf25 []requestLeaseResponseAlloc
	var buf26 []reverseScanResponseAlloc
	var buf27 []computeChecksumResponseAlloc
	var buf28 []checkConsistencyResponseAlloc
	var buf29 []initPutResponseAlloc
	var buf30 []requestLeaseResponseAlloc
	var buf31 []leaseInfoResponseAlloc
	var buf32 []writeBatchResponseAlloc
	var buf33 []exportResponseAlloc
	var buf34 []importResponseAlloc
	var buf35 []queryTxnResponseAlloc
	var buf36 []queryIntentResponseAlloc
	var buf37 []adminScatterResponseAlloc
	var buf38 []addSSTableResponseAlloc
	var buf39 []recomputeStatsResponseAlloc
	var buf40 []refreshResponseAlloc
	var buf41 []refreshRangeResponseAlloc
	var buf42 []subsumeResponseAlloc
	var buf43 []rangeStatsResponseAlloc

	for i, r := range ba.Requests {
		switch r.GetValue().(type) {
		case *RequestUnion_Get:
			if buf0 == nil {
				buf0 = make([]getResponseAlloc, counts[0])
			}
			buf0[0].union.Get = &buf0[0].resp
			br.Responses[i].Value = &buf0[0].union
			buf0 = buf0[1:]
		case *RequestUnion_Put:
			if buf1 == nil {
				buf1 = make([]putResponseAlloc, counts[1])
			}
			buf1[0].union.Put = &buf1[0].resp
			br.Responses[i].Value = &buf1[0].union
			buf1 = buf1[1:]
		case *RequestUnion_ConditionalPut:
			if buf2 == nil {
				buf2 = make([]conditionalPutResponseAlloc, counts[2])
			}
			buf2[0].union.ConditionalPut = &buf2[0].resp
			br.Responses[i].Value = &buf2[0].union
			buf2 = buf2[1:]
		case *RequestUnion_Increment:
			if buf3 == nil {
				buf3 = make([]incrementResponseAlloc, counts[3])
			}
			buf3[0].union.Increment = &buf3[0].resp
			br.Responses[i].Value = &buf3[0].union
			buf3 = buf3[1:]
		case *RequestUnion_Delete:
			if buf4 == nil {
				buf4 = make([]deleteResponseAlloc, counts[4])
			}
			buf4[0].union.Delete = &buf4[0].resp
			br.Responses[i].Value = &buf4[0].union
			buf4 = buf4[1:]
		case *RequestUnion_DeleteRange:
			if buf5 == nil {
				buf5 = make([]deleteRangeResponseAlloc, counts[5])
			}
			buf5[0].union.DeleteRange = &buf5[0].resp
			br.Responses[i].Value = &buf5[0].union
			buf5 = buf5[1:]
		case *RequestUnion_ClearRange:
			if buf6 == nil {
				buf6 = make([]clearRangeResponseAlloc, counts[6])
			}
			buf6[0].union.ClearRange = &buf6[0].resp
			br.Responses[i].Value = &buf6[0].union
			buf6 = buf6[1:]
		case *RequestUnion_RevertRange:
			if buf7 == nil {
				buf7 = make([]revertRangeResponseAlloc, counts[7])
			}
			buf7[0].union.RevertRange = &buf7[0].resp
			br.Responses[i].Value = &buf7[0].union
			buf7 = buf7[1:]
		case *RequestUnion_Scan:
			if buf8 == nil {
				buf8 = make([]scanResponseAlloc, counts[8])
			}
			buf8[0].union.Scan = &buf8[0].resp
			br.Responses[i].Value = &buf8[0].union
			buf8 = buf8[1:]
		case *RequestUnion_BeginTransaction:
			if buf9 == nil {
				buf9 = make([]beginTransactionResponseAlloc, counts[9])
			}
			buf9[0].union.BeginTransaction = &buf9[0].resp
			br.Responses[i].Value = &buf9[0].union
			buf9 = buf9[1:]
		case *RequestUnion_EndTransaction:
			if buf10 == nil {
				buf10 = make([]endTransactionResponseAlloc, counts[10])
			}
			buf10[0].union.EndTransaction = &buf10[0].resp
			br.Responses[i].Value = &buf10[0].union
			buf10 = buf10[1:]
		case *RequestUnion_AdminSplit:
			if buf11 == nil {
				buf11 = make([]adminSplitResponseAlloc, counts[11])
			}
			buf11[0].union.AdminSplit = &buf11[0].resp
			br.Responses[i].Value = &buf11[0].union
			buf11 = buf11[1:]
		case *RequestUnion_AdminUnsplit:
			if buf12 == nil {
				buf12 = make([]adminUnsplitResponseAlloc, counts[12])
			}
			buf12[0].union.AdminUnsplit = &buf12[0].resp
			br.Responses[i].Value = &buf12[0].union
			buf12 = buf12[1:]
		case *RequestUnion_AdminMerge:
			if buf13 == nil {
				buf13 = make([]adminMergeResponseAlloc, counts[13])
			}
			buf13[0].union.AdminMerge = &buf13[0].resp
			br.Responses[i].Value = &buf13[0].union
			buf13 = buf13[1:]
		case *RequestUnion_AdminTransferLease:
			if buf14 == nil {
				buf14 = make([]adminTransferLeaseResponseAlloc, counts[14])
			}
			buf14[0].union.AdminTransferLease = &buf14[0].resp
			br.Responses[i].Value = &buf14[0].union
			buf14 = buf14[1:]
		case *RequestUnion_AdminChangeReplicas:
			if buf15 == nil {
				buf15 = make([]adminChangeReplicasResponseAlloc, counts[15])
			}
			buf15[0].union.AdminChangeReplicas = &buf15[0].resp
			br.Responses[i].Value = &buf15[0].union
			buf15 = buf15[1:]
		case *RequestUnion_AdminRelocateRange:
			if buf16 == nil {
				buf16 = make([]adminRelocateRangeResponseAlloc, counts[16])
			}
			buf16[0].union.AdminRelocateRange = &buf16[0].resp
			br.Responses[i].Value = &buf16[0].union
			buf16 = buf16[1:]
		case *RequestUnion_HeartbeatTxn:
			if buf17 == nil {
				buf17 = make([]heartbeatTxnResponseAlloc, counts[17])
			}
			buf17[0].union.HeartbeatTxn = &buf17[0].resp
			br.Responses[i].Value = &buf17[0].union
			buf17 = buf17[1:]
		case *RequestUnion_Gc:
			if buf18 == nil {
				buf18 = make([]gCResponseAlloc, counts[18])
			}
			buf18[0].union.Gc = &buf18[0].resp
			br.Responses[i].Value = &buf18[0].union
			buf18 = buf18[1:]
		case *RequestUnion_PushTxn:
			if buf19 == nil {
				buf19 = make([]pushTxnResponseAlloc, counts[19])
			}
			buf19[0].union.PushTxn = &buf19[0].resp
			br.Responses[i].Value = &buf19[0].union
			buf19 = buf19[1:]
		case *RequestUnion_RecoverTxn:
			if buf20 == nil {
				buf20 = make([]recoverTxnResponseAlloc, counts[20])
			}
			buf20[0].union.RecoverTxn = &buf20[0].resp
			br.Responses[i].Value = &buf20[0].union
			buf20 = buf20[1:]
		case *RequestUnion_ResolveIntent:
			if buf21 == nil {
				buf21 = make([]resolveIntentResponseAlloc, counts[21])
			}
			buf21[0].union.ResolveIntent = &buf21[0].resp
			br.Responses[i].Value = &buf21[0].union
			buf21 = buf21[1:]
		case *RequestUnion_ResolveIntentRange:
			if buf22 == nil {
				buf22 = make([]resolveIntentRangeResponseAlloc, counts[22])
			}
			buf22[0].union.ResolveIntentRange = &buf22[0].resp
			br.Responses[i].Value = &buf22[0].union
			buf22 = buf22[1:]
		case *RequestUnion_Merge:
			if buf23 == nil {
				buf23 = make([]mergeResponseAlloc, counts[23])
			}
			buf23[0].union.Merge = &buf23[0].resp
			br.Responses[i].Value = &buf23[0].union
			buf23 = buf23[1:]
		case *RequestUnion_TruncateLog:
			if buf24 == nil {
				buf24 = make([]truncateLogResponseAlloc, counts[24])
			}
			buf24[0].union.TruncateLog = &buf24[0].resp
			br.Responses[i].Value = &buf24[0].union
			buf24 = buf24[1:]
		case *RequestUnion_RequestLease:
			if buf25 == nil {
				buf25 = make([]requestLeaseResponseAlloc, counts[25])
			}
			buf25[0].union.RequestLease = &buf25[0].resp
			br.Responses[i].Value = &buf25[0].union
			buf25 = buf25[1:]
		case *RequestUnion_ReverseScan:
			if buf26 == nil {
				buf26 = make([]reverseScanResponseAlloc, counts[26])
			}
			buf26[0].union.ReverseScan = &buf26[0].resp
			br.Responses[i].Value = &buf26[0].union
			buf26 = buf26[1:]
		case *RequestUnion_ComputeChecksum:
			if buf27 == nil {
				buf27 = make([]computeChecksumResponseAlloc, counts[27])
			}
			buf27[0].union.ComputeChecksum = &buf27[0].resp
			br.Responses[i].Value = &buf27[0].union
			buf27 = buf27[1:]
		case *RequestUnion_CheckConsistency:
			if buf28 == nil {
				buf28 = make([]checkConsistencyResponseAlloc, counts[28])
			}
			buf28[0].union.CheckConsistency = &buf28[0].resp
			br.Responses[i].Value = &buf28[0].union
			buf28 = buf28[1:]
		case *RequestUnion_InitPut:
			if buf29 == nil {
				buf29 = make([]initPutResponseAlloc, counts[29])
			}
			buf29[0].union.InitPut = &buf29[0].resp
			br.Responses[i].Value = &buf29[0].union
			buf29 = buf29[1:]
		case *RequestUnion_TransferLease:
			if buf30 == nil {
				buf30 = make([]requestLeaseResponseAlloc, counts[30])
			}
			buf30[0].union.RequestLease = &buf30[0].resp
			br.Responses[i].Value = &buf30[0].union
			buf30 = buf30[1:]
		case *RequestUnion_LeaseInfo:
			if buf31 == nil {
				buf31 = make([]leaseInfoResponseAlloc, counts[31])
			}
			buf31[0].union.LeaseInfo = &buf31[0].resp
			br.Responses[i].Value = &buf31[0].union
			buf31 = buf31[1:]
		case *RequestUnion_WriteBatch:
			if buf32 == nil {
				buf32 = make([]writeBatchResponseAlloc, counts[32])
			}
			buf32[0].union.WriteBatch = &buf32[0].resp
			br.Responses[i].Value = &buf32[0].union
			buf32 = buf32[1:]
		case *RequestUnion_Export:
			if buf33 == nil {
				buf33 = make([]exportResponseAlloc, counts[33])
			}
			buf33[0].union.Export = &buf33[0].resp
			br.Responses[i].Value = &buf33[0].union
			buf33 = buf33[1:]
		case *RequestUnion_Import:
			if buf34 == nil {
				buf34 = make([]importResponseAlloc, counts[34])
			}
			buf34[0].union.Import = &buf34[0].resp
			br.Responses[i].Value = &buf34[0].union
			buf34 = buf34[1:]
		case *RequestUnion_QueryTxn:
			if buf35 == nil {
				buf35 = make([]queryTxnResponseAlloc, counts[35])
			}
			buf35[0].union.QueryTxn = &buf35[0].resp
			br.Responses[i].Value = &buf35[0].union
			buf35 = buf35[1:]
		case *RequestUnion_QueryIntent:
			if buf36 == nil {
				buf36 = make([]queryIntentResponseAlloc, counts[36])
			}
			buf36[0].union.QueryIntent = &buf36[0].resp
			br.Responses[i].Value = &buf36[0].union
			buf36 = buf36[1:]
		case *RequestUnion_AdminScatter:
			if buf37 == nil {
				buf37 = make([]adminScatterResponseAlloc, counts[37])
			}
			buf37[0].union.AdminScatter = &buf37[0].resp
			br.Responses[i].Value = &buf37[0].union
			buf37 = buf37[1:]
		case *RequestUnion_AddSstable:
			if buf38 == nil {
				buf38 = make([]addSSTableResponseAlloc, counts[38])
			}
			buf38[0].union.AddSstable = &buf38[0].resp
			br.Responses[i].Value = &buf38[0].union
			buf38 = buf38[1:]
		case *RequestUnion_RecomputeStats:
			if buf39 == nil {
				buf39 = make([]recomputeStatsResponseAlloc, counts[39])
			}
			buf39[0].union.RecomputeStats = &buf39[0].resp
			br.Responses[i].Value = &buf39[0].union
			buf39 = buf39[1:]
		case *RequestUnion_Refresh:
			if buf40 == nil {
				buf40 = make([]refreshResponseAlloc, counts[40])
			}
			buf40[0].union.Refresh = &buf40[0].resp
			br.Responses[i].Value = &buf40[0].union
			buf40 = buf40[1:]
		case *RequestUnion_RefreshRange:
			if buf41 == nil {
				buf41 = make([]refreshRangeResponseAlloc, counts[41])
			}
			buf41[0].union.RefreshRange = &buf41[0].resp
			br.Responses[i].Value = &buf41[0].union
			buf41 = buf41[1:]
		case *RequestUnion_Subsume:
			if buf42 == nil {
				buf42 = make([]subsumeResponseAlloc, counts[42])
			}
			buf42[0].union.Subsume = &buf42[0].resp
			br.Responses[i].Value = &buf42[0].union
			buf42 = buf42[1:]
		case *RequestUnion_RangeStats:
			if buf43 == nil {
				buf43 = make([]rangeStatsResponseAlloc, counts[43])
			}
			buf43[0].union.RangeStats = &buf43[0].resp
			br.Responses[i].Value = &buf43[0].union
			buf43 = buf43[1:]
		default:
			panic(fmt.Sprintf("unsupported request: %+v", r))
		}
	}
	return br
}
