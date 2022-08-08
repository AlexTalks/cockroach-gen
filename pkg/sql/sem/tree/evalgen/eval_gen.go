// Copyright 2022 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

// Command evalgen is used to generate interfaces and "visitor" methods for
// expr and op evaluation.
//
// Generated files can be regenerated with either of the follow commands:
//
//   ./dev generate go
//   go generate ./pkg/sql/sem/tree
//
package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"

	"github.com/cockroachdb/errors"
	"golang.org/x/tools/go/ast/inspector"
)

const (
	exprEvalFileName = "eval_expr_generated.go"
	opEvalFileName   = "eval_op_generated.go"
)

var generated = stringSet{
	exprEvalFileName: {},
	opEvalFileName:   {},
}

func main() {
	out := flag.String("out", ".", "path to write files")
	flag.Parse()
	fileSet := token.NewFileSet()
	var files []*ast.File
	byName := map[string]*ast.File{}
	for _, arg := range flag.Args() {
		expanded, err := filepath.Glob(arg)
		if err != nil {
			panic(fmt.Sprintf("failed to expand %s: %v", arg, err))
		}
		for _, fp := range expanded {
			name := filepath.Base(fp)
			if _, exists := byName[name]; generated.contains(name) || exists {
				continue
			}
			f, err := parser.ParseFile(fileSet, fp, nil, 0)
			if err != nil {
				panic(fmt.Sprintf("failed to parse %s: %v", arg, err))
			}
			files = append(files, f)
			byName[name] = f
		}
	}
	if err := generateExprEval(
		filepath.Join(*out, exprEvalFileName), files,
	); err != nil {
		panic(err)
	}
	if err := generateOpsFile(
		filepath.Join(*out, opEvalFileName), files, byName,
	); err != nil {
		panic(err)
	}
}

func writeFile(name string, fn func(f *os.File) error) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer func() {
		closeErr := f.Close()
		if err == nil {
			err = closeErr
		}
		if err != nil {
			_ = os.Remove(name)
		}
	}()
	return fn(f)
}

const header = `// Copyright 2022 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

// Code generated by eval_gen.go. DO NOT EDIT.
// Regenerate this file with either of the following commands:
//
//   ./dev generate go
//   go generate ./pkg/sql/sem/tree
//
// If you use the dev command and you have added a new tree expression, like
// tree.XYZ in a new file, you may get the confusing error: undefined: XYZ.
// Run './dev generate bazel' to fix this.
package tree
`

func findTypesWithMethod(ins *inspector.Inspector, filter func(decl *ast.FuncDecl) bool) stringSet {
	ss := stringSet{}
	ins.Preorder([]ast.Node{(*ast.FuncDecl)(nil)}, func(node ast.Node) {
		fd := node.(*ast.FuncDecl)
		if fd.Recv == nil || !filter(fd) {
			return
		}
		ss.add(extractMethodReceiverName(fd))
	})
	return ss
}

func extractMethodReceiver(fd *ast.FuncDecl) (name string, isPtr bool) {
	switch n := fd.Recv.List[0].Type.(type) {
	case *ast.StarExpr:
		return n.X.(*ast.Ident).Name, true
	case *ast.Ident:
		return n.Name, false
	default:
		panic(errors.Errorf("unexpected receiver type node %T", n))
	}
}

func extractMethodReceiverName(fd *ast.FuncDecl) (name string) {
	name, _ = extractMethodReceiver(fd)
	return name
}

func findStructTypesWithEmbeddedType(
	ins *inspector.Inspector, filter func(name string) bool, embeddedStruct string,
) stringSet {
	hasEmbedded := stringSet{}
	ins.Preorder([]ast.Node{
		(*ast.TypeSpec)(nil),
	}, func(node ast.Node) {
		ts := node.(*ast.TypeSpec)
		if !filter(ts.Name.Name) {
			return
		}
		st, ok := ts.Type.(*ast.StructType)
		if !ok {
			return
		}
		for _, f := range st.Fields.List {
			if len(f.Names) > 0 {
				continue
			}
			id, ok := f.Type.(*ast.Ident)
			if !ok {
				continue
			}
			if id.Name == embeddedStruct {
				hasEmbedded.add(ts.Name.Name)
				return
			}
		}
	})
	return hasEmbedded
}
