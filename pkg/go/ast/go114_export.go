// export by github.com/goplus/igop/cmd/qexp

//+build go1.14,!go1.15

package ast

import (
	q "go/ast"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "ast",
		Path: "go/ast",
		Deps: map[string]string{
			"bytes":      "bytes",
			"fmt":        "fmt",
			"go/scanner": "scanner",
			"go/token":   "token",
			"io":         "io",
			"os":         "os",
			"reflect":    "reflect",
			"sort":       "sort",
			"strconv":    "strconv",
			"strings":    "strings",
		},
		Interfaces: map[string]reflect.Type{
			"Decl":    reflect.TypeOf((*q.Decl)(nil)).Elem(),
			"Expr":    reflect.TypeOf((*q.Expr)(nil)).Elem(),
			"Node":    reflect.TypeOf((*q.Node)(nil)).Elem(),
			"Spec":    reflect.TypeOf((*q.Spec)(nil)).Elem(),
			"Stmt":    reflect.TypeOf((*q.Stmt)(nil)).Elem(),
			"Visitor": reflect.TypeOf((*q.Visitor)(nil)).Elem(),
		},
		NamedTypes: map[string]reflect.Type{
			"ArrayType":      reflect.TypeOf((*q.ArrayType)(nil)).Elem(),
			"AssignStmt":     reflect.TypeOf((*q.AssignStmt)(nil)).Elem(),
			"BadDecl":        reflect.TypeOf((*q.BadDecl)(nil)).Elem(),
			"BadExpr":        reflect.TypeOf((*q.BadExpr)(nil)).Elem(),
			"BadStmt":        reflect.TypeOf((*q.BadStmt)(nil)).Elem(),
			"BasicLit":       reflect.TypeOf((*q.BasicLit)(nil)).Elem(),
			"BinaryExpr":     reflect.TypeOf((*q.BinaryExpr)(nil)).Elem(),
			"BlockStmt":      reflect.TypeOf((*q.BlockStmt)(nil)).Elem(),
			"BranchStmt":     reflect.TypeOf((*q.BranchStmt)(nil)).Elem(),
			"CallExpr":       reflect.TypeOf((*q.CallExpr)(nil)).Elem(),
			"CaseClause":     reflect.TypeOf((*q.CaseClause)(nil)).Elem(),
			"ChanDir":        reflect.TypeOf((*q.ChanDir)(nil)).Elem(),
			"ChanType":       reflect.TypeOf((*q.ChanType)(nil)).Elem(),
			"CommClause":     reflect.TypeOf((*q.CommClause)(nil)).Elem(),
			"Comment":        reflect.TypeOf((*q.Comment)(nil)).Elem(),
			"CommentGroup":   reflect.TypeOf((*q.CommentGroup)(nil)).Elem(),
			"CommentMap":     reflect.TypeOf((*q.CommentMap)(nil)).Elem(),
			"CompositeLit":   reflect.TypeOf((*q.CompositeLit)(nil)).Elem(),
			"DeclStmt":       reflect.TypeOf((*q.DeclStmt)(nil)).Elem(),
			"DeferStmt":      reflect.TypeOf((*q.DeferStmt)(nil)).Elem(),
			"Ellipsis":       reflect.TypeOf((*q.Ellipsis)(nil)).Elem(),
			"EmptyStmt":      reflect.TypeOf((*q.EmptyStmt)(nil)).Elem(),
			"ExprStmt":       reflect.TypeOf((*q.ExprStmt)(nil)).Elem(),
			"Field":          reflect.TypeOf((*q.Field)(nil)).Elem(),
			"FieldFilter":    reflect.TypeOf((*q.FieldFilter)(nil)).Elem(),
			"FieldList":      reflect.TypeOf((*q.FieldList)(nil)).Elem(),
			"File":           reflect.TypeOf((*q.File)(nil)).Elem(),
			"Filter":         reflect.TypeOf((*q.Filter)(nil)).Elem(),
			"ForStmt":        reflect.TypeOf((*q.ForStmt)(nil)).Elem(),
			"FuncDecl":       reflect.TypeOf((*q.FuncDecl)(nil)).Elem(),
			"FuncLit":        reflect.TypeOf((*q.FuncLit)(nil)).Elem(),
			"FuncType":       reflect.TypeOf((*q.FuncType)(nil)).Elem(),
			"GenDecl":        reflect.TypeOf((*q.GenDecl)(nil)).Elem(),
			"GoStmt":         reflect.TypeOf((*q.GoStmt)(nil)).Elem(),
			"Ident":          reflect.TypeOf((*q.Ident)(nil)).Elem(),
			"IfStmt":         reflect.TypeOf((*q.IfStmt)(nil)).Elem(),
			"ImportSpec":     reflect.TypeOf((*q.ImportSpec)(nil)).Elem(),
			"Importer":       reflect.TypeOf((*q.Importer)(nil)).Elem(),
			"IncDecStmt":     reflect.TypeOf((*q.IncDecStmt)(nil)).Elem(),
			"IndexExpr":      reflect.TypeOf((*q.IndexExpr)(nil)).Elem(),
			"InterfaceType":  reflect.TypeOf((*q.InterfaceType)(nil)).Elem(),
			"KeyValueExpr":   reflect.TypeOf((*q.KeyValueExpr)(nil)).Elem(),
			"LabeledStmt":    reflect.TypeOf((*q.LabeledStmt)(nil)).Elem(),
			"MapType":        reflect.TypeOf((*q.MapType)(nil)).Elem(),
			"MergeMode":      reflect.TypeOf((*q.MergeMode)(nil)).Elem(),
			"ObjKind":        reflect.TypeOf((*q.ObjKind)(nil)).Elem(),
			"Object":         reflect.TypeOf((*q.Object)(nil)).Elem(),
			"Package":        reflect.TypeOf((*q.Package)(nil)).Elem(),
			"ParenExpr":      reflect.TypeOf((*q.ParenExpr)(nil)).Elem(),
			"RangeStmt":      reflect.TypeOf((*q.RangeStmt)(nil)).Elem(),
			"ReturnStmt":     reflect.TypeOf((*q.ReturnStmt)(nil)).Elem(),
			"Scope":          reflect.TypeOf((*q.Scope)(nil)).Elem(),
			"SelectStmt":     reflect.TypeOf((*q.SelectStmt)(nil)).Elem(),
			"SelectorExpr":   reflect.TypeOf((*q.SelectorExpr)(nil)).Elem(),
			"SendStmt":       reflect.TypeOf((*q.SendStmt)(nil)).Elem(),
			"SliceExpr":      reflect.TypeOf((*q.SliceExpr)(nil)).Elem(),
			"StarExpr":       reflect.TypeOf((*q.StarExpr)(nil)).Elem(),
			"StructType":     reflect.TypeOf((*q.StructType)(nil)).Elem(),
			"SwitchStmt":     reflect.TypeOf((*q.SwitchStmt)(nil)).Elem(),
			"TypeAssertExpr": reflect.TypeOf((*q.TypeAssertExpr)(nil)).Elem(),
			"TypeSpec":       reflect.TypeOf((*q.TypeSpec)(nil)).Elem(),
			"TypeSwitchStmt": reflect.TypeOf((*q.TypeSwitchStmt)(nil)).Elem(),
			"UnaryExpr":      reflect.TypeOf((*q.UnaryExpr)(nil)).Elem(),
			"ValueSpec":      reflect.TypeOf((*q.ValueSpec)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"FileExports":       reflect.ValueOf(q.FileExports),
			"FilterDecl":        reflect.ValueOf(q.FilterDecl),
			"FilterFile":        reflect.ValueOf(q.FilterFile),
			"FilterPackage":     reflect.ValueOf(q.FilterPackage),
			"Fprint":            reflect.ValueOf(q.Fprint),
			"Inspect":           reflect.ValueOf(q.Inspect),
			"IsExported":        reflect.ValueOf(q.IsExported),
			"MergePackageFiles": reflect.ValueOf(q.MergePackageFiles),
			"NewCommentMap":     reflect.ValueOf(q.NewCommentMap),
			"NewIdent":          reflect.ValueOf(q.NewIdent),
			"NewObj":            reflect.ValueOf(q.NewObj),
			"NewPackage":        reflect.ValueOf(q.NewPackage),
			"NewScope":          reflect.ValueOf(q.NewScope),
			"NotNilFilter":      reflect.ValueOf(q.NotNilFilter),
			"PackageExports":    reflect.ValueOf(q.PackageExports),
			"Print":             reflect.ValueOf(q.Print),
			"SortImports":       reflect.ValueOf(q.SortImports),
			"Walk":              reflect.ValueOf(q.Walk),
		},
		TypedConsts: map[string]igop.TypedConst{
			"Bad":                        {reflect.TypeOf(q.Bad), constant.MakeInt64(int64(q.Bad))},
			"Con":                        {reflect.TypeOf(q.Con), constant.MakeInt64(int64(q.Con))},
			"FilterFuncDuplicates":       {reflect.TypeOf(q.FilterFuncDuplicates), constant.MakeInt64(int64(q.FilterFuncDuplicates))},
			"FilterImportDuplicates":     {reflect.TypeOf(q.FilterImportDuplicates), constant.MakeInt64(int64(q.FilterImportDuplicates))},
			"FilterUnassociatedComments": {reflect.TypeOf(q.FilterUnassociatedComments), constant.MakeInt64(int64(q.FilterUnassociatedComments))},
			"Fun":                        {reflect.TypeOf(q.Fun), constant.MakeInt64(int64(q.Fun))},
			"Lbl":                        {reflect.TypeOf(q.Lbl), constant.MakeInt64(int64(q.Lbl))},
			"Pkg":                        {reflect.TypeOf(q.Pkg), constant.MakeInt64(int64(q.Pkg))},
			"RECV":                       {reflect.TypeOf(q.RECV), constant.MakeInt64(int64(q.RECV))},
			"SEND":                       {reflect.TypeOf(q.SEND), constant.MakeInt64(int64(q.SEND))},
			"Typ":                        {reflect.TypeOf(q.Typ), constant.MakeInt64(int64(q.Typ))},
			"Var":                        {reflect.TypeOf(q.Var), constant.MakeInt64(int64(q.Var))},
		},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
