// export by github.com/goplus/igop/cmd/qexp

//+build go1.16,!go1.17

package tls

import (
	q "crypto/tls"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "tls",
		Path: "crypto/tls",
		Deps: map[string]string{
			"bytes":           "bytes",
			"container/list":  "list",
			"context":         "context",
			"crypto":          "crypto",
			"crypto/aes":      "aes",
			"crypto/cipher":   "cipher",
			"crypto/des":      "des",
			"crypto/ecdsa":    "ecdsa",
			"crypto/ed25519":  "ed25519",
			"crypto/elliptic": "elliptic",
			"crypto/hmac":     "hmac",
			"crypto/md5":      "md5",
			"crypto/rand":     "rand",
			"crypto/rc4":      "rc4",
			"crypto/rsa":      "rsa",
			"crypto/sha1":     "sha1",
			"crypto/sha256":   "sha256",
			"crypto/sha512":   "sha512",
			"crypto/subtle":   "subtle",
			"crypto/x509":     "x509",
			"encoding/pem":    "pem",
			"errors":          "errors",
			"fmt":             "fmt",
			"hash":            "hash",
			"internal/cpu":    "cpu",
			"io":              "io",
			"math/big":        "big",
			"net":             "net",
			"os":              "os",
			"runtime":         "runtime",
			"sort":            "sort",
			"strconv":         "strconv",
			"strings":         "strings",
			"sync":            "sync",
			"sync/atomic":     "atomic",
			"time":            "time",
			"vendor/golang.org/x/crypto/chacha20poly1305": "chacha20poly1305",
			"vendor/golang.org/x/crypto/cryptobyte":       "cryptobyte",
			"vendor/golang.org/x/crypto/curve25519":       "curve25519",
			"vendor/golang.org/x/crypto/hkdf":             "hkdf",
		},
		Interfaces: map[string]reflect.Type{
			"ClientSessionCache": reflect.TypeOf((*q.ClientSessionCache)(nil)).Elem(),
		},
		NamedTypes: map[string]reflect.Type{
			"Certificate":            reflect.TypeOf((*q.Certificate)(nil)).Elem(),
			"CertificateRequestInfo": reflect.TypeOf((*q.CertificateRequestInfo)(nil)).Elem(),
			"CipherSuite":            reflect.TypeOf((*q.CipherSuite)(nil)).Elem(),
			"ClientAuthType":         reflect.TypeOf((*q.ClientAuthType)(nil)).Elem(),
			"ClientHelloInfo":        reflect.TypeOf((*q.ClientHelloInfo)(nil)).Elem(),
			"ClientSessionState":     reflect.TypeOf((*q.ClientSessionState)(nil)).Elem(),
			"Config":                 reflect.TypeOf((*q.Config)(nil)).Elem(),
			"Conn":                   reflect.TypeOf((*q.Conn)(nil)).Elem(),
			"ConnectionState":        reflect.TypeOf((*q.ConnectionState)(nil)).Elem(),
			"CurveID":                reflect.TypeOf((*q.CurveID)(nil)).Elem(),
			"Dialer":                 reflect.TypeOf((*q.Dialer)(nil)).Elem(),
			"RecordHeaderError":      reflect.TypeOf((*q.RecordHeaderError)(nil)).Elem(),
			"RenegotiationSupport":   reflect.TypeOf((*q.RenegotiationSupport)(nil)).Elem(),
			"SignatureScheme":        reflect.TypeOf((*q.SignatureScheme)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"CipherSuiteName":          reflect.ValueOf(q.CipherSuiteName),
			"CipherSuites":             reflect.ValueOf(q.CipherSuites),
			"Client":                   reflect.ValueOf(q.Client),
			"Dial":                     reflect.ValueOf(q.Dial),
			"DialWithDialer":           reflect.ValueOf(q.DialWithDialer),
			"InsecureCipherSuites":     reflect.ValueOf(q.InsecureCipherSuites),
			"Listen":                   reflect.ValueOf(q.Listen),
			"LoadX509KeyPair":          reflect.ValueOf(q.LoadX509KeyPair),
			"NewLRUClientSessionCache": reflect.ValueOf(q.NewLRUClientSessionCache),
			"NewListener":              reflect.ValueOf(q.NewListener),
			"Server":                   reflect.ValueOf(q.Server),
			"X509KeyPair":              reflect.ValueOf(q.X509KeyPair),
		},
		TypedConsts: map[string]igop.TypedConst{
			"CurveP256":                                     {reflect.TypeOf(q.CurveP256), constant.MakeInt64(int64(q.CurveP256))},
			"CurveP384":                                     {reflect.TypeOf(q.CurveP384), constant.MakeInt64(int64(q.CurveP384))},
			"CurveP521":                                     {reflect.TypeOf(q.CurveP521), constant.MakeInt64(int64(q.CurveP521))},
			"ECDSAWithP256AndSHA256":                        {reflect.TypeOf(q.ECDSAWithP256AndSHA256), constant.MakeInt64(int64(q.ECDSAWithP256AndSHA256))},
			"ECDSAWithP384AndSHA384":                        {reflect.TypeOf(q.ECDSAWithP384AndSHA384), constant.MakeInt64(int64(q.ECDSAWithP384AndSHA384))},
			"ECDSAWithP521AndSHA512":                        {reflect.TypeOf(q.ECDSAWithP521AndSHA512), constant.MakeInt64(int64(q.ECDSAWithP521AndSHA512))},
			"ECDSAWithSHA1":                                 {reflect.TypeOf(q.ECDSAWithSHA1), constant.MakeInt64(int64(q.ECDSAWithSHA1))},
			"Ed25519":                                       {reflect.TypeOf(q.Ed25519), constant.MakeInt64(int64(q.Ed25519))},
			"NoClientCert":                                  {reflect.TypeOf(q.NoClientCert), constant.MakeInt64(int64(q.NoClientCert))},
			"PKCS1WithSHA1":                                 {reflect.TypeOf(q.PKCS1WithSHA1), constant.MakeInt64(int64(q.PKCS1WithSHA1))},
			"PKCS1WithSHA256":                               {reflect.TypeOf(q.PKCS1WithSHA256), constant.MakeInt64(int64(q.PKCS1WithSHA256))},
			"PKCS1WithSHA384":                               {reflect.TypeOf(q.PKCS1WithSHA384), constant.MakeInt64(int64(q.PKCS1WithSHA384))},
			"PKCS1WithSHA512":                               {reflect.TypeOf(q.PKCS1WithSHA512), constant.MakeInt64(int64(q.PKCS1WithSHA512))},
			"PSSWithSHA256":                                 {reflect.TypeOf(q.PSSWithSHA256), constant.MakeInt64(int64(q.PSSWithSHA256))},
			"PSSWithSHA384":                                 {reflect.TypeOf(q.PSSWithSHA384), constant.MakeInt64(int64(q.PSSWithSHA384))},
			"PSSWithSHA512":                                 {reflect.TypeOf(q.PSSWithSHA512), constant.MakeInt64(int64(q.PSSWithSHA512))},
			"RenegotiateFreelyAsClient":                     {reflect.TypeOf(q.RenegotiateFreelyAsClient), constant.MakeInt64(int64(q.RenegotiateFreelyAsClient))},
			"RenegotiateNever":                              {reflect.TypeOf(q.RenegotiateNever), constant.MakeInt64(int64(q.RenegotiateNever))},
			"RenegotiateOnceAsClient":                       {reflect.TypeOf(q.RenegotiateOnceAsClient), constant.MakeInt64(int64(q.RenegotiateOnceAsClient))},
			"RequestClientCert":                             {reflect.TypeOf(q.RequestClientCert), constant.MakeInt64(int64(q.RequestClientCert))},
			"RequireAndVerifyClientCert":                    {reflect.TypeOf(q.RequireAndVerifyClientCert), constant.MakeInt64(int64(q.RequireAndVerifyClientCert))},
			"RequireAnyClientCert":                          {reflect.TypeOf(q.RequireAnyClientCert), constant.MakeInt64(int64(q.RequireAnyClientCert))},
			"TLS_AES_128_GCM_SHA256":                        {reflect.TypeOf(q.TLS_AES_128_GCM_SHA256), constant.MakeInt64(int64(q.TLS_AES_128_GCM_SHA256))},
			"TLS_AES_256_GCM_SHA384":                        {reflect.TypeOf(q.TLS_AES_256_GCM_SHA384), constant.MakeInt64(int64(q.TLS_AES_256_GCM_SHA384))},
			"TLS_CHACHA20_POLY1305_SHA256":                  {reflect.TypeOf(q.TLS_CHACHA20_POLY1305_SHA256), constant.MakeInt64(int64(q.TLS_CHACHA20_POLY1305_SHA256))},
			"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA":          {reflect.TypeOf(q.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA), constant.MakeInt64(int64(q.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA))},
			"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256":       {reflect.TypeOf(q.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256), constant.MakeInt64(int64(q.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256))},
			"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256":       {reflect.TypeOf(q.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256), constant.MakeInt64(int64(q.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256))},
			"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA":          {reflect.TypeOf(q.TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA), constant.MakeInt64(int64(q.TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA))},
			"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384":       {reflect.TypeOf(q.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384), constant.MakeInt64(int64(q.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384))},
			"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305":        {reflect.TypeOf(q.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305), constant.MakeInt64(int64(q.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305))},
			"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256": {reflect.TypeOf(q.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256), constant.MakeInt64(int64(q.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256))},
			"TLS_ECDHE_ECDSA_WITH_RC4_128_SHA":              {reflect.TypeOf(q.TLS_ECDHE_ECDSA_WITH_RC4_128_SHA), constant.MakeInt64(int64(q.TLS_ECDHE_ECDSA_WITH_RC4_128_SHA))},
			"TLS_ECDHE_RSA_WITH_3DES_EDE_CBC_SHA":           {reflect.TypeOf(q.TLS_ECDHE_RSA_WITH_3DES_EDE_CBC_SHA), constant.MakeInt64(int64(q.TLS_ECDHE_RSA_WITH_3DES_EDE_CBC_SHA))},
			"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA":            {reflect.TypeOf(q.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA), constant.MakeInt64(int64(q.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA))},
			"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256":         {reflect.TypeOf(q.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256), constant.MakeInt64(int64(q.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256))},
			"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256":         {reflect.TypeOf(q.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256), constant.MakeInt64(int64(q.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256))},
			"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA":            {reflect.TypeOf(q.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA), constant.MakeInt64(int64(q.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA))},
			"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384":         {reflect.TypeOf(q.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384), constant.MakeInt64(int64(q.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384))},
			"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305":          {reflect.TypeOf(q.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305), constant.MakeInt64(int64(q.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305))},
			"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256":   {reflect.TypeOf(q.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256), constant.MakeInt64(int64(q.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256))},
			"TLS_ECDHE_RSA_WITH_RC4_128_SHA":                {reflect.TypeOf(q.TLS_ECDHE_RSA_WITH_RC4_128_SHA), constant.MakeInt64(int64(q.TLS_ECDHE_RSA_WITH_RC4_128_SHA))},
			"TLS_FALLBACK_SCSV":                             {reflect.TypeOf(q.TLS_FALLBACK_SCSV), constant.MakeInt64(int64(q.TLS_FALLBACK_SCSV))},
			"TLS_RSA_WITH_3DES_EDE_CBC_SHA":                 {reflect.TypeOf(q.TLS_RSA_WITH_3DES_EDE_CBC_SHA), constant.MakeInt64(int64(q.TLS_RSA_WITH_3DES_EDE_CBC_SHA))},
			"TLS_RSA_WITH_AES_128_CBC_SHA":                  {reflect.TypeOf(q.TLS_RSA_WITH_AES_128_CBC_SHA), constant.MakeInt64(int64(q.TLS_RSA_WITH_AES_128_CBC_SHA))},
			"TLS_RSA_WITH_AES_128_CBC_SHA256":               {reflect.TypeOf(q.TLS_RSA_WITH_AES_128_CBC_SHA256), constant.MakeInt64(int64(q.TLS_RSA_WITH_AES_128_CBC_SHA256))},
			"TLS_RSA_WITH_AES_128_GCM_SHA256":               {reflect.TypeOf(q.TLS_RSA_WITH_AES_128_GCM_SHA256), constant.MakeInt64(int64(q.TLS_RSA_WITH_AES_128_GCM_SHA256))},
			"TLS_RSA_WITH_AES_256_CBC_SHA":                  {reflect.TypeOf(q.TLS_RSA_WITH_AES_256_CBC_SHA), constant.MakeInt64(int64(q.TLS_RSA_WITH_AES_256_CBC_SHA))},
			"TLS_RSA_WITH_AES_256_GCM_SHA384":               {reflect.TypeOf(q.TLS_RSA_WITH_AES_256_GCM_SHA384), constant.MakeInt64(int64(q.TLS_RSA_WITH_AES_256_GCM_SHA384))},
			"TLS_RSA_WITH_RC4_128_SHA":                      {reflect.TypeOf(q.TLS_RSA_WITH_RC4_128_SHA), constant.MakeInt64(int64(q.TLS_RSA_WITH_RC4_128_SHA))},
			"VerifyClientCertIfGiven":                       {reflect.TypeOf(q.VerifyClientCertIfGiven), constant.MakeInt64(int64(q.VerifyClientCertIfGiven))},
			"X25519":                                        {reflect.TypeOf(q.X25519), constant.MakeInt64(int64(q.X25519))},
		},
		UntypedConsts: map[string]igop.UntypedConst{
			"VersionSSL30": {"untyped int", constant.MakeInt64(int64(q.VersionSSL30))},
			"VersionTLS10": {"untyped int", constant.MakeInt64(int64(q.VersionTLS10))},
			"VersionTLS11": {"untyped int", constant.MakeInt64(int64(q.VersionTLS11))},
			"VersionTLS12": {"untyped int", constant.MakeInt64(int64(q.VersionTLS12))},
			"VersionTLS13": {"untyped int", constant.MakeInt64(int64(q.VersionTLS13))},
		},
	})
}
