// Code generated by "stringer -output=pkg/sql/opt/rule_name_string.go -type=RuleName pkg/sql/opt/rule_name.go pkg/sql/opt/rule_name.og.go"; DO NOT EDIT.

package opt

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[InvalidRuleName-0]
	_ = x[SimplifyRootOrdering-1]
	_ = x[PruneRootCols-2]
	_ = x[SimplifyZeroCardinalityGroup-3]
	_ = x[NumManualRuleNames-4]
	_ = x[startAutoRule-4]
	_ = x[EliminateAggDistinct-5]
	_ = x[NormalizeNestedAnds-6]
	_ = x[SimplifyTrueAnd-7]
	_ = x[SimplifyAndTrue-8]
	_ = x[SimplifyFalseAnd-9]
	_ = x[SimplifyAndFalse-10]
	_ = x[SimplifyTrueOr-11]
	_ = x[SimplifyOrTrue-12]
	_ = x[SimplifyFalseOr-13]
	_ = x[SimplifyOrFalse-14]
	_ = x[SimplifyRange-15]
	_ = x[FoldNullAndOr-16]
	_ = x[FoldNotTrue-17]
	_ = x[FoldNotFalse-18]
	_ = x[NegateComparison-19]
	_ = x[EliminateNot-20]
	_ = x[NegateAnd-21]
	_ = x[NegateOr-22]
	_ = x[ExtractRedundantConjunct-23]
	_ = x[CommuteVarInequality-24]
	_ = x[CommuteConstInequality-25]
	_ = x[NormalizeCmpPlusConst-26]
	_ = x[NormalizeCmpMinusConst-27]
	_ = x[NormalizeCmpConstMinus-28]
	_ = x[NormalizeTupleEquality-29]
	_ = x[FoldNullComparisonLeft-30]
	_ = x[FoldNullComparisonRight-31]
	_ = x[FoldIsNull-32]
	_ = x[FoldNonNullIsNull-33]
	_ = x[FoldIsNotNull-34]
	_ = x[FoldNonNullIsNotNull-35]
	_ = x[CommuteNullIs-36]
	_ = x[DecorrelateJoin-37]
	_ = x[DecorrelateProjectSet-38]
	_ = x[TryDecorrelateSelect-39]
	_ = x[TryDecorrelateProject-40]
	_ = x[TryDecorrelateProjectSelect-41]
	_ = x[TryDecorrelateProjectInnerJoin-42]
	_ = x[TryDecorrelateInnerJoin-43]
	_ = x[TryDecorrelateInnerLeftJoin-44]
	_ = x[TryDecorrelateGroupBy-45]
	_ = x[TryDecorrelateScalarGroupBy-46]
	_ = x[TryDecorrelateSemiJoin-47]
	_ = x[TryDecorrelateLimitOne-48]
	_ = x[TryDecorrelateProjectSet-49]
	_ = x[TryDecorrelateWindow-50]
	_ = x[HoistSelectExists-51]
	_ = x[HoistSelectNotExists-52]
	_ = x[HoistSelectSubquery-53]
	_ = x[HoistProjectSubquery-54]
	_ = x[HoistJoinSubquery-55]
	_ = x[HoistValuesSubquery-56]
	_ = x[HoistProjectSetSubquery-57]
	_ = x[NormalizeSelectAnyFilter-58]
	_ = x[NormalizeJoinAnyFilter-59]
	_ = x[NormalizeSelectNotAnyFilter-60]
	_ = x[NormalizeJoinNotAnyFilter-61]
	_ = x[FoldNullCast-62]
	_ = x[FoldNullUnary-63]
	_ = x[FoldNullBinaryLeft-64]
	_ = x[FoldNullBinaryRight-65]
	_ = x[FoldNullInNonEmpty-66]
	_ = x[FoldNullInEmpty-67]
	_ = x[FoldNullNotInEmpty-68]
	_ = x[FoldArray-69]
	_ = x[FoldBinary-70]
	_ = x[FoldUnary-71]
	_ = x[FoldComparison-72]
	_ = x[FoldCast-73]
	_ = x[FoldIndirection-74]
	_ = x[FoldColumnAccess-75]
	_ = x[FoldFunction-76]
	_ = x[FoldEqualsAnyNull-77]
	_ = x[ConvertGroupByToDistinct-78]
	_ = x[EliminateDistinct-79]
	_ = x[EliminateGroupByProject-80]
	_ = x[ReduceGroupingCols-81]
	_ = x[EliminateAggDistinctForKeys-82]
	_ = x[EliminateDistinctOnNoColumns-83]
	_ = x[InlineProjectConstants-84]
	_ = x[InlineSelectConstants-85]
	_ = x[InlineJoinConstantsLeft-86]
	_ = x[InlineJoinConstantsRight-87]
	_ = x[PushSelectIntoInlinableProject-88]
	_ = x[InlineProjectInProject-89]
	_ = x[SimplifyJoinFilters-90]
	_ = x[DetectJoinContradiction-91]
	_ = x[PushFilterIntoJoinLeftAndRight-92]
	_ = x[MapFilterIntoJoinLeft-93]
	_ = x[MapFilterIntoJoinRight-94]
	_ = x[MapEqualityIntoJoinLeftAndRight-95]
	_ = x[PushFilterIntoJoinLeft-96]
	_ = x[PushFilterIntoJoinRight-97]
	_ = x[SimplifyLeftJoinWithoutFilters-98]
	_ = x[SimplifyRightJoinWithoutFilters-99]
	_ = x[SimplifyLeftJoinWithFilters-100]
	_ = x[SimplifyRightJoinWithFilters-101]
	_ = x[EliminateSemiJoin-102]
	_ = x[SimplifyZeroCardinalitySemiJoin-103]
	_ = x[EliminateAntiJoin-104]
	_ = x[SimplifyZeroCardinalityAntiJoin-105]
	_ = x[EliminateJoinNoColsLeft-106]
	_ = x[EliminateJoinNoColsRight-107]
	_ = x[HoistJoinProjectRight-108]
	_ = x[HoistJoinProjectLeft-109]
	_ = x[SimplifyJoinNotNullEquality-110]
	_ = x[ExtractJoinEqualities-111]
	_ = x[SortFiltersInJoin-112]
	_ = x[EliminateLimit-113]
	_ = x[EliminateOffset-114]
	_ = x[PushLimitIntoProject-115]
	_ = x[PushOffsetIntoProject-116]
	_ = x[PushLimitIntoOffset-117]
	_ = x[PushLimitIntoOrdinality-118]
	_ = x[EliminateMax1Row-119]
	_ = x[FoldPlusZero-120]
	_ = x[FoldZeroPlus-121]
	_ = x[FoldMinusZero-122]
	_ = x[FoldMultOne-123]
	_ = x[FoldOneMult-124]
	_ = x[FoldDivOne-125]
	_ = x[InvertMinus-126]
	_ = x[EliminateUnaryMinus-127]
	_ = x[SimplifyLimitOrdering-128]
	_ = x[SimplifyOffsetOrdering-129]
	_ = x[SimplifyGroupByOrdering-130]
	_ = x[SimplifyOrdinalityOrdering-131]
	_ = x[SimplifyExplainOrdering-132]
	_ = x[EliminateProject-133]
	_ = x[MergeProjects-134]
	_ = x[MergeProjectWithValues-135]
	_ = x[PruneProjectCols-136]
	_ = x[PruneScanCols-137]
	_ = x[PruneSelectCols-138]
	_ = x[PruneLimitCols-139]
	_ = x[PruneOffsetCols-140]
	_ = x[PruneJoinLeftCols-141]
	_ = x[PruneJoinRightCols-142]
	_ = x[PruneSemiAntiJoinRightCols-143]
	_ = x[PruneAggCols-144]
	_ = x[PruneGroupByCols-145]
	_ = x[PruneValuesCols-146]
	_ = x[PruneOrdinalityCols-147]
	_ = x[PruneExplainCols-148]
	_ = x[PruneProjectSetCols-149]
	_ = x[PruneWindowOutputCols-150]
	_ = x[PruneWindowInputCols-151]
	_ = x[PruneMutationFetchCols-152]
	_ = x[PruneMutationInputCols-153]
	_ = x[PruneMutationReturnCols-154]
	_ = x[PruneWithScanCols-155]
	_ = x[PruneWithCols-156]
	_ = x[PruneUnionAllCols-157]
	_ = x[RejectNullsLeftJoin-158]
	_ = x[RejectNullsRightJoin-159]
	_ = x[RejectNullsGroupBy-160]
	_ = x[CommuteVar-161]
	_ = x[CommuteConst-162]
	_ = x[EliminateCoalesce-163]
	_ = x[SimplifyCoalesce-164]
	_ = x[EliminateCast-165]
	_ = x[NormalizeInConst-166]
	_ = x[FoldInNull-167]
	_ = x[UnifyComparisonTypes-168]
	_ = x[EliminateExistsProject-169]
	_ = x[EliminateExistsGroupBy-170]
	_ = x[IntroduceExistsLimit-171]
	_ = x[NormalizeJSONFieldAccess-172]
	_ = x[NormalizeJSONContains-173]
	_ = x[SimplifyCaseWhenConstValue-174]
	_ = x[SimplifyEqualsAnyTuple-175]
	_ = x[SimplifyAnyScalarArray-176]
	_ = x[FoldCollate-177]
	_ = x[NormalizeArrayFlattenToAgg-178]
	_ = x[SimplifySelectFilters-179]
	_ = x[ConsolidateSelectFilters-180]
	_ = x[DetectSelectContradiction-181]
	_ = x[EliminateSelect-182]
	_ = x[MergeSelects-183]
	_ = x[PushSelectIntoProject-184]
	_ = x[MergeSelectInnerJoin-185]
	_ = x[PushSelectCondLeftIntoJoinLeftAndRight-186]
	_ = x[PushSelectCondRightIntoJoinLeftAndRight-187]
	_ = x[PushSelectIntoJoinLeft-188]
	_ = x[PushSelectIntoJoinRight-189]
	_ = x[PushSelectIntoGroupBy-190]
	_ = x[RemoveNotNullCondition-191]
	_ = x[InlineConstVar-192]
	_ = x[EliminateUnionAllLeft-193]
	_ = x[EliminateUnionAllRight-194]
	_ = x[PushFilterIntoSetOp-195]
	_ = x[EliminateWindow-196]
	_ = x[ReduceWindowPartitionCols-197]
	_ = x[SimplifyWindowOrdering-198]
	_ = x[PushSelectIntoWindow-199]
	_ = x[PushLimitIntoWindow-200]
	_ = x[InlineWith-201]
	_ = x[startExploreRule-202]
	_ = x[ReplaceMinWithLimit-203]
	_ = x[ReplaceMaxWithLimit-204]
	_ = x[GenerateStreamingGroupBy-205]
	_ = x[CommuteJoin-206]
	_ = x[CommuteLeftJoin-207]
	_ = x[CommuteRightJoin-208]
	_ = x[CommuteSemiJoin-209]
	_ = x[GenerateMergeJoins-210]
	_ = x[GenerateLookupJoins-211]
	_ = x[GenerateZigzagJoins-212]
	_ = x[GenerateInvertedIndexZigzagJoins-213]
	_ = x[GenerateLookupJoinsWithFilter-214]
	_ = x[AssociateJoin-215]
	_ = x[GenerateLimitedScans-216]
	_ = x[PushLimitIntoConstrainedScan-217]
	_ = x[PushLimitIntoIndexJoin-218]
	_ = x[GenerateIndexScans-219]
	_ = x[GenerateConstrainedScans-220]
	_ = x[GenerateInvertedIndexScans-221]
	_ = x[NumRuleNames-222]
}

const _RuleName_name = "InvalidRuleNameSimplifyRootOrderingPruneRootColsSimplifyZeroCardinalityGroupNumManualRuleNamesEliminateAggDistinctNormalizeNestedAndsSimplifyTrueAndSimplifyAndTrueSimplifyFalseAndSimplifyAndFalseSimplifyTrueOrSimplifyOrTrueSimplifyFalseOrSimplifyOrFalseSimplifyRangeFoldNullAndOrFoldNotTrueFoldNotFalseNegateComparisonEliminateNotNegateAndNegateOrExtractRedundantConjunctCommuteVarInequalityCommuteConstInequalityNormalizeCmpPlusConstNormalizeCmpMinusConstNormalizeCmpConstMinusNormalizeTupleEqualityFoldNullComparisonLeftFoldNullComparisonRightFoldIsNullFoldNonNullIsNullFoldIsNotNullFoldNonNullIsNotNullCommuteNullIsDecorrelateJoinDecorrelateProjectSetTryDecorrelateSelectTryDecorrelateProjectTryDecorrelateProjectSelectTryDecorrelateProjectInnerJoinTryDecorrelateInnerJoinTryDecorrelateInnerLeftJoinTryDecorrelateGroupByTryDecorrelateScalarGroupByTryDecorrelateSemiJoinTryDecorrelateLimitOneTryDecorrelateProjectSetTryDecorrelateWindowHoistSelectExistsHoistSelectNotExistsHoistSelectSubqueryHoistProjectSubqueryHoistJoinSubqueryHoistValuesSubqueryHoistProjectSetSubqueryNormalizeSelectAnyFilterNormalizeJoinAnyFilterNormalizeSelectNotAnyFilterNormalizeJoinNotAnyFilterFoldNullCastFoldNullUnaryFoldNullBinaryLeftFoldNullBinaryRightFoldNullInNonEmptyFoldNullInEmptyFoldNullNotInEmptyFoldArrayFoldBinaryFoldUnaryFoldComparisonFoldCastFoldIndirectionFoldColumnAccessFoldFunctionFoldEqualsAnyNullConvertGroupByToDistinctEliminateDistinctEliminateGroupByProjectReduceGroupingColsEliminateAggDistinctForKeysEliminateDistinctOnNoColumnsInlineProjectConstantsInlineSelectConstantsInlineJoinConstantsLeftInlineJoinConstantsRightPushSelectIntoInlinableProjectInlineProjectInProjectSimplifyJoinFiltersDetectJoinContradictionPushFilterIntoJoinLeftAndRightMapFilterIntoJoinLeftMapFilterIntoJoinRightMapEqualityIntoJoinLeftAndRightPushFilterIntoJoinLeftPushFilterIntoJoinRightSimplifyLeftJoinWithoutFiltersSimplifyRightJoinWithoutFiltersSimplifyLeftJoinWithFiltersSimplifyRightJoinWithFiltersEliminateSemiJoinSimplifyZeroCardinalitySemiJoinEliminateAntiJoinSimplifyZeroCardinalityAntiJoinEliminateJoinNoColsLeftEliminateJoinNoColsRightHoistJoinProjectRightHoistJoinProjectLeftSimplifyJoinNotNullEqualityExtractJoinEqualitiesSortFiltersInJoinEliminateLimitEliminateOffsetPushLimitIntoProjectPushOffsetIntoProjectPushLimitIntoOffsetPushLimitIntoOrdinalityEliminateMax1RowFoldPlusZeroFoldZeroPlusFoldMinusZeroFoldMultOneFoldOneMultFoldDivOneInvertMinusEliminateUnaryMinusSimplifyLimitOrderingSimplifyOffsetOrderingSimplifyGroupByOrderingSimplifyOrdinalityOrderingSimplifyExplainOrderingEliminateProjectMergeProjectsMergeProjectWithValuesPruneProjectColsPruneScanColsPruneSelectColsPruneLimitColsPruneOffsetColsPruneJoinLeftColsPruneJoinRightColsPruneSemiAntiJoinRightColsPruneAggColsPruneGroupByColsPruneValuesColsPruneOrdinalityColsPruneExplainColsPruneProjectSetColsPruneWindowOutputColsPruneWindowInputColsPruneMutationFetchColsPruneMutationInputColsPruneMutationReturnColsPruneWithScanColsPruneWithColsPruneUnionAllColsRejectNullsLeftJoinRejectNullsRightJoinRejectNullsGroupByCommuteVarCommuteConstEliminateCoalesceSimplifyCoalesceEliminateCastNormalizeInConstFoldInNullUnifyComparisonTypesEliminateExistsProjectEliminateExistsGroupByIntroduceExistsLimitNormalizeJSONFieldAccessNormalizeJSONContainsSimplifyCaseWhenConstValueSimplifyEqualsAnyTupleSimplifyAnyScalarArrayFoldCollateNormalizeArrayFlattenToAggSimplifySelectFiltersConsolidateSelectFiltersDetectSelectContradictionEliminateSelectMergeSelectsPushSelectIntoProjectMergeSelectInnerJoinPushSelectCondLeftIntoJoinLeftAndRightPushSelectCondRightIntoJoinLeftAndRightPushSelectIntoJoinLeftPushSelectIntoJoinRightPushSelectIntoGroupByRemoveNotNullConditionInlineConstVarEliminateUnionAllLeftEliminateUnionAllRightPushFilterIntoSetOpEliminateWindowReduceWindowPartitionColsSimplifyWindowOrderingPushSelectIntoWindowPushLimitIntoWindowInlineWithstartExploreRuleReplaceMinWithLimitReplaceMaxWithLimitGenerateStreamingGroupByCommuteJoinCommuteLeftJoinCommuteRightJoinCommuteSemiJoinGenerateMergeJoinsGenerateLookupJoinsGenerateZigzagJoinsGenerateInvertedIndexZigzagJoinsGenerateLookupJoinsWithFilterAssociateJoinGenerateLimitedScansPushLimitIntoConstrainedScanPushLimitIntoIndexJoinGenerateIndexScansGenerateConstrainedScansGenerateInvertedIndexScansNumRuleNames"

var _RuleName_index = [...]uint16{0, 15, 35, 48, 76, 94, 114, 133, 148, 163, 179, 195, 209, 223, 238, 253, 266, 279, 290, 302, 318, 330, 339, 347, 371, 391, 413, 434, 456, 478, 500, 522, 545, 555, 572, 585, 605, 618, 633, 654, 674, 695, 722, 752, 775, 802, 823, 850, 872, 894, 918, 938, 955, 975, 994, 1014, 1031, 1050, 1073, 1097, 1119, 1146, 1171, 1183, 1196, 1214, 1233, 1251, 1266, 1284, 1293, 1303, 1312, 1326, 1334, 1349, 1365, 1377, 1394, 1418, 1435, 1458, 1476, 1503, 1531, 1553, 1574, 1597, 1621, 1651, 1673, 1692, 1715, 1745, 1766, 1788, 1819, 1841, 1864, 1894, 1925, 1952, 1980, 1997, 2028, 2045, 2076, 2099, 2123, 2144, 2164, 2191, 2212, 2229, 2243, 2258, 2278, 2299, 2318, 2341, 2357, 2369, 2381, 2394, 2405, 2416, 2426, 2437, 2456, 2477, 2499, 2522, 2548, 2571, 2587, 2600, 2622, 2638, 2651, 2666, 2680, 2695, 2712, 2730, 2756, 2768, 2784, 2799, 2818, 2834, 2853, 2874, 2894, 2916, 2938, 2961, 2978, 2991, 3008, 3027, 3047, 3065, 3075, 3087, 3104, 3120, 3133, 3149, 3159, 3179, 3201, 3223, 3243, 3267, 3288, 3314, 3336, 3358, 3369, 3395, 3416, 3440, 3465, 3480, 3492, 3513, 3533, 3571, 3610, 3632, 3655, 3676, 3698, 3712, 3733, 3755, 3774, 3789, 3814, 3836, 3856, 3875, 3885, 3901, 3920, 3939, 3963, 3974, 3989, 4005, 4020, 4038, 4057, 4076, 4108, 4137, 4150, 4170, 4198, 4220, 4238, 4262, 4288, 4300}

func (i RuleName) String() string {
	if i >= RuleName(len(_RuleName_index)-1) {
		return "RuleName(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _RuleName_name[_RuleName_index[i]:_RuleName_index[i+1]]
}
