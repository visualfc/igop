// export by github.com/goplus/interp/cmd/qexp

// +build freebsd netbsd openbsd

package syscall

import (
	"syscall"

	"github.com/goplus/interp"
)

func init() {
	interp.RegisterPackage("syscall", extMap_3244060493, typList_3244060493)
}

var extMap_3244060493 = map[string]interface{}{}

var typList_3244060493 = []interface{}{
	(*syscall.IfAnnounceMsghdr)(nil),
	(*syscall.InterfaceAnnounceMessage)(nil),
}