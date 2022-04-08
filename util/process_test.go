package util

import (
	"github.com/ssst0n3/awesome_libs/log"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"os/exec"
	"testing"
	"time"
)

func TestGetProcessNameByPid(t *testing.T) {
	expectedProcess := "cat"
	cmd := exec.Command(expectedProcess)
	assert.NoError(t, cmd.Start())
	processName, err := GetProcessNameByPid(cmd.Process.Pid)
	assert.NoError(t, err)
	assert.Equal(t, expectedProcess, processName)
}

func TestGetProcessPathFromEnvByPid(t *testing.T) {
	shebang := "/tmp/ctrsploit_TestGetProcessPathFromEnvByPid"
	assert.NoError(t, os.RemoveAll(shebang))
	assert.NoError(t, ioutil.WriteFile(shebang, []byte("#!/bin/bash\nsleep 10"), 0755))
	cmd := exec.Command("/bin/bash", "-c", shebang)
	assert.NoError(t, cmd.Start())
	time.Sleep(time.Second)
	path, err := GetProcessPathFromEnvByPid(cmd.Process.Pid)
	assert.NoError(t, err)
	assert.Equal(t, shebang, path)
	assert.NoError(t, os.RemoveAll(shebang))
}

func TestIsSheBang(t *testing.T) {
	shebang := "/tmp/ctrsploit_TestIsShebang"
	assert.NoError(t, os.RemoveAll(shebang))
	assert.NoError(t, ioutil.WriteFile(shebang, []byte("#!/bin/bash\nsleep 10"), 0755))
	cmd := exec.Command("/bin/sh", "-c", "sh -c "+shebang)
	assert.NoError(t, cmd.Start())
	time.Sleep(time.Second)
	isSheBang, err := IsSheBang(cmd.Process.Pid)
	assert.NoError(t, err)
	assert.True(t, len(isSheBang) > 0)
	assert.NoError(t, os.RemoveAll(shebang))
}

func TestGetLastArgInCmdLine(t *testing.T) {
	shebang := "/tmp/ctrsploit_TestGetLastArgInCmdLine"
	assert.NoError(t, os.RemoveAll(shebang))
	assert.NoError(t, ioutil.WriteFile(shebang, []byte("#!/bin/bash\nsleep 10"), 0755))
	cmd := exec.Command("/bin/bash", "-c", shebang+" sh")
	assert.NoError(t, cmd.Start())
	time.Sleep(time.Second)
	lastArg, err := GetLastArgInCmdLine(cmd.Process.Pid)
	assert.NoError(t, err)
	log.Logger.Info(lastArg)
	assert.NoError(t, os.RemoveAll(shebang))
}

func TestGetCmdline(t *testing.T) {
	shebang := "/tmp/ctrsploit_TestGetCmdline"
	assert.NoError(t, os.RemoveAll(shebang))
	assert.NoError(t, ioutil.WriteFile(shebang, []byte("#!/bin/bash\nsleep 10"), 0755))
	cmd := exec.Command("/bin/bash", "-c", shebang+" sh")
	assert.NoError(t, cmd.Start())
	time.Sleep(time.Second)
	cmdline, err := GetCmdline(cmd.Process.Pid)
	assert.NoError(t, err)
	assert.Equal(t, []string{"/bin/bash", "/tmp/ctrsploit_TestGetCmdline", "sh"}, cmdline)
	assert.NoError(t, os.RemoveAll(shebang))
}
