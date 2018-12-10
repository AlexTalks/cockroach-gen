// Code generated by help.awk. DO NOT EDIT.
// GENERATED FILE DO NOT EDIT

package parser

var helpMessages = map[string]HelpMessageBody{
	//line sql.y: 1075
	`ALTER`: {
		//line sql.y: 1076
		Category: hGroup,
		//line sql.y: 1077
		Text: `ALTER TABLE, ALTER INDEX, ALTER VIEW, ALTER SEQUENCE, ALTER DATABASE, ALTER USER
`,
	},
	//line sql.y: 1091
	`ALTER TABLE`: {
		ShortDescription: `change the definition of a table`,
		//line sql.y: 1092
		Category: hDDL,
		//line sql.y: 1093
		Text: `
ALTER TABLE [IF EXISTS] <tablename> <command> [, ...]

Commands:
  ALTER TABLE ... ADD [COLUMN] [IF NOT EXISTS] <colname> <type> [<qualifiers...>]
  ALTER TABLE ... ADD <constraint>
  ALTER TABLE ... DROP [COLUMN] [IF EXISTS] <colname> [RESTRICT | CASCADE]
  ALTER TABLE ... DROP CONSTRAINT [IF EXISTS] <constraintname> [RESTRICT | CASCADE]
  ALTER TABLE ... ALTER [COLUMN] <colname> {SET DEFAULT <expr> | DROP DEFAULT}
  ALTER TABLE ... ALTER [COLUMN] <colname> DROP NOT NULL
  ALTER TABLE ... ALTER [COLUMN] <colname> DROP STORED
  ALTER TABLE ... ALTER [COLUMN] <colname> [SET DATA] TYPE <type> [COLLATE <collation>]
  ALTER TABLE ... RENAME TO <newname>
  ALTER TABLE ... RENAME [COLUMN] <colname> TO <newname>
  ALTER TABLE ... VALIDATE CONSTRAINT <constraintname>
  ALTER TABLE ... SPLIT AT <selectclause>
  ALTER TABLE ... SCATTER [ FROM ( <exprs...> ) TO ( <exprs...> ) ]
  ALTER TABLE ... INJECT STATISTICS ...  (experimental)
  ALTER TABLE ... PARTITION BY RANGE ( <name...> ) ( <rangespec> )
  ALTER TABLE ... PARTITION BY LIST ( <name...> ) ( <listspec> )
  ALTER TABLE ... PARTITION BY NOTHING
  ALTER TABLE ... CONFIGURE ZONE <zoneconfig>
  ALTER PARTITION ... OF TABLE ... CONFIGURE ZONE <zoneconfig>

Column qualifiers:
  [CONSTRAINT <constraintname>] {NULL | NOT NULL | UNIQUE | PRIMARY KEY | CHECK (<expr>) | DEFAULT <expr>}
  FAMILY <familyname>, CREATE [IF NOT EXISTS] FAMILY [<familyname>]
  REFERENCES <tablename> [( <colnames...> )]
  COLLATE <collationname>

Zone configurations:
  DISCARD
  USING <var> = <expr> [, ...]
  { TO | = } <expr>

`,
		//line sql.y: 1128
		SeeAlso: `WEBDOCS/alter-table.html
`,
	},
	//line sql.y: 1142
	`ALTER VIEW`: {
		ShortDescription: `change the definition of a view`,
		//line sql.y: 1143
		Category: hDDL,
		//line sql.y: 1144
		Text: `
ALTER VIEW [IF EXISTS] <name> RENAME TO <newname>
`,
		//line sql.y: 1146
		SeeAlso: `WEBDOCS/alter-view.html
`,
	},
	//line sql.y: 1153
	`ALTER SEQUENCE`: {
		ShortDescription: `change the definition of a sequence`,
		//line sql.y: 1154
		Category: hDDL,
		//line sql.y: 1155
		Text: `
ALTER SEQUENCE [IF EXISTS] <name>
  [INCREMENT <increment>]
  [MINVALUE <minvalue> | NO MINVALUE]
  [MAXVALUE <maxvalue> | NO MAXVALUE]
  [START <start>]
  [[NO] CYCLE]
ALTER SEQUENCE [IF EXISTS] <name> RENAME TO <newname>
`,
	},
	//line sql.y: 1188
	`ALTER USER`: {
		ShortDescription: `change user properties`,
		//line sql.y: 1189
		Category: hPriv,
		//line sql.y: 1190
		Text: `
ALTER USER [IF EXISTS] <name> WITH PASSWORD <password>
`,
		//line sql.y: 1192
		SeeAlso: `CREATE USER
`,
	},
	//line sql.y: 1197
	`ALTER DATABASE`: {
		ShortDescription: `change the definition of a database`,
		//line sql.y: 1198
		Category: hDDL,
		//line sql.y: 1199
		Text: `
ALTER DATABASE <name> RENAME TO <newname>
`,
		//line sql.y: 1201
		SeeAlso: `WEBDOCS/alter-database.html
`,
	},
	//line sql.y: 1209
	`ALTER RANGE`: {
		ShortDescription: `change the parameters of a range`,
		//line sql.y: 1210
		Category: hDDL,
		//line sql.y: 1211
		Text: `
ALTER RANGE <zonename> <command>

Commands:
  ALTER RANGE ... CONFIGURE ZONE <zoneconfig>

Zone configurations:
  DISCARD
  USING <var> = <expr> [, ...]
  { TO | = } <expr>

`,
		//line sql.y: 1222
		SeeAlso: `ALTER TABLE
`,
	},
	//line sql.y: 1227
	`ALTER INDEX`: {
		ShortDescription: `change the definition of an index`,
		//line sql.y: 1228
		Category: hDDL,
		//line sql.y: 1229
		Text: `
ALTER INDEX [IF EXISTS] <idxname> <command>

Commands:
  ALTER INDEX ... RENAME TO <newname>
  ALTER INDEX ... SPLIT AT <selectclause>
  ALTER INDEX ... SCATTER [ FROM ( <exprs...> ) TO ( <exprs...> ) ]

`,
		//line sql.y: 1237
		SeeAlso: `WEBDOCS/alter-index.html
`,
	},
	//line sql.y: 1656
	`BACKUP`: {
		ShortDescription: `back up data to external storage`,
		//line sql.y: 1657
		Category: hCCL,
		//line sql.y: 1658
		Text: `
BACKUP <targets...> TO <location...>
       [ AS OF SYSTEM TIME <expr> ]
       [ INCREMENTAL FROM <location...> ]
       [ WITH <option> [= <value>] [, ...] ]

Targets:
   TABLE <pattern> [, ...]
   DATABASE <databasename> [, ...]

Location:
   "[scheme]://[host]/[path to backup]?[parameters]"

Options:
   INTO_DB
   SKIP_MISSING_FOREIGN_KEYS

`,
		//line sql.y: 1675
		SeeAlso: `RESTORE, WEBDOCS/backup.html
`,
	},
	//line sql.y: 1683
	`RESTORE`: {
		ShortDescription: `restore data from external storage`,
		//line sql.y: 1684
		Category: hCCL,
		//line sql.y: 1685
		Text: `
RESTORE <targets...> FROM <location...>
        [ AS OF SYSTEM TIME <expr> ]
        [ WITH <option> [= <value>] [, ...] ]

Targets:
   TABLE <pattern> [, ...]
   DATABASE <databasename> [, ...]

Locations:
   "[scheme]://[host]/[path to backup]?[parameters]"

Options:
   INTO_DB
   SKIP_MISSING_FOREIGN_KEYS

`,
		//line sql.y: 1701
		SeeAlso: `BACKUP, WEBDOCS/restore.html
`,
	},
	//line sql.y: 1719
	`IMPORT`: {
		ShortDescription: `load data from file in a distributed manner`,
		//line sql.y: 1720
		Category: hCCL,
		//line sql.y: 1721
		Text: `
-- Import both schema and table data:
IMPORT [ TABLE <tablename> FROM ]
       <format> <datafile>
       [ WITH <option> [= <value>] [, ...] ]

-- Import using specific schema, use only table data from external file:
IMPORT TABLE <tablename>
       { ( <elements> ) | CREATE USING <schemafile> }
       <format>
       DATA ( <datafile> [, ...] )
       [ WITH <option> [= <value>] [, ...] ]

Formats:
   CSV
   MYSQLOUTFILE
   MYSQLDUMP (mysqldump's SQL output)
   PGCOPY
   PGDUMP

Options:
   distributed = '...'
   sstsize = '...'
   temp = '...'
   delimiter = '...'      [CSV, PGCOPY-specific]
   nullif = '...'         [CSV, PGCOPY-specific]
   comment = '...'        [CSV-specific]

`,
		//line sql.y: 1749
		SeeAlso: `CREATE TABLE
`,
	},
	//line sql.y: 1799
	`EXPORT`: {
		ShortDescription: `export data to file in a distributed manner`,
		//line sql.y: 1800
		Category: hCCL,
		//line sql.y: 1801
		Text: `
EXPORT INTO <format> (<datafile> [WITH <option> [= value] [,...]]) FROM <query>

Formats:
   CSV

Options:
   delimiter = '...'   [CSV-specific]

`,
		//line sql.y: 1810
		SeeAlso: `SELECT
`,
	},
	//line sql.y: 1905
	`CANCEL`: {
		//line sql.y: 1906
		Category: hGroup,
		//line sql.y: 1907
		Text: `CANCEL JOBS, CANCEL QUERIES, CANCEL SESSIONS
`,
	},
	//line sql.y: 1914
	`CANCEL JOBS`: {
		ShortDescription: `cancel background jobs`,
		//line sql.y: 1915
		Category: hMisc,
		//line sql.y: 1916
		Text: `
CANCEL JOBS <selectclause>
CANCEL JOB <jobid>
`,
		//line sql.y: 1919
		SeeAlso: `SHOW JOBS, PAUSE JOBS, RESUME JOBS
`,
	},
	//line sql.y: 1937
	`CANCEL QUERIES`: {
		ShortDescription: `cancel running queries`,
		//line sql.y: 1938
		Category: hMisc,
		//line sql.y: 1939
		Text: `
CANCEL QUERIES [IF EXISTS] <selectclause>
CANCEL QUERY [IF EXISTS] <expr>
`,
		//line sql.y: 1942
		SeeAlso: `SHOW QUERIES
`,
	},
	//line sql.y: 1973
	`CANCEL SESSIONS`: {
		ShortDescription: `cancel open sessions`,
		//line sql.y: 1974
		Category: hMisc,
		//line sql.y: 1975
		Text: `
CANCEL SESSIONS [IF EXISTS] <selectclause>
CANCEL SESSION [IF EXISTS] <sessionid>
`,
		//line sql.y: 1978
		SeeAlso: `SHOW SESSIONS
`,
	},
	//line sql.y: 2039
	`CREATE`: {
		//line sql.y: 2040
		Category: hGroup,
		//line sql.y: 2041
		Text: `
CREATE DATABASE, CREATE TABLE, CREATE INDEX, CREATE TABLE AS,
CREATE USER, CREATE VIEW, CREATE SEQUENCE, CREATE STATISTICS,
CREATE ROLE
`,
	},
	//line sql.y: 2122
	`CREATE STATISTICS`: {
		ShortDescription: `create a new table statistic (experimental)`,
		//line sql.y: 2123
		Category: hExperimental,
		//line sql.y: 2124
		Text: `
CREATE STATISTICS <statisticname>
  ON <colname> [, ...]
  FROM <tablename> [AS OF SYSTEM TIME <expr>]
`,
	},
	//line sql.y: 2187
	`DELETE`: {
		ShortDescription: `delete rows from a table`,
		//line sql.y: 2188
		Category: hDML,
		//line sql.y: 2189
		Text: `DELETE FROM <tablename> [WHERE <expr>]
              [ORDER BY <exprs...>]
              [LIMIT <expr>]
              [RETURNING <exprs...>]
`,
		//line sql.y: 2193
		SeeAlso: `WEBDOCS/delete.html
`,
	},
	//line sql.y: 2208
	`DISCARD`: {
		ShortDescription: `reset the session to its initial state`,
		//line sql.y: 2209
		Category: hCfg,
		//line sql.y: 2210
		Text: `DISCARD ALL
`,
	},
	//line sql.y: 2222
	`DROP`: {
		//line sql.y: 2223
		Category: hGroup,
		//line sql.y: 2224
		Text: `
DROP DATABASE, DROP INDEX, DROP TABLE, DROP VIEW, DROP SEQUENCE,
DROP USER, DROP ROLE
`,
	},
	//line sql.y: 2241
	`DROP VIEW`: {
		ShortDescription: `remove a view`,
		//line sql.y: 2242
		Category: hDDL,
		//line sql.y: 2243
		Text: `DROP VIEW [IF EXISTS] <tablename> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 2244
		SeeAlso: `WEBDOCS/drop-index.html
`,
	},
	//line sql.y: 2256
	`DROP SEQUENCE`: {
		ShortDescription: `remove a sequence`,
		//line sql.y: 2257
		Category: hDDL,
		//line sql.y: 2258
		Text: `DROP SEQUENCE [IF EXISTS] <sequenceName> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 2259
		SeeAlso: `DROP
`,
	},
	//line sql.y: 2271
	`DROP TABLE`: {
		ShortDescription: `remove a table`,
		//line sql.y: 2272
		Category: hDDL,
		//line sql.y: 2273
		Text: `DROP TABLE [IF EXISTS] <tablename> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 2274
		SeeAlso: `WEBDOCS/drop-table.html
`,
	},
	//line sql.y: 2286
	`DROP INDEX`: {
		ShortDescription: `remove an index`,
		//line sql.y: 2287
		Category: hDDL,
		//line sql.y: 2288
		Text: `DROP INDEX [IF EXISTS] <idxname> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 2289
		SeeAlso: `WEBDOCS/drop-index.html
`,
	},
	//line sql.y: 2309
	`DROP DATABASE`: {
		ShortDescription: `remove a database`,
		//line sql.y: 2310
		Category: hDDL,
		//line sql.y: 2311
		Text: `DROP DATABASE [IF EXISTS] <databasename> [CASCADE | RESTRICT]
`,
		//line sql.y: 2312
		SeeAlso: `WEBDOCS/drop-database.html
`,
	},
	//line sql.y: 2332
	`DROP USER`: {
		ShortDescription: `remove a user`,
		//line sql.y: 2333
		Category: hPriv,
		//line sql.y: 2334
		Text: `DROP USER [IF EXISTS] <user> [, ...]
`,
		//line sql.y: 2335
		SeeAlso: `CREATE USER, SHOW USERS
`,
	},
	//line sql.y: 2347
	`DROP ROLE`: {
		ShortDescription: `remove a role`,
		//line sql.y: 2348
		Category: hPriv,
		//line sql.y: 2349
		Text: `DROP ROLE [IF EXISTS] <role> [, ...]
`,
		//line sql.y: 2350
		SeeAlso: `CREATE ROLE, SHOW ROLES
`,
	},
	//line sql.y: 2382
	`EXPLAIN`: {
		ShortDescription: `show the logical plan of a query`,
		//line sql.y: 2383
		Category: hMisc,
		//line sql.y: 2384
		Text: `
EXPLAIN <statement>
EXPLAIN ([PLAN ,] <planoptions...> ) <statement>
EXPLAIN [ANALYZE] (DISTSQL) <statement>
EXPLAIN ANALYZE [(DISTSQL)] <statement>

Explainable statements:
    SELECT, CREATE, DROP, ALTER, INSERT, UPSERT, UPDATE, DELETE,
    SHOW, EXPLAIN

Plan options:
    TYPES, VERBOSE, OPT

`,
		//line sql.y: 2397
		SeeAlso: `WEBDOCS/explain.html
`,
	},
	//line sql.y: 2457
	`PREPARE`: {
		ShortDescription: `prepare a statement for later execution`,
		//line sql.y: 2458
		Category: hMisc,
		//line sql.y: 2459
		Text: `PREPARE <name> [ ( <types...> ) ] AS <query>
`,
		//line sql.y: 2460
		SeeAlso: `EXECUTE, DEALLOCATE, DISCARD
`,
	},
	//line sql.y: 2482
	`EXECUTE`: {
		ShortDescription: `execute a statement prepared previously`,
		//line sql.y: 2483
		Category: hMisc,
		//line sql.y: 2484
		Text: `EXECUTE <name> [ ( <exprs...> ) ]
`,
		//line sql.y: 2485
		SeeAlso: `PREPARE, DEALLOCATE, DISCARD
`,
	},
	//line sql.y: 2506
	`DEALLOCATE`: {
		ShortDescription: `remove a prepared statement`,
		//line sql.y: 2507
		Category: hMisc,
		//line sql.y: 2508
		Text: `DEALLOCATE [PREPARE] { <name> | ALL }
`,
		//line sql.y: 2509
		SeeAlso: `PREPARE, EXECUTE, DISCARD
`,
	},
	//line sql.y: 2529
	`GRANT`: {
		ShortDescription: `define access privileges and role memberships`,
		//line sql.y: 2530
		Category: hPriv,
		//line sql.y: 2531
		Text: `
Grant privileges:
  GRANT {ALL | <privileges...> } ON <targets...> TO <grantees...>
Grant role membership (CCL only):
  GRANT <roles...> TO <grantees...> [WITH ADMIN OPTION]

Privileges:
  CREATE, DROP, GRANT, SELECT, INSERT, DELETE, UPDATE

Targets:
  DATABASE <databasename> [, ...]
  [TABLE] [<databasename> .] { <tablename> | * } [, ...]

`,
		//line sql.y: 2544
		SeeAlso: `REVOKE, WEBDOCS/grant.html
`,
	},
	//line sql.y: 2560
	`REVOKE`: {
		ShortDescription: `remove access privileges and role memberships`,
		//line sql.y: 2561
		Category: hPriv,
		//line sql.y: 2562
		Text: `
Revoke privileges:
  REVOKE {ALL | <privileges...> } ON <targets...> FROM <grantees...>
Revoke role membership (CCL only):
  REVOKE [ADMIN OPTION FOR] <roles...> FROM <grantees...>

Privileges:
  CREATE, DROP, GRANT, SELECT, INSERT, DELETE, UPDATE

Targets:
  DATABASE <databasename> [, <databasename>]...
  [TABLE] [<databasename> .] { <tablename> | * } [, ...]

`,
		//line sql.y: 2575
		SeeAlso: `GRANT, WEBDOCS/revoke.html
`,
	},
	//line sql.y: 2630
	`RESET`: {
		ShortDescription: `reset a session variable to its default value`,
		//line sql.y: 2631
		Category: hCfg,
		//line sql.y: 2632
		Text: `RESET [SESSION] <var>
`,
		//line sql.y: 2633
		SeeAlso: `RESET CLUSTER SETTING, WEBDOCS/set-vars.html
`,
	},
	//line sql.y: 2645
	`RESET CLUSTER SETTING`: {
		ShortDescription: `reset a cluster setting to its default value`,
		//line sql.y: 2646
		Category: hCfg,
		//line sql.y: 2647
		Text: `RESET CLUSTER SETTING <var>
`,
		//line sql.y: 2648
		SeeAlso: `SET CLUSTER SETTING, RESET
`,
	},
	//line sql.y: 2657
	`USE`: {
		ShortDescription: `set the current database`,
		//line sql.y: 2658
		Category: hCfg,
		//line sql.y: 2659
		Text: `USE <dbname>

"USE <dbname>" is an alias for "SET [SESSION] database = <dbname>".
`,
		//line sql.y: 2662
		SeeAlso: `SET SESSION, WEBDOCS/set-vars.html
`,
	},
	//line sql.y: 2683
	`SCRUB`: {
		ShortDescription: `run checks against databases or tables`,
		//line sql.y: 2684
		Category: hExperimental,
		//line sql.y: 2685
		Text: `
EXPERIMENTAL SCRUB TABLE <table> ...
EXPERIMENTAL SCRUB DATABASE <database>

The various checks that ca be run with SCRUB includes:
  - Physical table data (encoding)
  - Secondary index integrity
  - Constraint integrity (NOT NULL, CHECK, FOREIGN KEY, UNIQUE)
`,
		//line sql.y: 2693
		SeeAlso: `SCRUB TABLE, SCRUB DATABASE
`,
	},
	//line sql.y: 2699
	`SCRUB DATABASE`: {
		ShortDescription: `run scrub checks on a database`,
		//line sql.y: 2700
		Category: hExperimental,
		//line sql.y: 2701
		Text: `
EXPERIMENTAL SCRUB DATABASE <database>
                            [AS OF SYSTEM TIME <expr>]

All scrub checks will be run on the database. This includes:
  - Physical table data (encoding)
  - Secondary index integrity
  - Constraint integrity (NOT NULL, CHECK, FOREIGN KEY, UNIQUE)
`,
		//line sql.y: 2709
		SeeAlso: `SCRUB TABLE, SCRUB
`,
	},
	//line sql.y: 2717
	`SCRUB TABLE`: {
		ShortDescription: `run scrub checks on a table`,
		//line sql.y: 2718
		Category: hExperimental,
		//line sql.y: 2719
		Text: `
SCRUB TABLE <tablename>
            [AS OF SYSTEM TIME <expr>]
            [WITH OPTIONS <option> [, ...]]

Options:
  EXPERIMENTAL SCRUB TABLE ... WITH OPTIONS INDEX ALL
  EXPERIMENTAL SCRUB TABLE ... WITH OPTIONS INDEX (<index>...)
  EXPERIMENTAL SCRUB TABLE ... WITH OPTIONS CONSTRAINT ALL
  EXPERIMENTAL SCRUB TABLE ... WITH OPTIONS CONSTRAINT (<constraint>...)
  EXPERIMENTAL SCRUB TABLE ... WITH OPTIONS PHYSICAL
`,
		//line sql.y: 2730
		SeeAlso: `SCRUB DATABASE, SRUB
`,
	},
	//line sql.y: 2790
	`SET CLUSTER SETTING`: {
		ShortDescription: `change a cluster setting`,
		//line sql.y: 2791
		Category: hCfg,
		//line sql.y: 2792
		Text: `SET CLUSTER SETTING <var> { TO | = } <value>
`,
		//line sql.y: 2793
		SeeAlso: `SHOW CLUSTER SETTING, RESET CLUSTER SETTING, SET SESSION,
WEBDOCS/cluster-settings.html
`,
	},
	//line sql.y: 2814
	`SET SESSION`: {
		ShortDescription: `change a session variable`,
		//line sql.y: 2815
		Category: hCfg,
		//line sql.y: 2816
		Text: `
SET [SESSION] <var> { TO | = } <values...>
SET [SESSION] TIME ZONE <tz>
SET [SESSION] CHARACTERISTICS AS TRANSACTION ISOLATION LEVEL { SNAPSHOT | SERIALIZABLE }
SET [SESSION] TRACING { TO | = } { on | off | cluster | local | kv | results } [,...]

`,
		//line sql.y: 2822
		SeeAlso: `SHOW SESSION, RESET, DISCARD, SHOW, SET CLUSTER SETTING, SET TRANSACTION,
WEBDOCS/set-vars.html
`,
	},
	//line sql.y: 2839
	`SET TRANSACTION`: {
		ShortDescription: `configure the transaction settings`,
		//line sql.y: 2840
		Category: hTxn,
		//line sql.y: 2841
		Text: `
SET [SESSION] TRANSACTION <txnparameters...>

Transaction parameters:
   ISOLATION LEVEL { SNAPSHOT | SERIALIZABLE }
   PRIORITY { LOW | NORMAL | HIGH }

`,
		//line sql.y: 2848
		SeeAlso: `SHOW TRANSACTION, SET SESSION,
WEBDOCS/set-transaction.html
`,
	},
	//line sql.y: 3031
	`SHOW`: {
		//line sql.y: 3032
		Category: hGroup,
		//line sql.y: 3033
		Text: `
SHOW BACKUP, SHOW CLUSTER SETTING, SHOW COLUMNS, SHOW CONSTRAINTS,
SHOW CREATE, SHOW DATABASES, SHOW HISTOGRAM, SHOW INDEXES, SHOW JOBS,
SHOW QUERIES, SHOW ROLES, SHOW SESSION, SHOW SESSIONS, SHOW STATISTICS,
SHOW SYNTAX, SHOW TABLES, SHOW TRACE SHOW TRANSACTION, SHOW USERS
`,
	},
	//line sql.y: 3065
	`SHOW SESSION`: {
		ShortDescription: `display session variables`,
		//line sql.y: 3066
		Category: hCfg,
		//line sql.y: 3067
		Text: `SHOW [SESSION] { <var> | ALL }
`,
		//line sql.y: 3068
		SeeAlso: `WEBDOCS/show-vars.html
`,
	},
	//line sql.y: 3089
	`SHOW STATISTICS`: {
		ShortDescription: `display table statistics (experimental)`,
		//line sql.y: 3090
		Category: hExperimental,
		//line sql.y: 3091
		Text: `SHOW STATISTICS [USING JSON] FOR TABLE <table_name>

Returns the available statistics for a table.
The statistics can include a histogram ID, which can
be used with SHOW HISTOGRAM.
If USING JSON is specified, the statistics and histograms
are encoded in JSON format.
`,
		//line sql.y: 3098
		SeeAlso: `SHOW HISTOGRAM
`,
	},
	//line sql.y: 3121
	`SHOW HISTOGRAM`: {
		ShortDescription: `display histogram (experimental)`,
		//line sql.y: 3122
		Category: hExperimental,
		//line sql.y: 3123
		Text: `SHOW HISTOGRAM <histogram_id>

Returns the data in the histogram with the
given ID (as returned by SHOW STATISTICS).
`,
		//line sql.y: 3127
		SeeAlso: `SHOW STATISTICS
`,
	},
	//line sql.y: 3141
	`SHOW BACKUP`: {
		ShortDescription: `list backup contents`,
		//line sql.y: 3142
		Category: hCCL,
		//line sql.y: 3143
		Text: `SHOW BACKUP [FILES|RANGES] <location>
`,
		//line sql.y: 3144
		SeeAlso: `WEBDOCS/show-backup.html
`,
	},
	//line sql.y: 3171
	`SHOW CLUSTER SETTING`: {
		ShortDescription: `display cluster settings`,
		//line sql.y: 3172
		Category: hCfg,
		//line sql.y: 3173
		Text: `
SHOW CLUSTER SETTING <var>
SHOW ALL CLUSTER SETTINGS
`,
		//line sql.y: 3176
		SeeAlso: `WEBDOCS/cluster-settings.html
`,
	},
	//line sql.y: 3193
	`SHOW COLUMNS`: {
		ShortDescription: `list columns in relation`,
		//line sql.y: 3194
		Category: hDDL,
		//line sql.y: 3195
		Text: `SHOW COLUMNS FROM <tablename>
`,
		//line sql.y: 3196
		SeeAlso: `WEBDOCS/show-columns.html
`,
	},
	//line sql.y: 3209
	`SHOW DATABASES`: {
		ShortDescription: `list databases`,
		//line sql.y: 3210
		Category: hDDL,
		//line sql.y: 3211
		Text: `SHOW DATABASES
`,
		//line sql.y: 3212
		SeeAlso: `WEBDOCS/show-databases.html
`,
	},
	//line sql.y: 3220
	`SHOW GRANTS`: {
		ShortDescription: `list grants`,
		//line sql.y: 3221
		Category: hPriv,
		//line sql.y: 3222
		Text: `
Show privilege grants:
  SHOW GRANTS [ON <targets...>] [FOR <users...>]
Show role grants:
  SHOW GRANTS ON ROLE [<roles...>] [FOR <grantees...>]

`,
		//line sql.y: 3228
		SeeAlso: `WEBDOCS/show-grants.html
`,
	},
	//line sql.y: 3241
	`SHOW INDEXES`: {
		ShortDescription: `list indexes`,
		//line sql.y: 3242
		Category: hDDL,
		//line sql.y: 3243
		Text: `SHOW INDEXES FROM <tablename>
`,
		//line sql.y: 3244
		SeeAlso: `WEBDOCS/show-index.html
`,
	},
	//line sql.y: 3277
	`SHOW CONSTRAINTS`: {
		ShortDescription: `list constraints`,
		//line sql.y: 3278
		Category: hDDL,
		//line sql.y: 3279
		Text: `SHOW CONSTRAINTS FROM <tablename>
`,
		//line sql.y: 3280
		SeeAlso: `WEBDOCS/show-constraints.html
`,
	},
	//line sql.y: 3303
	`SHOW QUERIES`: {
		ShortDescription: `list running queries`,
		//line sql.y: 3304
		Category: hMisc,
		//line sql.y: 3305
		Text: `SHOW [CLUSTER | LOCAL] QUERIES
`,
		//line sql.y: 3306
		SeeAlso: `CANCEL QUERIES
`,
	},
	//line sql.y: 3322
	`SHOW JOBS`: {
		ShortDescription: `list background jobs`,
		//line sql.y: 3323
		Category: hMisc,
		//line sql.y: 3324
		Text: `SHOW JOBS
`,
		//line sql.y: 3325
		SeeAlso: `CANCEL JOBS, PAUSE JOBS, RESUME JOBS
`,
	},
	//line sql.y: 3333
	`SHOW TRACE`: {
		ShortDescription: `display an execution trace`,
		//line sql.y: 3334
		Category: hMisc,
		//line sql.y: 3335
		Text: `
SHOW [COMPACT] [KV] TRACE FOR SESSION
`,
		//line sql.y: 3337
		SeeAlso: `EXPLAIN
`,
	},
	//line sql.y: 3360
	`SHOW SESSIONS`: {
		ShortDescription: `list open client sessions`,
		//line sql.y: 3361
		Category: hMisc,
		//line sql.y: 3362
		Text: `SHOW [CLUSTER | LOCAL] SESSIONS
`,
		//line sql.y: 3363
		SeeAlso: `CANCEL SESSIONS
`,
	},
	//line sql.y: 3379
	`SHOW TABLES`: {
		ShortDescription: `list tables`,
		//line sql.y: 3380
		Category: hDDL,
		//line sql.y: 3381
		Text: `SHOW TABLES [FROM <databasename> [ . <schemaname> ] ]
`,
		//line sql.y: 3382
		SeeAlso: `WEBDOCS/show-tables.html
`,
	},
	//line sql.y: 3414
	`SHOW SCHEMAS`: {
		ShortDescription: `list schemas`,
		//line sql.y: 3415
		Category: hDDL,
		//line sql.y: 3416
		Text: `SHOW SCHEMAS [FROM <databasename> ]
`,
	},
	//line sql.y: 3428
	`SHOW SYNTAX`: {
		ShortDescription: `analyze SQL syntax`,
		//line sql.y: 3429
		Category: hMisc,
		//line sql.y: 3430
		Text: `SHOW SYNTAX <string>
`,
	},
	//line sql.y: 3439
	`SHOW TRANSACTION`: {
		ShortDescription: `display current transaction properties`,
		//line sql.y: 3440
		Category: hCfg,
		//line sql.y: 3441
		Text: `SHOW TRANSACTION {ISOLATION LEVEL | PRIORITY | STATUS}
`,
		//line sql.y: 3442
		SeeAlso: `WEBDOCS/show-transaction.html
`,
	},
	//line sql.y: 3461
	`SHOW CREATE`: {
		ShortDescription: `display the CREATE statement for a table, sequence or view`,
		//line sql.y: 3462
		Category: hDDL,
		//line sql.y: 3463
		Text: `SHOW CREATE [ TABLE | SEQUENCE | VIEW ] <tablename>
`,
		//line sql.y: 3464
		SeeAlso: `WEBDOCS/show-create-table.html
`,
	},
	//line sql.y: 3492
	`SHOW USERS`: {
		ShortDescription: `list defined users`,
		//line sql.y: 3493
		Category: hPriv,
		//line sql.y: 3494
		Text: `SHOW USERS
`,
		//line sql.y: 3495
		SeeAlso: `CREATE USER, DROP USER, WEBDOCS/show-users.html
`,
	},
	//line sql.y: 3503
	`SHOW ROLES`: {
		ShortDescription: `list defined roles`,
		//line sql.y: 3504
		Category: hPriv,
		//line sql.y: 3505
		Text: `SHOW ROLES
`,
		//line sql.y: 3506
		SeeAlso: `CREATE ROLE, DROP ROLE
`,
	},
	//line sql.y: 3561
	`SHOW RANGES`: {
		ShortDescription: `list ranges`,
		//line sql.y: 3562
		Category: hMisc,
		//line sql.y: 3563
		Text: `
SHOW EXPERIMENTAL_RANGES FROM TABLE <tablename>
SHOW EXPERIMENTAL_RANGES FROM INDEX [ <tablename> @ ] <indexname>
`,
	},
	//line sql.y: 3809
	`PAUSE JOBS`: {
		ShortDescription: `pause background jobs`,
		//line sql.y: 3810
		Category: hMisc,
		//line sql.y: 3811
		Text: `
PAUSE JOBS <selectclause>
PAUSE JOB <jobid>
`,
		//line sql.y: 3814
		SeeAlso: `SHOW JOBS, CANCEL JOBS, RESUME JOBS
`,
	},
	//line sql.y: 3831
	`CREATE TABLE`: {
		ShortDescription: `create a new table`,
		//line sql.y: 3832
		Category: hDDL,
		//line sql.y: 3833
		Text: `
CREATE TABLE [IF NOT EXISTS] <tablename> ( <elements...> ) [<interleave>]
CREATE TABLE [IF NOT EXISTS] <tablename> [( <colnames...> )] AS <source>

Table elements:
   <name> <type> [<qualifiers...>]
   [UNIQUE | INVERTED] INDEX [<name>] ( <colname> [ASC | DESC] [, ...] )
                           [STORING ( <colnames...> )] [<interleave>]
   FAMILY [<name>] ( <colnames...> )
   [CONSTRAINT <name>] <constraint>

Table constraints:
   PRIMARY KEY ( <colnames...> )
   FOREIGN KEY ( <colnames...> ) REFERENCES <tablename> [( <colnames...> )] [ON DELETE {NO ACTION | RESTRICT}] [ON UPDATE {NO ACTION | RESTRICT}]
   UNIQUE ( <colnames... ) [STORING ( <colnames...> )] [<interleave>]
   CHECK ( <expr> )

Column qualifiers:
  [CONSTRAINT <constraintname>] {NULL | NOT NULL | UNIQUE | PRIMARY KEY | CHECK (<expr>) | DEFAULT <expr>}
  FAMILY <familyname>, CREATE [IF NOT EXISTS] FAMILY [<familyname>]
  REFERENCES <tablename> [( <colnames...> )] [ON DELETE {NO ACTION | RESTRICT}] [ON UPDATE {NO ACTION | RESTRICT}]
  COLLATE <collationname>
  AS ( <expr> ) STORED

Interleave clause:
   INTERLEAVE IN PARENT <tablename> ( <colnames...> ) [CASCADE | RESTRICT]

`,
		//line sql.y: 3860
		SeeAlso: `SHOW TABLES, CREATE VIEW, SHOW CREATE,
WEBDOCS/create-table.html
WEBDOCS/create-table-as.html
`,
	},
	//line sql.y: 4465
	`CREATE SEQUENCE`: {
		ShortDescription: `create a new sequence`,
		//line sql.y: 4466
		Category: hDDL,
		//line sql.y: 4467
		Text: `
CREATE SEQUENCE <seqname>
  [INCREMENT <increment>]
  [MINVALUE <minvalue> | NO MINVALUE]
  [MAXVALUE <maxvalue> | NO MAXVALUE]
  [START [WITH] <start>]
  [CACHE <cache>]
  [NO CYCLE]
  [VIRTUAL]

`,
		//line sql.y: 4477
		SeeAlso: `CREATE TABLE
`,
	},
	//line sql.y: 4532
	`TRUNCATE`: {
		ShortDescription: `empty one or more tables`,
		//line sql.y: 4533
		Category: hDML,
		//line sql.y: 4534
		Text: `TRUNCATE [TABLE] <tablename> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 4535
		SeeAlso: `WEBDOCS/truncate.html
`,
	},
	//line sql.y: 4543
	`CREATE USER`: {
		ShortDescription: `define a new user`,
		//line sql.y: 4544
		Category: hPriv,
		//line sql.y: 4545
		Text: `CREATE USER [IF NOT EXISTS] <name> [ [WITH] PASSWORD <passwd> ]
`,
		//line sql.y: 4546
		SeeAlso: `DROP USER, SHOW USERS, WEBDOCS/create-user.html
`,
	},
	//line sql.y: 4568
	`CREATE ROLE`: {
		ShortDescription: `define a new role`,
		//line sql.y: 4569
		Category: hPriv,
		//line sql.y: 4570
		Text: `CREATE ROLE [IF NOT EXISTS] <name>
`,
		//line sql.y: 4571
		SeeAlso: `DROP ROLE, SHOW ROLES
`,
	},
	//line sql.y: 4589
	`CREATE VIEW`: {
		ShortDescription: `create a new view`,
		//line sql.y: 4590
		Category: hDDL,
		//line sql.y: 4591
		Text: `CREATE VIEW <viewname> [( <colnames...> )] AS <source>
`,
		//line sql.y: 4592
		SeeAlso: `CREATE TABLE, SHOW CREATE, WEBDOCS/create-view.html
`,
	},
	//line sql.y: 4630
	`CREATE INDEX`: {
		ShortDescription: `create a new index`,
		//line sql.y: 4631
		Category: hDDL,
		//line sql.y: 4632
		Text: `
CREATE [UNIQUE | INVERTED] INDEX [IF NOT EXISTS] [<idxname>]
       ON <tablename> ( <colname> [ASC | DESC] [, ...] )
       [STORING ( <colnames...> )] [<interleave>]

Interleave clause:
   INTERLEAVE IN PARENT <tablename> ( <colnames...> ) [CASCADE | RESTRICT]

`,
		//line sql.y: 4640
		SeeAlso: `CREATE TABLE, SHOW INDEXES, SHOW CREATE,
WEBDOCS/create-index.html
`,
	},
	//line sql.y: 4954
	`RELEASE`: {
		ShortDescription: `complete a retryable block`,
		//line sql.y: 4955
		Category: hTxn,
		//line sql.y: 4956
		Text: `RELEASE [SAVEPOINT] cockroach_restart
`,
		//line sql.y: 4957
		SeeAlso: `SAVEPOINT, WEBDOCS/savepoint.html
`,
	},
	//line sql.y: 4965
	`RESUME JOBS`: {
		ShortDescription: `resume background jobs`,
		//line sql.y: 4966
		Category: hMisc,
		//line sql.y: 4967
		Text: `
RESUME JOBS <selectclause>
RESUME JOB <jobid>
`,
		//line sql.y: 4970
		SeeAlso: `SHOW JOBS, CANCEL JOBS, PAUSE JOBS
`,
	},
	//line sql.y: 4987
	`SAVEPOINT`: {
		ShortDescription: `start a retryable block`,
		//line sql.y: 4988
		Category: hTxn,
		//line sql.y: 4989
		Text: `SAVEPOINT cockroach_restart
`,
		//line sql.y: 4990
		SeeAlso: `RELEASE, WEBDOCS/savepoint.html
`,
	},
	//line sql.y: 5005
	`BEGIN`: {
		ShortDescription: `start a transaction`,
		//line sql.y: 5006
		Category: hTxn,
		//line sql.y: 5007
		Text: `
BEGIN [TRANSACTION] [ <txnparameter> [[,] ...] ]
START TRANSACTION [ <txnparameter> [[,] ...] ]

Transaction parameters:
   ISOLATION LEVEL { SNAPSHOT | SERIALIZABLE }
   PRIORITY { LOW | NORMAL | HIGH }

`,
		//line sql.y: 5015
		SeeAlso: `COMMIT, ROLLBACK, WEBDOCS/begin-transaction.html
`,
	},
	//line sql.y: 5028
	`COMMIT`: {
		ShortDescription: `commit the current transaction`,
		//line sql.y: 5029
		Category: hTxn,
		//line sql.y: 5030
		Text: `
COMMIT [TRANSACTION]
END [TRANSACTION]
`,
		//line sql.y: 5033
		SeeAlso: `BEGIN, ROLLBACK, WEBDOCS/commit-transaction.html
`,
	},
	//line sql.y: 5057
	`ROLLBACK`: {
		ShortDescription: `abort the current transaction`,
		//line sql.y: 5058
		Category: hTxn,
		//line sql.y: 5059
		Text: `ROLLBACK [TRANSACTION] [TO [SAVEPOINT] cockroach_restart]
`,
		//line sql.y: 5060
		SeeAlso: `BEGIN, COMMIT, SAVEPOINT, WEBDOCS/rollback-transaction.html
`,
	},
	//line sql.y: 5174
	`CREATE DATABASE`: {
		ShortDescription: `create a new database`,
		//line sql.y: 5175
		Category: hDDL,
		//line sql.y: 5176
		Text: `CREATE DATABASE [IF NOT EXISTS] <name>
`,
		//line sql.y: 5177
		SeeAlso: `WEBDOCS/create-database.html
`,
	},
	//line sql.y: 5246
	`INSERT`: {
		ShortDescription: `create new rows in a table`,
		//line sql.y: 5247
		Category: hDML,
		//line sql.y: 5248
		Text: `
INSERT INTO <tablename> [[AS] <name>] [( <colnames...> )]
       <selectclause>
       [ON CONFLICT [( <colnames...> )] {DO UPDATE SET ... [WHERE <expr>] | DO NOTHING}]
       [RETURNING <exprs...>]
`,
		//line sql.y: 5253
		SeeAlso: `UPSERT, UPDATE, DELETE, WEBDOCS/insert.html
`,
	},
	//line sql.y: 5272
	`UPSERT`: {
		ShortDescription: `create or replace rows in a table`,
		//line sql.y: 5273
		Category: hDML,
		//line sql.y: 5274
		Text: `
UPSERT INTO <tablename> [AS <name>] [( <colnames...> )]
       <selectclause>
       [RETURNING <exprs...>]
`,
		//line sql.y: 5278
		SeeAlso: `INSERT, UPDATE, DELETE, WEBDOCS/upsert.html
`,
	},
	//line sql.y: 5393
	`UPDATE`: {
		ShortDescription: `update rows of a table`,
		//line sql.y: 5394
		Category: hDML,
		//line sql.y: 5395
		Text: `
UPDATE <tablename> [[AS] <name>]
       SET ...
       [WHERE <expr>]
       [ORDER BY <exprs...>]
       [LIMIT <expr>]
       [RETURNING <exprs...>]
`,
		//line sql.y: 5402
		SeeAlso: `INSERT, UPSERT, DELETE, WEBDOCS/update.html
`,
	},
	//line sql.y: 5576
	`<SELECTCLAUSE>`: {
		ShortDescription: `access tabular data`,
		//line sql.y: 5577
		Category: hDML,
		//line sql.y: 5578
		Text: `
Select clause:
  TABLE <tablename>
  VALUES ( <exprs...> ) [ , ... ]
  SELECT ... [ { INTERSECT | UNION | EXCEPT } [ ALL | DISTINCT ] <selectclause> ]
`,
	},
	//line sql.y: 5589
	`SELECT`: {
		ShortDescription: `retrieve rows from a data source and compute a result`,
		//line sql.y: 5590
		Category: hDML,
		//line sql.y: 5591
		Text: `
SELECT [DISTINCT [ ON ( <expr> [ , ... ] ) ] ]
       { <expr> [[AS] <name>] | [ [<dbname>.] <tablename>. ] * } [, ...]
       [ FROM <source> ]
       [ WHERE <expr> ]
       [ GROUP BY <expr> [ , ... ] ]
       [ HAVING <expr> ]
       [ WINDOW <name> AS ( <definition> ) ]
       [ { UNION | INTERSECT | EXCEPT } [ ALL | DISTINCT ] <selectclause> ]
       [ ORDER BY <expr> [ ASC | DESC ] [, ...] ]
       [ LIMIT { <expr> | ALL } ]
       [ OFFSET <expr> [ ROW | ROWS ] ]
`,
		//line sql.y: 5603
		SeeAlso: `WEBDOCS/select-clause.html
`,
	},
	//line sql.y: 5678
	`TABLE`: {
		ShortDescription: `select an entire table`,
		//line sql.y: 5679
		Category: hDML,
		//line sql.y: 5680
		Text: `TABLE <tablename>
`,
		//line sql.y: 5681
		SeeAlso: `SELECT, VALUES, WEBDOCS/table-expressions.html
`,
	},
	//line sql.y: 5969
	`VALUES`: {
		ShortDescription: `select a given set of values`,
		//line sql.y: 5970
		Category: hDML,
		//line sql.y: 5971
		Text: `VALUES ( <exprs...> ) [, ...]
`,
		//line sql.y: 5972
		SeeAlso: `SELECT, TABLE, WEBDOCS/table-expressions.html
`,
	},
	//line sql.y: 6062
	`<SOURCE>`: {
		ShortDescription: `define a data source for SELECT`,
		//line sql.y: 6063
		Category: hDML,
		//line sql.y: 6064
		Text: `
Data sources:
  <tablename> [ @ { <idxname> | <indexhint> } ]
  <tablefunc> ( <exprs...> )
  ( { <selectclause> | <source> } )
  <source> [AS] <alias> [( <colnames...> )]
  <source> { [INNER] | { LEFT | RIGHT | FULL } [OUTER] } JOIN <source> ON <expr>
  <source> { [INNER] | { LEFT | RIGHT | FULL } [OUTER] } JOIN <source> USING ( <colnames...> )
  <source> NATURAL { [INNER] | { LEFT | RIGHT | FULL } [OUTER] } JOIN <source>
  <source> CROSS JOIN <source>
  <source> WITH ORDINALITY
  '[' EXPLAIN ... ']'
  '[' SHOW ... ']'

Index flags:
  '{' FORCE_INDEX = <idxname> [, ...] '}'
  '{' NO_INDEX_JOIN [, ...] '}'

`,
		//line sql.y: 6082
		SeeAlso: `WEBDOCS/table-expressions.html
`,
	},
}
