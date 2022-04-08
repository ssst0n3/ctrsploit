package util

import (
	"bytes"
	"fmt"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"io"
	"io/ioutil"
	"os"
	"strconv"
)

func CheckPathExists(path string) bool {
	_, err := os.Lstat(path)
	return !os.IsNotExist(err)
}

func ReadIntFromFile(path string) (result int, err error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	content = bytes.TrimSpace(content)
	result, err = strconv.Atoi(string(content))
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	return
}

func ReplaceContent(path string, source, dest []byte) (err error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	err = ioutil.WriteFile(path, bytes.Replace(content, source, dest, -1), 0)
	if err != nil {
		awesome_error.CheckWarning(err)
		return
	}
	return
}

func ReadFirstTwoBytesOfFile(file string) (header [2]byte, err error) {
	r, err := os.Open(file)
	if err != nil {
		if os.IsNotExist(err) {
			return
		}
		awesome_error.CheckErr(err)
		return
	}
	defer r.Close()
	n, err := io.ReadFull(r, header[:])
	if err != nil {
		return
	}
	if n < 2 {
		err = fmt.Errorf("the size of file %s = %d < 2", file, n)
	}
	return
}
