// export by github.com/goplus/igop/cmd/qexp

//go:build go1.22 && !go1.23
// +build go1.22,!go1.23

package constraint

import (
	q "go/build/constraint"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "constraint",
		Path: "go/build/constraint",
		Deps: map[string]string{
			"errors":       "errors",
			"strconv":      "strconv",
			"strings":      "strings",
			"unicode":      "unicode",
			"unicode/utf8": "utf8",
		},
		Interfaces: map[string]reflect.Type{
			"Expr": reflect.TypeOf((*q.Expr)(nil)).Elem(),
		},
		NamedTypes: map[string]reflect.Type{
			"AndExpr":     reflect.TypeOf((*q.AndExpr)(nil)).Elem(),
			"NotExpr":     reflect.TypeOf((*q.NotExpr)(nil)).Elem(),
			"OrExpr":      reflect.TypeOf((*q.OrExpr)(nil)).Elem(),
			"SyntaxError": reflect.TypeOf((*q.SyntaxError)(nil)).Elem(),
			"TagExpr":     reflect.TypeOf((*q.TagExpr)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"GoVersion":      reflect.ValueOf(q.GoVersion),
			"IsGoBuild":      reflect.ValueOf(q.IsGoBuild),
			"IsPlusBuild":    reflect.ValueOf(q.IsPlusBuild),
			"Parse":          reflect.ValueOf(q.Parse),
			"PlusBuildLines": reflect.ValueOf(q.PlusBuildLines),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
