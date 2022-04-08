package util

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func TestCheckFileExists(t *testing.T) {
	assert.True(t, CheckPathExists("/etc/passwd"))
	assert.False(t, CheckPathExists("/not_exists"))
}

func TestReadIntFromFile(t *testing.T) {
	result, err := ReadIntFromFile("/proc/sys/kernel/pid_max")
	assert.NoError(t, err)
	assert.True(t, result > 0)
}

func TestReplaceContent(t *testing.T) {
	assert.NoError(t, ioutil.WriteFile("/tmp/replace_test", []byte("source"), 0755))
	assert.NoError(t, ReplaceContent("/tmp/replace_test", []byte("source"), []byte("dest")))
	content, err := ioutil.ReadFile("/tmp/replace_test")
	assert.NoError(t, err)
	assert.Equal(t, []byte("dest"), content)
}

func TestReadFirstTwoBytesOfFile(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		file := "/tmp/ctrsploit_TestReadFirstTwoBytesOfFile"
		assert.NoError(t, os.RemoveAll(file))
		content := []byte("#!")
		assert.NoError(t, ioutil.WriteFile(file, content, 0644))
		header, err := ReadFirstTwoBytesOfFile(file)
		assert.NoError(t, err)
		assert.Equal(t, content, header[:])
		assert.NoError(t, os.RemoveAll(file))
	})
}
