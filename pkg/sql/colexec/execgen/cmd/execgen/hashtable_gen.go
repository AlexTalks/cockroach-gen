// Copyright 2020 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package main

import (
	"io"
	"io/ioutil"
	"strings"
	"text/template"

	"github.com/cockroachdb/cockroach/pkg/sql/sem/tree"
)

func genHashTable(wr io.Writer) error {
	t, err := ioutil.ReadFile("pkg/sql/colexec/hashtable_tmpl.go")
	if err != nil {
		return err
	}

	s := string(t)

	s = strings.Replace(s, "_TYPES_T", "coltypes.{{.LTyp}}", -1)
	s = strings.Replace(s, "_TYPE", "{{.LTyp}}", -1)
	s = strings.Replace(s, "_TemplateType", "{{.LTyp}}", -1)

	assignNeRe := makeFunctionRegex("_ASSIGN_NE", 3)
	s = assignNeRe.ReplaceAllString(s, `{{.Global.Assign "$1" "$2" "$3"}}`)

	checkCol := makeFunctionRegex("_CHECK_COL_WITH_NULLS", 7)
	s = checkCol.ReplaceAllString(s, `{{template "checkColWithNulls" buildDict "Global" . "UseSel" $7}}`)

	checkColBody := makeFunctionRegex("_CHECK_COL_BODY", 9)
	s = checkColBody.ReplaceAllString(
		s,
		`{{template "checkColBody" buildDict "Global" .Global "UseSel" .UseSel "ProbeHasNulls" $7 "BuildHasNulls" $8 "AllowNullEquality" $9}}`)

	s = replaceManipulationFuncs(".Global.LTyp", s)

	tmpl, err := template.New("hashtable").Funcs(template.FuncMap{"buildDict": buildDict}).Parse(s)
	if err != nil {
		return err
	}

	return tmpl.Execute(wr, sameTypeComparisonOpToOverloads[tree.NE])
}

func init() {
	registerGenerator(genHashTable, "hashtable.eg.go")
}
