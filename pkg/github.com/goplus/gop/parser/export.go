// export by github.com/goplus/igop/cmd/qexp

package parser

import (
	q "github.com/goplus/gop/parser"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "parser",
		Path: "github.com/goplus/gop/parser",
		Deps: map[string]string{
			"errors":                           "errors",
			"fmt":                              "fmt",
			"github.com/goplus/gop/ast":        "ast",
			"github.com/goplus/gop/parser/fsx": "fsx",
			"github.com/goplus/gop/parser/iox": "iox",
			"github.com/goplus/gop/scanner":    "scanner",
			"github.com/goplus/gop/token":      "token",
			"github.com/goplus/gop/tpl/ast":    "ast",
			"github.com/goplus/gop/tpl/parser": "parser",
			"github.com/qiniu/x/log":           "log",
			"go/ast":                           "ast",
			"go/parser":                        "parser",
			"go/scanner":                       "scanner",
			"io/fs":                            "fs",
			"path":                             "path",
			"strconv":                          "strconv",
			"strings":                          "strings",
			"unicode":                          "unicode",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Config": reflect.TypeOf((*q.Config)(nil)).Elem(),
			"Mode":   reflect.TypeOf((*q.Mode)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{
			"FileSystem": reflect.TypeOf((*q.FileSystem)(nil)).Elem(),
		},
		Vars: map[string]reflect.Value{
			"ErrUnknownFileKind": reflect.ValueOf(&q.ErrUnknownFileKind),
		},
		Funcs: map[string]reflect.Value{
			"Parse":          reflect.ValueOf(q.Parse),
			"ParseDir":       reflect.ValueOf(q.ParseDir),
			"ParseDirEx":     reflect.ValueOf(q.ParseDirEx),
			"ParseEntries":   reflect.ValueOf(q.ParseEntries),
			"ParseEntry":     reflect.ValueOf(q.ParseEntry),
			"ParseExpr":      reflect.ValueOf(q.ParseExpr),
			"ParseExprEx":    reflect.ValueOf(q.ParseExprEx),
			"ParseExprFrom":  reflect.ValueOf(q.ParseExprFrom),
			"ParseFSDir":     reflect.ValueOf(q.ParseFSDir),
			"ParseFSEntries": reflect.ValueOf(q.ParseFSEntries),
			"ParseFSEntry":   reflect.ValueOf(q.ParseFSEntry),
			"ParseFSFile":    reflect.ValueOf(q.ParseFSFile),
			"ParseFSFiles":   reflect.ValueOf(q.ParseFSFiles),
			"ParseFile":      reflect.ValueOf(q.ParseFile),
			"ParseFiles":     reflect.ValueOf(q.ParseFiles),
			"SetDebug":       reflect.ValueOf(q.SetDebug),
		},
		TypedConsts: map[string]igop.TypedConst{
			"AllErrors":            {reflect.TypeOf(q.AllErrors), constant.MakeInt64(int64(q.AllErrors))},
			"DeclarationErrors":    {reflect.TypeOf(q.DeclarationErrors), constant.MakeInt64(int64(q.DeclarationErrors))},
			"ImportsOnly":          {reflect.TypeOf(q.ImportsOnly), constant.MakeInt64(int64(q.ImportsOnly))},
			"PackageClauseOnly":    {reflect.TypeOf(q.PackageClauseOnly), constant.MakeInt64(int64(q.PackageClauseOnly))},
			"ParseComments":        {reflect.TypeOf(q.ParseComments), constant.MakeInt64(int64(q.ParseComments))},
			"ParseGoAsGoPlus":      {reflect.TypeOf(q.ParseGoAsGoPlus), constant.MakeInt64(int64(q.ParseGoAsGoPlus))},
			"ParseGoPlusClass":     {reflect.TypeOf(q.ParseGoPlusClass), constant.MakeInt64(int64(q.ParseGoPlusClass))},
			"SaveAbsFile":          {reflect.TypeOf(q.SaveAbsFile), constant.MakeInt64(int64(q.SaveAbsFile))},
			"SkipObjectResolution": {reflect.TypeOf(q.SkipObjectResolution), constant.MakeInt64(int64(q.SkipObjectResolution))},
			"Trace":                {reflect.TypeOf(q.Trace), constant.MakeInt64(int64(q.Trace))},
		},
		UntypedConsts: map[string]igop.UntypedConst{
			"DbgFlagAll":         {"untyped int", constant.MakeInt64(int64(q.DbgFlagAll))},
			"DbgFlagParseError":  {"untyped int", constant.MakeInt64(int64(q.DbgFlagParseError))},
			"DbgFlagParseOutput": {"untyped int", constant.MakeInt64(int64(q.DbgFlagParseOutput))},
		},
	})
}
