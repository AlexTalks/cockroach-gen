// Copyright 2023 The Cockroach Authors.
//
// Licensed as a CockroachDB Enterprise file under the Cockroach Community
// License (the "License"); you may not use this file except in compliance with
// the License. You may obtain a copy of the License at
//
//     https://github.com/cockroachdb/cockroach/blob/master/licenses/CCL.txt

// Code generated by pkg/ccl/backupccl/testgen, DO NOT EDIT.

package backupccl

import (
	"testing"

	"github.com/cockroachdb/cockroach/pkg/testutils/skip"
	"github.com/cockroachdb/cockroach/pkg/util/leaktest"
	"github.com/cockroachdb/cockroach/pkg/util/log"
)

func TestDataDriven_alter_schedule_backup_options(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	skip.UnderRace(t, "takes ~3mins to run")

	runTestDataDriven(t, "pkg/ccl/backupccl/testdata/backup-restore/alter-schedule/backup-options")
}

func TestDataDriven_alter_schedule_missing_schedule(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	skip.UnderRace(t, "takes ~3mins to run")

	runTestDataDriven(t, "pkg/ccl/backupccl/testdata/backup-restore/alter-schedule/missing-schedule")
}

func TestDataDriven_alter_schedule_recurrence(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	skip.UnderRace(t, "takes ~3mins to run")

	runTestDataDriven(t, "pkg/ccl/backupccl/testdata/backup-restore/alter-schedule/recurrence")
}

func TestDataDriven_alter_schedule_schedule_options(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	skip.UnderRace(t, "takes ~3mins to run")

	runTestDataDriven(t, "pkg/ccl/backupccl/testdata/backup-restore/alter-schedule/schedule-options")
}

func TestDataDriven_backup_dropped_descriptors(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	skip.UnderRace(t, "takes ~3mins to run")

	runTestDataDriven(t, "pkg/ccl/backupccl/testdata/backup-restore/backup-dropped-descriptors")
}

func TestDataDriven_backup_dropped_descriptors_declarative(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	skip.UnderRace(t, "takes ~3mins to run")

	runTestDataDriven(t, "pkg/ccl/backupccl/testdata/backup-restore/backup-dropped-descriptors-declarative")
}

func TestDataDriven_backup_permissions(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	skip.UnderRace(t, "takes ~3mins to run")

	runTestDataDriven(t, "pkg/ccl/backupccl/testdata/backup-restore/backup-permissions")
}

func TestDataDriven_backup_permissions_deprecated(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	skip.UnderRace(t, "takes ~3mins to run")

	runTestDataDriven(t, "pkg/ccl/backupccl/testdata/backup-restore/backup-permissions-deprecated")
}

func TestDataDriven_column_families(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	skip.UnderRace(t, "takes ~3mins to run")

	runTestDataDriven(t, "pkg/ccl/backupccl/testdata/backup-restore/column-families")
}

func TestDataDriven_descriptor_broadening(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	skip.UnderRace(t, "takes ~3mins to run")

	runTestDataDriven(t, "pkg/ccl/backupccl/testdata/backup-restore/descriptor-broadening")
}

func TestDataDriven_descriptor_conflicts(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	skip.UnderRace(t, "takes ~3mins to run")

	runTestDataDriven(t, "pkg/ccl/backupccl/testdata/backup-restore/descriptor-conflicts")
}

func TestDataDriven_drop_schedule_backup(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	skip.UnderRace(t, "takes ~3mins to run")

	runTestDataDriven(t, "pkg/ccl/backupccl/testdata/backup-restore/drop-schedule-backup")
}

func TestDataDriven_encrypted_backups(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	skip.UnderRace(t, "takes ~3mins to run")

	runTestDataDriven(t, "pkg/ccl/backupccl/testdata/backup-restore/encrypted-backups")
}

func TestDataDriven_external_connections_nodelocal(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	skip.UnderRace(t, "takes ~3mins to run")

	runTestDataDriven(t, "pkg/ccl/backupccl/testdata/backup-restore/external-connections-nodelocal")
}

func TestDataDriven_external_connections_privileges(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	skip.UnderRace(t, "takes ~3mins to run")

	runTestDataDriven(t, "pkg/ccl/backupccl/testdata/backup-restore/external-connections-privileges")
}

func TestDataDriven_external_connections_userfile(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	skip.UnderRace(t, "takes ~3mins to run")

	runTestDataDriven(t, "pkg/ccl/backupccl/testdata/backup-restore/external-connections-userfile")
}

func TestDataDriven_feature_flags(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	skip.UnderRace(t, "takes ~3mins to run")

	runTestDataDriven(t, "pkg/ccl/backupccl/testdata/backup-restore/feature-flags")
}

func TestDataDriven_file_table_read_write(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	skip.UnderRace(t, "takes ~3mins to run")

	runTestDataDriven(t, "pkg/ccl/backupccl/testdata/backup-restore/file_table_read_write")
}

func TestDataDriven_import_epoch(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	skip.UnderRace(t, "takes ~3mins to run")

	runTestDataDriven(t, "pkg/ccl/backupccl/testdata/backup-restore/import-epoch")
}

func TestDataDriven_import_start_time(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	skip.UnderRace(t, "takes ~3mins to run")

	runTestDataDriven(t, "pkg/ccl/backupccl/testdata/backup-restore/import-start-time")
}

func TestDataDriven_in_progress_import_rollback(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	skip.UnderRace(t, "takes ~3mins to run")

	runTestDataDriven(t, "pkg/ccl/backupccl/testdata/backup-restore/in-progress-import-rollback")
}

func TestDataDriven_in_progress_imports(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	skip.UnderRace(t, "takes ~3mins to run")

	runTestDataDriven(t, "pkg/ccl/backupccl/testdata/backup-restore/in-progress-imports")
}

func TestDataDriven_in_progress_restores(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	skip.UnderRace(t, "takes ~3mins to run")

	runTestDataDriven(t, "pkg/ccl/backupccl/testdata/backup-restore/in-progress-restores")
}

func TestDataDriven_lock_concurrent_backups(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	skip.UnderRace(t, "takes ~3mins to run")

	runTestDataDriven(t, "pkg/ccl/backupccl/testdata/backup-restore/lock-concurrent-backups")
}

func TestDataDriven_materialized_view(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	skip.UnderRace(t, "takes ~3mins to run")

	runTestDataDriven(t, "pkg/ccl/backupccl/testdata/backup-restore/materialized_view")
}

func TestDataDriven_max_row_size(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	skip.UnderRace(t, "takes ~3mins to run")

	runTestDataDriven(t, "pkg/ccl/backupccl/testdata/backup-restore/max-row-size")
}

func TestDataDriven_metadata(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	skip.UnderRace(t, "takes ~3mins to run")

	runTestDataDriven(t, "pkg/ccl/backupccl/testdata/backup-restore/metadata")
}

func TestDataDriven_mismatched_localities(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	skip.UnderRace(t, "takes ~3mins to run")

	runTestDataDriven(t, "pkg/ccl/backupccl/testdata/backup-restore/mismatched-localities")
}

func TestDataDriven_multiregion(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	skip.UnderRace(t, "takes ~3mins to run")
	skip.UnderDeadlockWithIssue(t, 117927)

	runTestDataDriven(t, "pkg/ccl/backupccl/testdata/backup-restore/multiregion")
}

func TestDataDriven_online_restore_auto_stats(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	skip.UnderRace(t, "takes ~3mins to run")

	runTestDataDriven(t, "pkg/ccl/backupccl/testdata/backup-restore/online-restore-auto-stats")
}

func TestDataDriven_online_restore_cluster(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	skip.UnderRace(t, "takes ~3mins to run")

	runTestDataDriven(t, "pkg/ccl/backupccl/testdata/backup-restore/online-restore-cluster")
}

func TestDataDriven_online_restore_empty_database(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	skip.UnderRace(t, "takes ~3mins to run")

	runTestDataDriven(t, "pkg/ccl/backupccl/testdata/backup-restore/online-restore-empty-database")
}

func TestDataDriven_online_restore_in_progress_imports(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	skip.UnderRace(t, "takes ~3mins to run")

	runTestDataDriven(t, "pkg/ccl/backupccl/testdata/backup-restore/online-restore-in-progress-imports")
}

func TestDataDriven_online_restore_prefix_backups(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	skip.UnderRace(t, "takes ~3mins to run")

	runTestDataDriven(t, "pkg/ccl/backupccl/testdata/backup-restore/online-restore-prefix-backups")
}

func TestDataDriven_plpgsql_procedures(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	skip.UnderRace(t, "takes ~3mins to run")

	runTestDataDriven(t, "pkg/ccl/backupccl/testdata/backup-restore/plpgsql_procedures")
}

func TestDataDriven_plpgsql_user_defined_functions(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	skip.UnderRace(t, "takes ~3mins to run")

	runTestDataDriven(t, "pkg/ccl/backupccl/testdata/backup-restore/plpgsql_user_defined_functions")
}

func TestDataDriven_procedures(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	skip.UnderRace(t, "takes ~3mins to run")

	runTestDataDriven(t, "pkg/ccl/backupccl/testdata/backup-restore/procedures")
}

func TestDataDriven_rangekeys(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	skip.UnderRace(t, "takes ~3mins to run")

	runTestDataDriven(t, "pkg/ccl/backupccl/testdata/backup-restore/rangekeys")
}

func TestDataDriven_rangekeys_revision_history(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	skip.UnderRace(t, "takes ~3mins to run")

	runTestDataDriven(t, "pkg/ccl/backupccl/testdata/backup-restore/rangekeys-revision-history")
}

func TestDataDriven_regression_tests(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	skip.UnderRace(t, "takes ~3mins to run")

	runTestDataDriven(t, "pkg/ccl/backupccl/testdata/backup-restore/regression-tests")
}

func TestDataDriven_restore_check_descriptors(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	skip.UnderRace(t, "takes ~3mins to run")

	runTestDataDriven(t, "pkg/ccl/backupccl/testdata/backup-restore/restore-check-descriptors")
}

func TestDataDriven_restore_grants(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	skip.UnderRace(t, "takes ~3mins to run")

	runTestDataDriven(t, "pkg/ccl/backupccl/testdata/backup-restore/restore-grants")
}

func TestDataDriven_restore_mixed_version(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	skip.UnderRace(t, "takes ~3mins to run")

	runTestDataDriven(t, "pkg/ccl/backupccl/testdata/backup-restore/restore-mixed-version")
}

func TestDataDriven_restore_on_fail_or_cancel_fast_drop(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	skip.UnderRace(t, "takes ~3mins to run")

	runTestDataDriven(t, "pkg/ccl/backupccl/testdata/backup-restore/restore-on-fail-or-cancel-fast-drop")
}

func TestDataDriven_restore_on_fail_or_cancel_retry(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	skip.UnderRace(t, "takes ~3mins to run")

	runTestDataDriven(t, "pkg/ccl/backupccl/testdata/backup-restore/restore-on-fail-or-cancel-retry")
}

func TestDataDriven_restore_on_fail_or_cancel_schema_objects(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	skip.UnderRace(t, "takes ~3mins to run")

	runTestDataDriven(t, "pkg/ccl/backupccl/testdata/backup-restore/restore-on-fail-or-cancel-schema-objects")
}

func TestDataDriven_restore_on_fail_or_cancel_ttl(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	skip.UnderRace(t, "takes ~3mins to run")

	runTestDataDriven(t, "pkg/ccl/backupccl/testdata/backup-restore/restore-on-fail-or-cancel-ttl")
}

func TestDataDriven_restore_permissions(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	skip.UnderRace(t, "takes ~3mins to run")

	runTestDataDriven(t, "pkg/ccl/backupccl/testdata/backup-restore/restore-permissions")
}

func TestDataDriven_restore_permissions_deprecated(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	skip.UnderRace(t, "takes ~3mins to run")

	runTestDataDriven(t, "pkg/ccl/backupccl/testdata/backup-restore/restore-permissions-deprecated")
}

func TestDataDriven_restore_regionless_alter_region(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	skip.UnderRace(t, "takes ~3mins to run")

	runTestDataDriven(t, "pkg/ccl/backupccl/testdata/backup-restore/restore-regionless-alter-region")
}

func TestDataDriven_restore_regionless_mixed_version(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	skip.UnderRace(t, "takes ~3mins to run")

	runTestDataDriven(t, "pkg/ccl/backupccl/testdata/backup-restore/restore-regionless-mixed-version")
}

func TestDataDriven_restore_regionless_on_regionless(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	skip.UnderRace(t, "takes ~3mins to run")

	runTestDataDriven(t, "pkg/ccl/backupccl/testdata/backup-restore/restore-regionless-on-regionless")
}

func TestDataDriven_restore_regionless_regional_by_row(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	skip.UnderRace(t, "takes ~3mins to run")

	runTestDataDriven(t, "pkg/ccl/backupccl/testdata/backup-restore/restore-regionless-regional-by-row")
}

func TestDataDriven_restore_schema_only(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	skip.UnderRace(t, "takes ~3mins to run")

	runTestDataDriven(t, "pkg/ccl/backupccl/testdata/backup-restore/restore-schema-only")
}

func TestDataDriven_restore_schema_only_multiregion(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	skip.UnderRace(t, "takes ~3mins to run")

	runTestDataDriven(t, "pkg/ccl/backupccl/testdata/backup-restore/restore-schema-only-multiregion")
}

func TestDataDriven_restore_tenants(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	skip.UnderRace(t, "takes ~3mins to run")

	runTestDataDriven(t, "pkg/ccl/backupccl/testdata/backup-restore/restore-tenants")
}

func TestDataDriven_restore_validation_only(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	skip.UnderRace(t, "takes ~3mins to run")

	runTestDataDriven(t, "pkg/ccl/backupccl/testdata/backup-restore/restore-validation-only")
}

func TestDataDriven_revision_history(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	skip.UnderRace(t, "takes ~3mins to run")

	runTestDataDriven(t, "pkg/ccl/backupccl/testdata/backup-restore/revision_history")
}

func TestDataDriven_row_level_ttl(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	skip.UnderRace(t, "takes ~3mins to run")

	runTestDataDriven(t, "pkg/ccl/backupccl/testdata/backup-restore/row_level_ttl")
}

func TestDataDriven_schedule_privileges(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	skip.UnderRace(t, "takes ~3mins to run")

	runTestDataDriven(t, "pkg/ccl/backupccl/testdata/backup-restore/schedule-privileges")
}

func TestDataDriven_show_backup_multiregion(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	skip.UnderRace(t, "takes ~3mins to run")

	runTestDataDriven(t, "pkg/ccl/backupccl/testdata/backup-restore/show-backup-multiregion")
}

func TestDataDriven_show_backup_union(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	skip.UnderRace(t, "takes ~3mins to run")

	runTestDataDriven(t, "pkg/ccl/backupccl/testdata/backup-restore/show-backup-union")
}

func TestDataDriven_show_schedules_old(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	skip.UnderRace(t, "takes ~3mins to run")

	runTestDataDriven(t, "pkg/ccl/backupccl/testdata/backup-restore/show-schedules-old")
}

func TestDataDriven_show_backup(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	skip.UnderRace(t, "takes ~3mins to run")

	runTestDataDriven(t, "pkg/ccl/backupccl/testdata/backup-restore/show_backup")
}

func TestDataDriven_system_privileges_table(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	skip.UnderRace(t, "takes ~3mins to run")

	runTestDataDriven(t, "pkg/ccl/backupccl/testdata/backup-restore/system-privileges-table")
}

func TestDataDriven_system_users(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	skip.UnderRace(t, "takes ~3mins to run")

	runTestDataDriven(t, "pkg/ccl/backupccl/testdata/backup-restore/system-users")
}

func TestDataDriven_temp_tables(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	skip.UnderRace(t, "takes ~3mins to run")

	runTestDataDriven(t, "pkg/ccl/backupccl/testdata/backup-restore/temp-tables")
}

func TestDataDriven_unique_without_index_constraint(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	skip.UnderRace(t, "takes ~3mins to run")

	runTestDataDriven(t, "pkg/ccl/backupccl/testdata/backup-restore/unique_without_index_constraint")
}

func TestDataDriven_user_defined_functions(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	skip.UnderRace(t, "takes ~3mins to run")

	runTestDataDriven(t, "pkg/ccl/backupccl/testdata/backup-restore/user-defined-functions")
}

func TestDataDriven_user_defined_functions_in_checks(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	skip.UnderRace(t, "takes ~3mins to run")

	runTestDataDriven(t, "pkg/ccl/backupccl/testdata/backup-restore/user-defined-functions-in-checks")
}

func TestDataDriven_user_defined_functions_in_defaults(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	skip.UnderRace(t, "takes ~3mins to run")

	runTestDataDriven(t, "pkg/ccl/backupccl/testdata/backup-restore/user-defined-functions-in-defaults")
}

func TestDataDriven_user_defined_types(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	skip.UnderRace(t, "takes ~3mins to run")

	runTestDataDriven(t, "pkg/ccl/backupccl/testdata/backup-restore/user-defined-types")
}

func TestDataDriven_views(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	skip.UnderRace(t, "takes ~3mins to run")

	runTestDataDriven(t, "pkg/ccl/backupccl/testdata/backup-restore/views")
}

func TestDataDriven_virtual_columns(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	skip.UnderRace(t, "takes ~3mins to run")

	runTestDataDriven(t, "pkg/ccl/backupccl/testdata/backup-restore/virtual-columns")
}
