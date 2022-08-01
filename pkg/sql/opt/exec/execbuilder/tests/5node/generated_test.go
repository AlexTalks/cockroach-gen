// Copyright 2022 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

// Code generated by generate-logictest, DO NOT EDIT.

package test5node

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/cockroachdb/cockroach/pkg/build/bazel"
	"github.com/cockroachdb/cockroach/pkg/security/securityassets"
	"github.com/cockroachdb/cockroach/pkg/security/securitytest"
	"github.com/cockroachdb/cockroach/pkg/server"
	"github.com/cockroachdb/cockroach/pkg/sql"
	"github.com/cockroachdb/cockroach/pkg/sql/logictest"
	"github.com/cockroachdb/cockroach/pkg/testutils/serverutils"
	"github.com/cockroachdb/cockroach/pkg/testutils/skip"
	"github.com/cockroachdb/cockroach/pkg/testutils/testcluster"
	"github.com/cockroachdb/cockroach/pkg/util/leaktest"
	"github.com/cockroachdb/cockroach/pkg/util/randutil"
)

const configIdx = 7

var execBuildLogicTestDir string

func init() {
	if bazel.BuiltWithBazel() {
		var err error
		execBuildLogicTestDir, err = bazel.Runfile("pkg/sql/opt/exec/execbuilder/testdata")
		if err != nil {
			panic(err)
		}
	} else {
		execBuildLogicTestDir = "../../../../../../sql/opt/exec/execbuilder/testdata"
	}
}

func TestMain(m *testing.M) {
	securityassets.SetLoader(securitytest.EmbeddedAssets)
	randutil.SeedForTests()
	serverutils.InitTestServerFactory(server.TestServerFactory)
	serverutils.InitTestClusterFactory(testcluster.TestClusterFactory)
	os.Exit(m.Run())
}

func runExecBuildLogicTest(t *testing.T, file string) {
	defer sql.TestingOverrideExplainEnvVersion("CockroachDB execbuilder test version")()
	skip.UnderDeadlock(t, "times out and/or hangs")
	serverArgs := logictest.TestServerArgs{
		DisableWorkmemRandomization: true,
		ForceProductionValues:       true,
	}
	logictest.RunLogicTest(t, serverArgs, configIdx, filepath.Join(execBuildLogicTestDir, file))
}

func TestExecBuild_dist_union(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runExecBuildLogicTest(t, "dist_union")
}

func TestExecBuild_dist_vectorize(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runExecBuildLogicTest(t, "dist_vectorize")
}

func TestExecBuild_distsql_agg(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runExecBuildLogicTest(t, "distsql_agg")
}

func TestExecBuild_distsql_auto_mode(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runExecBuildLogicTest(t, "distsql_auto_mode")
}

func TestExecBuild_distsql_distinct_on(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runExecBuildLogicTest(t, "distsql_distinct_on")
}

func TestExecBuild_distsql_indexjoin(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runExecBuildLogicTest(t, "distsql_indexjoin")
}

func TestExecBuild_distsql_inverted_index(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runExecBuildLogicTest(t, "distsql_inverted_index")
}

func TestExecBuild_distsql_join(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runExecBuildLogicTest(t, "distsql_join")
}

func TestExecBuild_distsql_merge_join(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runExecBuildLogicTest(t, "distsql_merge_join")
}

func TestExecBuild_distsql_misc(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runExecBuildLogicTest(t, "distsql_misc")
}

func TestExecBuild_distsql_numtables(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runExecBuildLogicTest(t, "distsql_numtables")
}

func TestExecBuild_distsql_ordinality(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runExecBuildLogicTest(t, "distsql_ordinality")
}

func TestExecBuild_distsql_single_flow(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runExecBuildLogicTest(t, "distsql_single_flow")
}

func TestExecBuild_distsql_tighten_spans(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runExecBuildLogicTest(t, "distsql_tighten_spans")
}

func TestExecBuild_distsql_union(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runExecBuildLogicTest(t, "distsql_union")
}

func TestExecBuild_distsql_window(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runExecBuildLogicTest(t, "distsql_window")
}

func TestExecBuild_experimental_distsql_planning_5node(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runExecBuildLogicTest(t, "experimental_distsql_planning_5node")
}

func TestExecBuild_explain_analyze_plans(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runExecBuildLogicTest(t, "explain_analyze_plans")
}

func TestExecBuild_inverted_filter_geospatial_dist(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runExecBuildLogicTest(t, "inverted_filter_geospatial_dist")
}

func TestExecBuild_inverted_filter_json_array_dist(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runExecBuildLogicTest(t, "inverted_filter_json_array_dist")
}

func TestExecBuild_inverted_join_geospatial_dist(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runExecBuildLogicTest(t, "inverted_join_geospatial_dist")
}

func TestExecBuild_inverted_join_geospatial_dist_vec(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runExecBuildLogicTest(t, "inverted_join_geospatial_dist_vec")
}

func TestExecBuild_inverted_join_json_array_dist(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runExecBuildLogicTest(t, "inverted_join_json_array_dist")
}

func TestExecBuild_inverted_join_multi_column_dist(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runExecBuildLogicTest(t, "inverted_join_multi_column_dist")
}

func TestExecBuild_lookup_join(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runExecBuildLogicTest(t, "lookup_join")
}

func TestExecBuild_merge_join_dist_vec(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runExecBuildLogicTest(t, "merge_join_dist_vec")
}

func TestExecBuild_scan_parallel(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runExecBuildLogicTest(t, "scan_parallel")
}

func TestExecBuild_stats(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runExecBuildLogicTest(t, "stats")
}
