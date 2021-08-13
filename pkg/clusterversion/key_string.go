// Code generated by "stringer -type=Key"; DO NOT EDIT.

package clusterversion

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Start20_2-0]
	_ = x[NodeMembershipStatus-1]
	_ = x[MinPasswordLength-2]
	_ = x[AbortSpanBytes-3]
	_ = x[CreateLoginPrivilege-4]
	_ = x[HBAForNonTLS-5]
	_ = x[V20_2-6]
	_ = x[Start21_1-7]
	_ = x[CPutInline-8]
	_ = x[ReplicaVersions-9]
	_ = x[replacedTruncatedAndRangeAppliedStateMigration-10]
	_ = x[replacedPostTruncatedAndRangeAppliedStateMigration-11]
	_ = x[TruncatedAndRangeAppliedStateMigration-12]
	_ = x[PostTruncatedAndRangeAppliedStateMigration-13]
	_ = x[SeparatedIntents-14]
	_ = x[TracingVerbosityIndependentSemantics-15]
	_ = x[ClosedTimestampsRaftTransport-16]
	_ = x[PriorReadSummaries-17]
	_ = x[NonVotingReplicas-18]
	_ = x[V21_1-19]
	_ = x[Start21_1PLUS-20]
	_ = x[Start21_2-21]
	_ = x[JoinTokensTable-22]
	_ = x[AcquisitionTypeInLeaseHistory-23]
	_ = x[SerializeViewUDTs-24]
	_ = x[ExpressionIndexes-25]
	_ = x[DeleteDeprecatedNamespaceTableDescriptorMigration-26]
	_ = x[FixDescriptors-27]
	_ = x[SQLStatsTable-28]
	_ = x[DatabaseRoleSettings-29]
	_ = x[TenantUsageTable-30]
	_ = x[SQLInstancesTable-31]
	_ = x[NewRetryableRangefeedErrors-32]
	_ = x[AlterSystemWebSessionsCreateIndexes-33]
	_ = x[SeparatedIntentsMigration-34]
	_ = x[PostSeparatedIntentsMigration-35]
}

const _Key_name = "Start20_2NodeMembershipStatusMinPasswordLengthAbortSpanBytesCreateLoginPrivilegeHBAForNonTLSV20_2Start21_1CPutInlineReplicaVersionsreplacedTruncatedAndRangeAppliedStateMigrationreplacedPostTruncatedAndRangeAppliedStateMigrationTruncatedAndRangeAppliedStateMigrationPostTruncatedAndRangeAppliedStateMigrationSeparatedIntentsTracingVerbosityIndependentSemanticsClosedTimestampsRaftTransportPriorReadSummariesNonVotingReplicasV21_1Start21_1PLUSStart21_2JoinTokensTableAcquisitionTypeInLeaseHistorySerializeViewUDTsExpressionIndexesDeleteDeprecatedNamespaceTableDescriptorMigrationFixDescriptorsSQLStatsTableDatabaseRoleSettingsTenantUsageTableSQLInstancesTableNewRetryableRangefeedErrorsAlterSystemWebSessionsCreateIndexesSeparatedIntentsMigrationPostSeparatedIntentsMigration"

var _Key_index = [...]uint16{0, 9, 29, 46, 60, 80, 92, 97, 106, 116, 131, 177, 227, 265, 307, 323, 359, 388, 406, 423, 428, 441, 450, 465, 494, 511, 528, 577, 591, 604, 624, 640, 657, 684, 719, 744, 773}

func (i Key) String() string {
	if i < 0 || i >= Key(len(_Key_index)-1) {
		return "Key(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Key_name[_Key_index[i]:_Key_index[i+1]]
}
