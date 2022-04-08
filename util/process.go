package util

import (
	"bytes"
	"fmt"
	"github.com/pkg/errors"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
)

func GetProcessNameByPid(pid int) (processName string, err error) {
	path, err := GetProcessPathByPid(pid)
	if err != nil {
		return
	}
	processName = filepath.Base(path)
	return
}

func GetProcessPathByPid(pid int) (path string, err error) {
	return getProcessPath(strconv.Itoa(pid))
}

func GetProcessPathFromEnvByPid(pid int) (path string, err error) {
	return getProcessPathFromEnv(strconv.Itoa(pid))
}

func getProcessPathFromEnv(pid string) (path string, err error) {
	content, err := ioutil.ReadFile(fmt.Sprintf("/proc/%s/environ", pid))
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	for _, env := range bytes.Split(content, []byte{0}) {
		if bytes.HasPrefix(env, []byte("_=")) {
			path = string(bytes.TrimPrefix(env, []byte("_=")))
			return
		}
	}
	return
}

func getProcessPath(pid string) (path string, err error) {
	exe := fmt.Sprintf("/proc/%s/exe", pid)
	path, err = filepath.EvalSymlinks(exe)
	if err != nil {
		if errors.Is(err, os.ErrPermission) || errors.Is(err, os.ErrNotExist) {
			return
		} else {
			awesome_error.CheckErr(err)
		}
		return
	}
	return
}

func getSelfPid() (pid int, err error) {
	path, err := filepath.EvalSymlinks("/proc/self")
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	process := strings.TrimPrefix(path, "/proc/")
	pid, err = strconv.Atoi(process)
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	return
}

func KillAll() (err error) {
	matches, err := filepath.Glob("/proc/[0-9]*")
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	selfPid, err := getSelfPid()
	if err != nil {
		return
	}
	for _, match := range matches {
		process := strings.TrimPrefix(match, "/proc/")
		pid, err := strconv.Atoi(process)
		if err != nil {
			awesome_error.CheckErr(err)
			continue
		}
		if pid == selfPid {
			continue
		}
		err = syscall.Kill(pid, syscall.Signal(9))
		if err != nil {
			awesome_error.CheckErr(err)
			continue
		}
	}
	return
}

func GetLastArgInCmdLine(pid int) (lastArg string, err error) {
	cmdline, err := ioutil.ReadFile(fmt.Sprintf("/proc/%d/cmdline", pid))
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	args := bytes.Split(cmdline, []byte{0})
	index := len(args) - 1
	if len(args) > 2 {
		index = len(args) - 2
	}
	lastArg = string(args[index])
	return
}

func IsSheBang(pid int) (shebang string, err error) {
	cmdline, err := GetCmdline(1)
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	if len(cmdline) <= 1 {
		return
	}
	for _, cmd := range cmdline[1:] {
		path := cmd
		if !strings.HasPrefix(cmd, "/") {
			path = fmt.Sprintf("/proc/%d/cwd/%s", pid, path)
		}
		header, err2 := ReadFirstTwoBytesOfFile(path)
		if err2 != nil {
			err = err2
			if os.IsNotExist(err) {
				continue
			}
			return
		}
		if header == [2]byte{'#', '!'} {
			realpath, err2 := filepath.EvalSymlinks(path)
			if err2 != nil {
				err = err2
				awesome_error.CheckErr(err)
				return
			}
			shebang = realpath
			return
		}
	}
	return
}

func GetCmdline(pid int) (cmdline []string, err error) {
	return getCmdline(strconv.Itoa(pid))
}

func getCmdline(pid string) (cmdline []string, err error) {
	content, err := ioutil.ReadFile(fmt.Sprintf("/proc/%s/cmdline", pid))
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	cmdline = strings.Split(string(content), "\x00")
	if len(cmdline) > 0 {
		if cmdline[len(cmdline)-1] == "" {
			cmdline = cmdline[:len(cmdline)-1]
		}
	}
	return
}
