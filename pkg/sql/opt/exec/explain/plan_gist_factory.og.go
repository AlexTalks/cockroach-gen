// Code generated by optgen; DO NOT EDIT.

package explain

import (
	"github.com/cockroachdb/cockroach/pkg/sql/catalog/colinfo"
	"github.com/cockroachdb/cockroach/pkg/sql/catalog/descpb"
	"github.com/cockroachdb/cockroach/pkg/sql/inverted"
	"github.com/cockroachdb/cockroach/pkg/sql/opt"
	"github.com/cockroachdb/cockroach/pkg/sql/opt/cat"
	"github.com/cockroachdb/cockroach/pkg/sql/opt/constraint"
	"github.com/cockroachdb/cockroach/pkg/sql/opt/exec"
	"github.com/cockroachdb/cockroach/pkg/sql/sem/tree"
	"github.com/cockroachdb/cockroach/pkg/sql/types"
	"github.com/cockroachdb/errors"
)

func (f *PlanGistFactory) ConstructScan(
	table cat.Table,
	index cat.Index,
	params exec.ScanParams,
	reqOrdering exec.OutputOrdering,
) (exec.Node, error) {
	f.encodeOperator(scanOp)
	f.encodeDataSource(table.ID(), table.Name())
	f.encodeDataSource(index.ID(), index.Name())
	f.encodeScanParams(params)
	node, err := f.wrappedFactory.ConstructScan(
		table,
		index,
		params,
		reqOrdering,
	)
	return node, err
}

func (f *PlanGistFactory) ConstructValues(
	rows [][]tree.TypedExpr,
	columns colinfo.ResultColumns,
) (exec.Node, error) {
	f.encodeOperator(valuesOp)
	f.encodeRows(rows)
	f.encodeResultColumns(columns)
	node, err := f.wrappedFactory.ConstructValues(
		rows,
		columns,
	)
	return node, err
}

func (f *PlanGistFactory) ConstructFilter(
	input exec.Node,
	filter tree.TypedExpr,
	reqOrdering exec.OutputOrdering,
) (exec.Node, error) {
	f.encodeOperator(filterOp)
	node, err := f.wrappedFactory.ConstructFilter(
		input,
		filter,
		reqOrdering,
	)
	return node, err
}

func (f *PlanGistFactory) ConstructInvertedFilter(
	input exec.Node,
	invFilter *inverted.SpanExpression,
	preFiltererExpr tree.TypedExpr,
	preFiltererType *types.T,
	invColumn exec.NodeColumnOrdinal,
) (exec.Node, error) {
	f.encodeOperator(invertedFilterOp)
	node, err := f.wrappedFactory.ConstructInvertedFilter(
		input,
		invFilter,
		preFiltererExpr,
		preFiltererType,
		invColumn,
	)
	return node, err
}

func (f *PlanGistFactory) ConstructSimpleProject(
	input exec.Node,
	cols []exec.NodeColumnOrdinal,
	reqOrdering exec.OutputOrdering,
) (exec.Node, error) {
	f.encodeOperator(simpleProjectOp)
	f.encodeNodeColumnOrdinals(cols)
	node, err := f.wrappedFactory.ConstructSimpleProject(
		input,
		cols,
		reqOrdering,
	)
	return node, err
}

func (f *PlanGistFactory) ConstructSerializingProject(
	input exec.Node,
	cols []exec.NodeColumnOrdinal,
	colNames []string,
) (exec.Node, error) {
	f.encodeOperator(serializingProjectOp)
	f.encodeNodeColumnOrdinals(cols)
	node, err := f.wrappedFactory.ConstructSerializingProject(
		input,
		cols,
		colNames,
	)
	return node, err
}

func (f *PlanGistFactory) ConstructRender(
	input exec.Node,
	columns colinfo.ResultColumns,
	exprs tree.TypedExprs,
	reqOrdering exec.OutputOrdering,
) (exec.Node, error) {
	f.encodeOperator(renderOp)
	f.encodeResultColumns(columns)
	node, err := f.wrappedFactory.ConstructRender(
		input,
		columns,
		exprs,
		reqOrdering,
	)
	return node, err
}

func (f *PlanGistFactory) ConstructApplyJoin(
	joinType descpb.JoinType,
	left exec.Node,
	rightColumns colinfo.ResultColumns,
	onCond tree.TypedExpr,
	planRightSideFn exec.ApplyJoinPlanRightSideFn,
) (exec.Node, error) {
	f.encodeOperator(applyJoinOp)
	f.encodeByte(byte(joinType))
	f.encodeResultColumns(rightColumns)
	node, err := f.wrappedFactory.ConstructApplyJoin(
		joinType,
		left,
		rightColumns,
		onCond,
		planRightSideFn,
	)
	return node, err
}

func (f *PlanGistFactory) ConstructHashJoin(
	joinType descpb.JoinType,
	left exec.Node,
	right exec.Node,
	leftEqCols []exec.NodeColumnOrdinal,
	rightEqCols []exec.NodeColumnOrdinal,
	leftEqColsAreKey bool,
	rightEqColsAreKey bool,
	extraOnCond tree.TypedExpr,
) (exec.Node, error) {
	f.encodeOperator(hashJoinOp)
	f.encodeByte(byte(joinType))
	f.encodeNodeColumnOrdinals(leftEqCols)
	f.encodeNodeColumnOrdinals(rightEqCols)
	f.encodeBool(leftEqColsAreKey)
	f.encodeBool(rightEqColsAreKey)
	node, err := f.wrappedFactory.ConstructHashJoin(
		joinType,
		left,
		right,
		leftEqCols,
		rightEqCols,
		leftEqColsAreKey,
		rightEqColsAreKey,
		extraOnCond,
	)
	return node, err
}

func (f *PlanGistFactory) ConstructMergeJoin(
	joinType descpb.JoinType,
	left exec.Node,
	right exec.Node,
	onCond tree.TypedExpr,
	leftOrdering colinfo.ColumnOrdering,
	rightOrdering colinfo.ColumnOrdering,
	reqOrdering exec.OutputOrdering,
	leftEqColsAreKey bool,
	rightEqColsAreKey bool,
) (exec.Node, error) {
	f.encodeOperator(mergeJoinOp)
	f.encodeByte(byte(joinType))
	f.encodeColumnOrdering(leftOrdering)
	f.encodeColumnOrdering(rightOrdering)
	f.encodeBool(leftEqColsAreKey)
	f.encodeBool(rightEqColsAreKey)
	node, err := f.wrappedFactory.ConstructMergeJoin(
		joinType,
		left,
		right,
		onCond,
		leftOrdering,
		rightOrdering,
		reqOrdering,
		leftEqColsAreKey,
		rightEqColsAreKey,
	)
	return node, err
}

func (f *PlanGistFactory) ConstructGroupBy(
	input exec.Node,
	groupCols []exec.NodeColumnOrdinal,
	// If set, the input is guaranteed to have this ordering and a "streaming"
	// aggregation is performed (i.e. aggregation happens separately for each
	// distinct set of values on the set of columns in the ordering).
	groupColOrdering colinfo.ColumnOrdering,
	aggregations []exec.AggInfo,
	// If set, the output must have this ordering, but it is guaranteed that
	// ReqOrdering is a prefix of GroupColOrdering.
	reqOrdering exec.OutputOrdering,
	// The grouping column order type (Streaming, PartialStreaming, or
	// NoStreaming).
	groupingOrderType exec.GroupingOrderType,
) (exec.Node, error) {
	f.encodeOperator(groupByOp)
	f.encodeNodeColumnOrdinals(groupCols)
	f.encodeColumnOrdering(groupColOrdering)
	node, err := f.wrappedFactory.ConstructGroupBy(
		input,
		groupCols,
		groupColOrdering,
		aggregations,
		reqOrdering,
		groupingOrderType,
	)
	return node, err
}

func (f *PlanGistFactory) ConstructScalarGroupBy(
	input exec.Node,
	aggregations []exec.AggInfo,
) (exec.Node, error) {
	f.encodeOperator(scalarGroupByOp)
	node, err := f.wrappedFactory.ConstructScalarGroupBy(
		input,
		aggregations,
	)
	return node, err
}

func (f *PlanGistFactory) ConstructDistinct(
	input exec.Node,
	distinctCols exec.NodeColumnOrdinalSet,
	orderedCols exec.NodeColumnOrdinalSet,
	reqOrdering exec.OutputOrdering,
	nullsAreDistinct bool,
	errorOnDup string,
) (exec.Node, error) {
	f.encodeOperator(distinctOp)
	node, err := f.wrappedFactory.ConstructDistinct(
		input,
		distinctCols,
		orderedCols,
		reqOrdering,
		nullsAreDistinct,
		errorOnDup,
	)
	return node, err
}

func (f *PlanGistFactory) ConstructHashSetOp(
	typ tree.UnionType,
	all bool,
	left exec.Node,
	right exec.Node,
) (exec.Node, error) {
	f.encodeOperator(hashSetOpOp)
	f.encodeBool(all)
	node, err := f.wrappedFactory.ConstructHashSetOp(
		typ,
		all,
		left,
		right,
	)
	return node, err
}

func (f *PlanGistFactory) ConstructStreamingSetOp(
	typ tree.UnionType,
	all bool,
	left exec.Node,
	right exec.Node,
	streamingOrdering colinfo.ColumnOrdering,
	reqOrdering exec.OutputOrdering,
) (exec.Node, error) {
	f.encodeOperator(streamingSetOpOp)
	f.encodeBool(all)
	f.encodeColumnOrdering(streamingOrdering)
	node, err := f.wrappedFactory.ConstructStreamingSetOp(
		typ,
		all,
		left,
		right,
		streamingOrdering,
		reqOrdering,
	)
	return node, err
}

func (f *PlanGistFactory) ConstructUnionAll(
	left exec.Node,
	right exec.Node,
	reqOrdering exec.OutputOrdering,
	hardLimit uint64,
) (exec.Node, error) {
	f.encodeOperator(unionAllOp)
	node, err := f.wrappedFactory.ConstructUnionAll(
		left,
		right,
		reqOrdering,
		hardLimit,
	)
	return node, err
}

func (f *PlanGistFactory) ConstructSort(
	input exec.Node,
	ordering exec.OutputOrdering,
	alreadyOrderedPrefix int,
) (exec.Node, error) {
	f.encodeOperator(sortOp)
	node, err := f.wrappedFactory.ConstructSort(
		input,
		ordering,
		alreadyOrderedPrefix,
	)
	return node, err
}

func (f *PlanGistFactory) ConstructOrdinality(
	input exec.Node,
	colName string,
) (exec.Node, error) {
	f.encodeOperator(ordinalityOp)
	node, err := f.wrappedFactory.ConstructOrdinality(
		input,
		colName,
	)
	return node, err
}

func (f *PlanGistFactory) ConstructIndexJoin(
	input exec.Node,
	table cat.Table,
	keyCols []exec.NodeColumnOrdinal,
	tableCols exec.TableColumnOrdinalSet,
	reqOrdering exec.OutputOrdering,
) (exec.Node, error) {
	f.encodeOperator(indexJoinOp)
	f.encodeDataSource(table.ID(), table.Name())
	f.encodeNodeColumnOrdinals(keyCols)
	node, err := f.wrappedFactory.ConstructIndexJoin(
		input,
		table,
		keyCols,
		tableCols,
		reqOrdering,
	)
	return node, err
}

func (f *PlanGistFactory) ConstructLookupJoin(
	joinType descpb.JoinType,
	input exec.Node,
	table cat.Table,
	index cat.Index,
	eqCols []exec.NodeColumnOrdinal,
	eqColsAreKey bool,
	lookupExpr tree.TypedExpr,
	remoteLookupExpr tree.TypedExpr,
	lookupCols exec.TableColumnOrdinalSet,
	onCond tree.TypedExpr,
	isSecondJoinInPairedJoiner bool,
	reqOrdering exec.OutputOrdering,
	locking *tree.LockingItem,
) (exec.Node, error) {
	f.encodeOperator(lookupJoinOp)
	f.encodeByte(byte(joinType))
	f.encodeDataSource(table.ID(), table.Name())
	f.encodeDataSource(index.ID(), index.Name())
	f.encodeNodeColumnOrdinals(eqCols)
	f.encodeBool(eqColsAreKey)
	node, err := f.wrappedFactory.ConstructLookupJoin(
		joinType,
		input,
		table,
		index,
		eqCols,
		eqColsAreKey,
		lookupExpr,
		remoteLookupExpr,
		lookupCols,
		onCond,
		isSecondJoinInPairedJoiner,
		reqOrdering,
		locking,
	)
	return node, err
}

func (f *PlanGistFactory) ConstructInvertedJoin(
	joinType descpb.JoinType,
	invertedExpr tree.TypedExpr,
	input exec.Node,
	table cat.Table,
	index cat.Index,
	prefixEqCols []exec.NodeColumnOrdinal,
	lookupCols exec.TableColumnOrdinalSet,
	onCond tree.TypedExpr,
	isFirstJoinInPairedJoiner bool,
	reqOrdering exec.OutputOrdering,
) (exec.Node, error) {
	f.encodeOperator(invertedJoinOp)
	f.encodeByte(byte(joinType))
	f.encodeDataSource(table.ID(), table.Name())
	f.encodeDataSource(index.ID(), index.Name())
	f.encodeNodeColumnOrdinals(prefixEqCols)
	node, err := f.wrappedFactory.ConstructInvertedJoin(
		joinType,
		invertedExpr,
		input,
		table,
		index,
		prefixEqCols,
		lookupCols,
		onCond,
		isFirstJoinInPairedJoiner,
		reqOrdering,
	)
	return node, err
}

func (f *PlanGistFactory) ConstructZigzagJoin(
	// Left table and index.
	leftTable cat.Table,
	leftIndex cat.Index,
	// leftCols are the columns that are scanned from the left index.
	leftCols exec.TableColumnOrdinalSet,
	// leftFixedVals contains values for the fixed columns (a prefix of the
	// index columns).
	leftFixedVals []tree.TypedExpr,
	// leftEqCols are the left table columns that have equality constraints,
	// corresponding 1-1 to RightEqCols.
	leftEqCols []exec.TableColumnOrdinal,
	// Right table and index.
	rightTable cat.Table,
	rightIndex cat.Index,
	// rightCols are the columns that are scanned from the right index.
	rightCols exec.TableColumnOrdinalSet,
	// rightFixedVals contains values for the fixed columns (a prefix of the
	// index columns).
	rightFixedVals []tree.TypedExpr,
	// rightEqCols are the right table columns that have equality constraints,
	// corresponding 1-1 to LeftEqCols.
	rightEqCols []exec.TableColumnOrdinal,
	// onCond is an extra filter that is evaluated on the results.
	// TODO(radu): remove this (it can be a separate Select).
	onCond tree.TypedExpr,
	reqOrdering exec.OutputOrdering,
) (exec.Node, error) {
	f.encodeOperator(zigzagJoinOp)
	f.encodeDataSource(leftTable.ID(), leftTable.Name())
	f.encodeDataSource(leftIndex.ID(), leftIndex.Name())
	f.encodeDataSource(rightTable.ID(), rightTable.Name())
	f.encodeDataSource(rightIndex.ID(), rightIndex.Name())
	node, err := f.wrappedFactory.ConstructZigzagJoin(
		leftTable,
		leftIndex,
		leftCols,
		leftFixedVals,
		leftEqCols,
		rightTable,
		rightIndex,
		rightCols,
		rightFixedVals,
		rightEqCols,
		onCond,
		reqOrdering,
	)
	return node, err
}

func (f *PlanGistFactory) ConstructLimit(
	input exec.Node,
	limit tree.TypedExpr,
	offset tree.TypedExpr,
) (exec.Node, error) {
	f.encodeOperator(limitOp)
	node, err := f.wrappedFactory.ConstructLimit(
		input,
		limit,
		offset,
	)
	return node, err
}

func (f *PlanGistFactory) ConstructTopK(
	input exec.Node,
	k int64,
	ordering exec.OutputOrdering,
	alreadyOrderedPrefix int,
) (exec.Node, error) {
	f.encodeOperator(topKOp)
	node, err := f.wrappedFactory.ConstructTopK(
		input,
		k,
		ordering,
		alreadyOrderedPrefix,
	)
	return node, err
}

func (f *PlanGistFactory) ConstructMax1Row(
	input exec.Node,
	errorText string,
) (exec.Node, error) {
	f.encodeOperator(max1RowOp)
	node, err := f.wrappedFactory.ConstructMax1Row(
		input,
		errorText,
	)
	return node, err
}

func (f *PlanGistFactory) ConstructProjectSet(
	input exec.Node,
	exprs tree.TypedExprs,
	zipCols colinfo.ResultColumns,
	numColsPerGen []int,
) (exec.Node, error) {
	f.encodeOperator(projectSetOp)
	f.encodeResultColumns(zipCols)
	node, err := f.wrappedFactory.ConstructProjectSet(
		input,
		exprs,
		zipCols,
		numColsPerGen,
	)
	return node, err
}

func (f *PlanGistFactory) ConstructWindow(
	input exec.Node,
	window exec.WindowInfo,
) (exec.Node, error) {
	f.encodeOperator(windowOp)
	node, err := f.wrappedFactory.ConstructWindow(
		input,
		window,
	)
	return node, err
}

func (f *PlanGistFactory) ConstructExplainOpt(
	plan string,
	envOpts exec.ExplainEnvData,
) (exec.Node, error) {
	f.encodeOperator(explainOptOp)
	node, err := f.wrappedFactory.ConstructExplainOpt(
		plan,
		envOpts,
	)
	return node, err
}

func (f *PlanGistFactory) ConstructExplain(
	options *tree.ExplainOptions,
	stmtType tree.StatementReturnType,
	buildFn exec.BuildPlanForExplainFn,
) (exec.Node, error) {
	f.encodeOperator(explainOp)
	node, err := f.wrappedFactory.ConstructExplain(
		options,
		stmtType,
		buildFn,
	)
	return node, err
}

func (f *PlanGistFactory) ConstructShowTrace(
	typ tree.ShowTraceType,
	compact bool,
) (exec.Node, error) {
	f.encodeOperator(showTraceOp)
	node, err := f.wrappedFactory.ConstructShowTrace(
		typ,
		compact,
	)
	return node, err
}

func (f *PlanGistFactory) ConstructInsert(
	input exec.Node,
	table cat.Table,
	arbiterIndexes cat.IndexOrdinals,
	arbiterConstraints cat.UniqueOrdinals,
	insertCols exec.TableColumnOrdinalSet,
	returnCols exec.TableColumnOrdinalSet,
	checkCols exec.CheckOrdinalSet,
	// If set, the operator will commit the transaction as part of its execution.
	// This is false when executing inside an explicit transaction, or there are
	// multiple mutations in a statement, or the output of the mutation is
	// processed through side-effecting expressions.
	autoCommit bool,
) (exec.Node, error) {
	f.encodeOperator(insertOp)
	f.encodeDataSource(table.ID(), table.Name())
	f.encodeBool(autoCommit)
	node, err := f.wrappedFactory.ConstructInsert(
		input,
		table,
		arbiterIndexes,
		arbiterConstraints,
		insertCols,
		returnCols,
		checkCols,
		autoCommit,
	)
	return node, err
}

func (f *PlanGistFactory) ConstructInsertFastPath(
	rows [][]tree.TypedExpr,
	table cat.Table,
	insertCols exec.TableColumnOrdinalSet,
	returnCols exec.TableColumnOrdinalSet,
	checkCols exec.CheckOrdinalSet,
	fkChecks []exec.InsertFastPathFKCheck,
	// If set, the operator will commit the transaction as part of its execution.
	// This is false when executing inside an explicit transaction.
	autoCommit bool,
) (exec.Node, error) {
	f.encodeOperator(insertFastPathOp)
	f.encodeRows(rows)
	f.encodeDataSource(table.ID(), table.Name())
	f.encodeBool(autoCommit)
	node, err := f.wrappedFactory.ConstructInsertFastPath(
		rows,
		table,
		insertCols,
		returnCols,
		checkCols,
		fkChecks,
		autoCommit,
	)
	return node, err
}

func (f *PlanGistFactory) ConstructUpdate(
	input exec.Node,
	table cat.Table,
	fetchCols exec.TableColumnOrdinalSet,
	updateCols exec.TableColumnOrdinalSet,
	returnCols exec.TableColumnOrdinalSet,
	checks exec.CheckOrdinalSet,
	passthrough colinfo.ResultColumns,
	// If set, the operator will commit the transaction as part of its execution.
	autoCommit bool,
) (exec.Node, error) {
	f.encodeOperator(updateOp)
	f.encodeDataSource(table.ID(), table.Name())
	f.encodeResultColumns(passthrough)
	f.encodeBool(autoCommit)
	node, err := f.wrappedFactory.ConstructUpdate(
		input,
		table,
		fetchCols,
		updateCols,
		returnCols,
		checks,
		passthrough,
		autoCommit,
	)
	return node, err
}

func (f *PlanGistFactory) ConstructUpsert(
	input exec.Node,
	table cat.Table,
	arbiterIndexes cat.IndexOrdinals,
	arbiterConstraints cat.UniqueOrdinals,
	canaryCol exec.NodeColumnOrdinal,
	insertCols exec.TableColumnOrdinalSet,
	fetchCols exec.TableColumnOrdinalSet,
	updateCols exec.TableColumnOrdinalSet,
	returnCols exec.TableColumnOrdinalSet,
	checks exec.CheckOrdinalSet,
	// If set, the operator will commit the transaction as part of its execution.
	// This is false when executing inside an explicit transaction, or there are
	// multiple mutations in a statement, or the output of the mutation is
	// processed through side-effecting expressions.
	autoCommit bool,
) (exec.Node, error) {
	f.encodeOperator(upsertOp)
	f.encodeDataSource(table.ID(), table.Name())
	f.encodeBool(autoCommit)
	node, err := f.wrappedFactory.ConstructUpsert(
		input,
		table,
		arbiterIndexes,
		arbiterConstraints,
		canaryCol,
		insertCols,
		fetchCols,
		updateCols,
		returnCols,
		checks,
		autoCommit,
	)
	return node, err
}

func (f *PlanGistFactory) ConstructDelete(
	input exec.Node,
	table cat.Table,
	fetchCols exec.TableColumnOrdinalSet,
	returnCols exec.TableColumnOrdinalSet,
	// If set, the operator will commit the transaction as part of its execution.
	// This is false when executing inside an explicit transaction, or there are
	// multiple mutations in a statement, or the output of the mutation is
	// processed through side-effecting expressions.
	autoCommit bool,
) (exec.Node, error) {
	f.encodeOperator(deleteOp)
	f.encodeDataSource(table.ID(), table.Name())
	f.encodeBool(autoCommit)
	node, err := f.wrappedFactory.ConstructDelete(
		input,
		table,
		fetchCols,
		returnCols,
		autoCommit,
	)
	return node, err
}

func (f *PlanGistFactory) ConstructDeleteRange(
	table cat.Table,
	needed exec.TableColumnOrdinalSet,
	indexConstraint *constraint.Constraint,
	// If set, the operator will commit the transaction as part of its execution.
	// This is false when executing inside an explicit transaction, or there are
	// multiple mutations in a statement, or the output of the mutation is
	// processed through side-effecting expressions, or the operation might
	// process too many rows.
	autoCommit bool,
) (exec.Node, error) {
	f.encodeOperator(deleteRangeOp)
	f.encodeDataSource(table.ID(), table.Name())
	f.encodeBool(autoCommit)
	node, err := f.wrappedFactory.ConstructDeleteRange(
		table,
		needed,
		indexConstraint,
		autoCommit,
	)
	return node, err
}

func (f *PlanGistFactory) ConstructCreateTable(
	schema cat.Schema,
	ct *tree.CreateTable,
) (exec.Node, error) {
	f.encodeOperator(createTableOp)
	f.encodeID(schema.ID())
	node, err := f.wrappedFactory.ConstructCreateTable(
		schema,
		ct,
	)
	return node, err
}

func (f *PlanGistFactory) ConstructCreateTableAs(
	input exec.Node,
	schema cat.Schema,
	ct *tree.CreateTable,
) (exec.Node, error) {
	f.encodeOperator(createTableAsOp)
	f.encodeID(schema.ID())
	node, err := f.wrappedFactory.ConstructCreateTableAs(
		input,
		schema,
		ct,
	)
	return node, err
}

func (f *PlanGistFactory) ConstructCreateView(
	schema cat.Schema,
	viewName *cat.DataSourceName,
	ifNotExists bool,
	replace bool,
	persistence tree.Persistence,
	materialized bool,
	viewQuery string,
	columns colinfo.ResultColumns,
	deps opt.ViewDeps,
	typeDeps opt.ViewTypeDeps,
) (exec.Node, error) {
	f.encodeOperator(createViewOp)
	f.encodeID(schema.ID())
	f.encodeResultColumns(columns)
	node, err := f.wrappedFactory.ConstructCreateView(
		schema,
		viewName,
		ifNotExists,
		replace,
		persistence,
		materialized,
		viewQuery,
		columns,
		deps,
		typeDeps,
	)
	return node, err
}

func (f *PlanGistFactory) ConstructSequenceSelect(
	sequence cat.Sequence,
) (exec.Node, error) {
	f.encodeOperator(sequenceSelectOp)
	node, err := f.wrappedFactory.ConstructSequenceSelect(
		sequence,
	)
	return node, err
}

func (f *PlanGistFactory) ConstructSaveTable(
	input exec.Node,
	table *cat.DataSourceName,
	colNames []string,
) (exec.Node, error) {
	f.encodeOperator(saveTableOp)
	node, err := f.wrappedFactory.ConstructSaveTable(
		input,
		table,
		colNames,
	)
	return node, err
}

func (f *PlanGistFactory) ConstructErrorIfRows(
	input exec.Node,
	// mkErr is used to create the error; it is passed an input row.
	mkErr exec.MkErrFn,
) (exec.Node, error) {
	f.encodeOperator(errorIfRowsOp)
	node, err := f.wrappedFactory.ConstructErrorIfRows(
		input,
		mkErr,
	)
	return node, err
}

func (f *PlanGistFactory) ConstructOpaque(
	metadata opt.OpaqueMetadata,
) (exec.Node, error) {
	f.encodeOperator(opaqueOp)
	node, err := f.wrappedFactory.ConstructOpaque(
		metadata,
	)
	return node, err
}

func (f *PlanGistFactory) ConstructAlterTableSplit(
	index cat.Index,
	input exec.Node,
	expiration tree.TypedExpr,
) (exec.Node, error) {
	f.encodeOperator(alterTableSplitOp)
	f.encodeDataSource(index.Table().ID(), index.Table().Name())
	f.encodeDataSource(index.ID(), index.Name())
	node, err := f.wrappedFactory.ConstructAlterTableSplit(
		index,
		input,
		expiration,
	)
	return node, err
}

func (f *PlanGistFactory) ConstructAlterTableUnsplit(
	index cat.Index,
	input exec.Node,
) (exec.Node, error) {
	f.encodeOperator(alterTableUnsplitOp)
	f.encodeDataSource(index.Table().ID(), index.Table().Name())
	f.encodeDataSource(index.ID(), index.Name())
	node, err := f.wrappedFactory.ConstructAlterTableUnsplit(
		index,
		input,
	)
	return node, err
}

func (f *PlanGistFactory) ConstructAlterTableUnsplitAll(
	index cat.Index,
) (exec.Node, error) {
	f.encodeOperator(alterTableUnsplitAllOp)
	f.encodeDataSource(index.Table().ID(), index.Table().Name())
	f.encodeDataSource(index.ID(), index.Name())
	node, err := f.wrappedFactory.ConstructAlterTableUnsplitAll(
		index,
	)
	return node, err
}

func (f *PlanGistFactory) ConstructAlterTableRelocate(
	index cat.Index,
	input exec.Node,
	subjectReplicas tree.RelocateSubject,
) (exec.Node, error) {
	f.encodeOperator(alterTableRelocateOp)
	f.encodeDataSource(index.Table().ID(), index.Table().Name())
	f.encodeDataSource(index.ID(), index.Name())
	node, err := f.wrappedFactory.ConstructAlterTableRelocate(
		index,
		input,
		subjectReplicas,
	)
	return node, err
}

func (f *PlanGistFactory) ConstructBuffer(
	input exec.Node,
	label string,
) (exec.Node, error) {
	f.encodeOperator(bufferOp)
	node, err := f.wrappedFactory.ConstructBuffer(
		input,
		label,
	)
	return node, err
}

func (f *PlanGistFactory) ConstructScanBuffer(
	ref exec.Node,
	label string,
) (exec.Node, error) {
	f.encodeOperator(scanBufferOp)
	node, err := f.wrappedFactory.ConstructScanBuffer(
		ref,
		label,
	)
	return node, err
}

func (f *PlanGistFactory) ConstructRecursiveCTE(
	initial exec.Node,
	fn exec.RecursiveCTEIterationFn,
	label string,
	deduplicate bool,
) (exec.Node, error) {
	f.encodeOperator(recursiveCTEOp)
	node, err := f.wrappedFactory.ConstructRecursiveCTE(
		initial,
		fn,
		label,
		deduplicate,
	)
	return node, err
}

func (f *PlanGistFactory) ConstructControlJobs(
	command tree.JobCommand,
	input exec.Node,
	reason tree.TypedExpr,
) (exec.Node, error) {
	f.encodeOperator(controlJobsOp)
	node, err := f.wrappedFactory.ConstructControlJobs(
		command,
		input,
		reason,
	)
	return node, err
}

func (f *PlanGistFactory) ConstructControlSchedules(
	command tree.ScheduleCommand,
	input exec.Node,
) (exec.Node, error) {
	f.encodeOperator(controlSchedulesOp)
	node, err := f.wrappedFactory.ConstructControlSchedules(
		command,
		input,
	)
	return node, err
}

func (f *PlanGistFactory) ConstructCancelQueries(
	input exec.Node,
	ifExists bool,
) (exec.Node, error) {
	f.encodeOperator(cancelQueriesOp)
	node, err := f.wrappedFactory.ConstructCancelQueries(
		input,
		ifExists,
	)
	return node, err
}

func (f *PlanGistFactory) ConstructCancelSessions(
	input exec.Node,
	ifExists bool,
) (exec.Node, error) {
	f.encodeOperator(cancelSessionsOp)
	node, err := f.wrappedFactory.ConstructCancelSessions(
		input,
		ifExists,
	)
	return node, err
}

func (f *PlanGistFactory) ConstructCreateStatistics(
	cs *tree.CreateStats,
) (exec.Node, error) {
	f.encodeOperator(createStatisticsOp)
	node, err := f.wrappedFactory.ConstructCreateStatistics(
		cs,
	)
	return node, err
}

func (f *PlanGistFactory) ConstructExport(
	input exec.Node,
	fileName tree.TypedExpr,
	fileFormat string,
	options []exec.KVOption,
	notNullCols exec.NodeColumnOrdinalSet,
) (exec.Node, error) {
	f.encodeOperator(exportOp)
	node, err := f.wrappedFactory.ConstructExport(
		input,
		fileName,
		fileFormat,
		options,
		notNullCols,
	)
	return node, err
}

func (f *PlanGistFactory) ConstructAlterRangeRelocate(
	input exec.Node,
	subjectReplicas tree.RelocateSubject,
	toStoreID tree.TypedExpr,
	fromStoreID tree.TypedExpr,
) (exec.Node, error) {
	f.encodeOperator(alterRangeRelocateOp)
	node, err := f.wrappedFactory.ConstructAlterRangeRelocate(
		input,
		subjectReplicas,
		toStoreID,
		fromStoreID,
	)
	return node, err
}

func (f *PlanGistFactory) decodeOperatorBody(op execOperator) (*Node, error) {
	var _n *Node
	var reqOrdering exec.OutputOrdering
	var err error
	var tbl cat.Table
	switch op {
	case scanOp:
		var args scanArgs
		args.Table = f.decodeTable()
		tbl = args.Table
		args.Index = f.decodeIndex(tbl)
		args.Params = f.decodeScanParams()
		_n, err = newNode(op, &args, reqOrdering)
	case valuesOp:
		var args valuesArgs
		args.Rows = f.decodeRows()
		args.Columns = f.decodeResultColumns()
		_n, err = newNode(op, &args, reqOrdering)
	case filterOp:
		var args filterArgs
		args.Input = f.popChild()
		_n, err = newNode(op, &args, reqOrdering, args.Input)
	case invertedFilterOp:
		var args invertedFilterArgs
		args.Input = f.popChild()
		_n, err = newNode(op, &args, reqOrdering, args.Input)
	case simpleProjectOp:
		var args simpleProjectArgs
		args.Cols = f.decodeNodeColumnOrdinals()
		args.Input = f.popChild()
		_n, err = newNode(op, &args, reqOrdering, args.Input)
	case serializingProjectOp:
		var args serializingProjectArgs
		args.Cols = f.decodeNodeColumnOrdinals()
		args.Input = f.popChild()
		_n, err = newNode(op, &args, reqOrdering, args.Input)
	case renderOp:
		var args renderArgs
		args.Columns = f.decodeResultColumns()
		args.Input = f.popChild()
		_n, err = newNode(op, &args, reqOrdering, args.Input)
	case applyJoinOp:
		var args applyJoinArgs
		args.JoinType = f.decodeJoinType()
		args.RightColumns = f.decodeResultColumns()
		args.Left = f.popChild()
		_n, err = newNode(op, &args, reqOrdering, args.Left)
	case hashJoinOp:
		var args hashJoinArgs
		args.JoinType = f.decodeJoinType()
		args.LeftEqCols = f.decodeNodeColumnOrdinals()
		args.RightEqCols = f.decodeNodeColumnOrdinals()
		args.LeftEqColsAreKey = f.decodeBool()
		args.RightEqColsAreKey = f.decodeBool()
		args.Right = f.popChild()
		args.Left = f.popChild()
		_n, err = newNode(op, &args, reqOrdering, args.Left, args.Right)
	case mergeJoinOp:
		var args mergeJoinArgs
		args.JoinType = f.decodeJoinType()
		args.LeftOrdering = f.decodeColumnOrdering()
		args.RightOrdering = f.decodeColumnOrdering()
		args.LeftEqColsAreKey = f.decodeBool()
		args.RightEqColsAreKey = f.decodeBool()
		args.Right = f.popChild()
		args.Left = f.popChild()
		_n, err = newNode(op, &args, reqOrdering, args.Left, args.Right)
	case groupByOp:
		var args groupByArgs
		args.GroupCols = f.decodeNodeColumnOrdinals()
		args.GroupColOrdering = f.decodeColumnOrdering()
		args.Input = f.popChild()
		_n, err = newNode(op, &args, reqOrdering, args.Input)
	case scalarGroupByOp:
		var args scalarGroupByArgs
		args.Input = f.popChild()
		_n, err = newNode(op, &args, reqOrdering, args.Input)
	case distinctOp:
		var args distinctArgs
		args.Input = f.popChild()
		_n, err = newNode(op, &args, reqOrdering, args.Input)
	case hashSetOpOp:
		var args hashSetOpArgs
		args.All = f.decodeBool()
		args.Right = f.popChild()
		args.Left = f.popChild()
		_n, err = newNode(op, &args, reqOrdering, args.Left, args.Right)
	case streamingSetOpOp:
		var args streamingSetOpArgs
		args.All = f.decodeBool()
		args.StreamingOrdering = f.decodeColumnOrdering()
		args.Right = f.popChild()
		args.Left = f.popChild()
		_n, err = newNode(op, &args, reqOrdering, args.Left, args.Right)
	case unionAllOp:
		var args unionAllArgs
		args.Right = f.popChild()
		args.Left = f.popChild()
		_n, err = newNode(op, &args, reqOrdering, args.Left, args.Right)
	case sortOp:
		var args sortArgs
		args.Input = f.popChild()
		_n, err = newNode(op, &args, reqOrdering, args.Input)
	case ordinalityOp:
		var args ordinalityArgs
		args.Input = f.popChild()
		_n, err = newNode(op, &args, reqOrdering, args.Input)
	case indexJoinOp:
		var args indexJoinArgs
		args.Table = f.decodeTable()
		args.KeyCols = f.decodeNodeColumnOrdinals()
		args.Input = f.popChild()
		_n, err = newNode(op, &args, reqOrdering, args.Input)
	case lookupJoinOp:
		var args lookupJoinArgs
		args.JoinType = f.decodeJoinType()
		args.Table = f.decodeTable()
		tbl = args.Table
		args.Index = f.decodeIndex(tbl)
		args.EqCols = f.decodeNodeColumnOrdinals()
		args.EqColsAreKey = f.decodeBool()
		args.Input = f.popChild()
		_n, err = newNode(op, &args, reqOrdering, args.Input)
	case invertedJoinOp:
		var args invertedJoinArgs
		args.JoinType = f.decodeJoinType()
		args.Table = f.decodeTable()
		tbl = args.Table
		args.Index = f.decodeIndex(tbl)
		args.PrefixEqCols = f.decodeNodeColumnOrdinals()
		args.Input = f.popChild()
		_n, err = newNode(op, &args, reqOrdering, args.Input)
	case zigzagJoinOp:
		var args zigzagJoinArgs
		args.LeftTable = f.decodeTable()
		tbl = args.LeftTable
		args.LeftIndex = f.decodeIndex(tbl)
		args.RightTable = f.decodeTable()
		tbl = args.RightTable
		args.RightIndex = f.decodeIndex(tbl)
		_n, err = newNode(op, &args, reqOrdering)
	case limitOp:
		var args limitArgs
		args.Input = f.popChild()
		_n, err = newNode(op, &args, reqOrdering, args.Input)
	case topKOp:
		var args topKArgs
		args.Input = f.popChild()
		_n, err = newNode(op, &args, reqOrdering, args.Input)
	case max1RowOp:
		var args max1RowArgs
		args.Input = f.popChild()
		_n, err = newNode(op, &args, reqOrdering, args.Input)
	case projectSetOp:
		var args projectSetArgs
		args.ZipCols = f.decodeResultColumns()
		args.Input = f.popChild()
		_n, err = newNode(op, &args, reqOrdering, args.Input)
	case windowOp:
		var args windowArgs
		args.Input = f.popChild()
		_n, err = newNode(op, &args, reqOrdering, args.Input)
	case explainOptOp:
		var args explainOptArgs
		_n, err = newNode(op, &args, reqOrdering)
	case explainOp:
		var args explainArgs
		_n, err = newNode(op, &args, reqOrdering)
	case showTraceOp:
		var args showTraceArgs
		_n, err = newNode(op, &args, reqOrdering)
	case insertOp:
		var args insertArgs
		args.Table = f.decodeTable()
		args.AutoCommit = f.decodeBool()
		args.Input = f.popChild()
		_n, err = newNode(op, &args, reqOrdering, args.Input)
	case insertFastPathOp:
		var args insertFastPathArgs
		args.Rows = f.decodeRows()
		args.Table = f.decodeTable()
		args.AutoCommit = f.decodeBool()
		_n, err = newNode(op, &args, reqOrdering)
	case updateOp:
		var args updateArgs
		args.Table = f.decodeTable()
		args.Passthrough = f.decodeResultColumns()
		args.AutoCommit = f.decodeBool()
		args.Input = f.popChild()
		_n, err = newNode(op, &args, reqOrdering, args.Input)
	case upsertOp:
		var args upsertArgs
		args.Table = f.decodeTable()
		args.AutoCommit = f.decodeBool()
		args.Input = f.popChild()
		_n, err = newNode(op, &args, reqOrdering, args.Input)
	case deleteOp:
		var args deleteArgs
		args.Table = f.decodeTable()
		args.AutoCommit = f.decodeBool()
		args.Input = f.popChild()
		_n, err = newNode(op, &args, reqOrdering, args.Input)
	case deleteRangeOp:
		var args deleteRangeArgs
		args.Table = f.decodeTable()
		args.AutoCommit = f.decodeBool()
		_n, err = newNode(op, &args, reqOrdering)
	case createTableOp:
		var args createTableArgs
		args.Schema = f.decodeSchema()
		_n, err = newNode(op, &args, reqOrdering)
	case createTableAsOp:
		var args createTableAsArgs
		args.Schema = f.decodeSchema()
		args.Input = f.popChild()
		_n, err = newNode(op, &args, reqOrdering, args.Input)
	case createViewOp:
		var args createViewArgs
		args.Schema = f.decodeSchema()
		args.Columns = f.decodeResultColumns()
		_n, err = newNode(op, &args, reqOrdering)
	case sequenceSelectOp:
		var args sequenceSelectArgs
		_n, err = newNode(op, &args, reqOrdering)
	case saveTableOp:
		var args saveTableArgs
		args.Input = f.popChild()
		_n, err = newNode(op, &args, reqOrdering, args.Input)
	case errorIfRowsOp:
		var args errorIfRowsArgs
		args.Input = f.popChild()
		_n, err = newNode(op, &args, reqOrdering, args.Input)
	case opaqueOp:
		var args opaqueArgs
		_n, err = newNode(op, &args, reqOrdering)
	case alterTableSplitOp:
		var args alterTableSplitArgs
		tbl := f.decodeTable()
		args.Index = f.decodeIndex(tbl)
		args.Input = f.popChild()
		_n, err = newNode(op, &args, reqOrdering, args.Input)
	case alterTableUnsplitOp:
		var args alterTableUnsplitArgs
		tbl := f.decodeTable()
		args.Index = f.decodeIndex(tbl)
		args.Input = f.popChild()
		_n, err = newNode(op, &args, reqOrdering, args.Input)
	case alterTableUnsplitAllOp:
		var args alterTableUnsplitAllArgs
		tbl := f.decodeTable()
		args.Index = f.decodeIndex(tbl)
		_n, err = newNode(op, &args, reqOrdering)
	case alterTableRelocateOp:
		var args alterTableRelocateArgs
		tbl := f.decodeTable()
		args.Index = f.decodeIndex(tbl)
		args.input = f.popChild()
		_n, err = newNode(op, &args, reqOrdering, args.input)
	case bufferOp:
		var args bufferArgs
		args.Input = f.popChild()
		_n, err = newNode(op, &args, reqOrdering, args.Input)
	case scanBufferOp:
		var args scanBufferArgs
		_n, err = newNode(op, &args, reqOrdering)
	case recursiveCTEOp:
		var args recursiveCTEArgs
		args.Initial = f.popChild()
		_n, err = newNode(op, &args, reqOrdering, args.Initial)
	case controlJobsOp:
		var args controlJobsArgs
		args.input = f.popChild()
		_n, err = newNode(op, &args, reqOrdering, args.input)
	case controlSchedulesOp:
		var args controlSchedulesArgs
		args.input = f.popChild()
		_n, err = newNode(op, &args, reqOrdering, args.input)
	case cancelQueriesOp:
		var args cancelQueriesArgs
		args.Input = f.popChild()
		_n, err = newNode(op, &args, reqOrdering, args.Input)
	case cancelSessionsOp:
		var args cancelSessionsArgs
		args.Input = f.popChild()
		_n, err = newNode(op, &args, reqOrdering, args.Input)
	case createStatisticsOp:
		var args createStatisticsArgs
		_n, err = newNode(op, &args, reqOrdering)
	case exportOp:
		var args exportArgs
		args.Input = f.popChild()
		_n, err = newNode(op, &args, reqOrdering, args.Input)
	case alterRangeRelocateOp:
		var args alterRangeRelocateArgs
		args.input = f.popChild()
		_n, err = newNode(op, &args, reqOrdering, args.input)
	default:
		return nil, errors.Newf("invalid op: %d", op)
	}
	return _n, err
}
