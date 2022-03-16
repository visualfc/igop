// export by github.com/goplus/gossa/cmd/qexp

//go:build go1.17 && !go1.18
// +build go1.17,!go1.18

package iotest

import (
	q "testing/iotest"

	"reflect"

	"github.com/goplus/gossa"
)

func init() {
	gossa.RegisterPackage(&gossa.Package{
		Name: "iotest",
		Path: "testing/iotest",
		Deps: map[string]string{
			"bytes":  "bytes",
			"errors": "errors",
			"fmt":    "fmt",
			"io":     "io",
			"log":    "log",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]gossa.NamedType{},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"ErrTimeout": reflect.ValueOf(&q.ErrTimeout),
		},
		Funcs: map[string]reflect.Value{
			"DataErrReader":  reflect.ValueOf(q.DataErrReader),
			"ErrReader":      reflect.ValueOf(q.ErrReader),
			"HalfReader":     reflect.ValueOf(q.HalfReader),
			"NewReadLogger":  reflect.ValueOf(q.NewReadLogger),
			"NewWriteLogger": reflect.ValueOf(q.NewWriteLogger),
			"OneByteReader":  reflect.ValueOf(q.OneByteReader),
			"TestReader":     reflect.ValueOf(q.TestReader),
			"TimeoutReader":  reflect.ValueOf(q.TimeoutReader),
			"TruncateWriter": reflect.ValueOf(q.TruncateWriter),
		},
		TypedConsts:   map[string]gossa.TypedConst{},
		UntypedConsts: map[string]gossa.UntypedConst{},
	})
}
