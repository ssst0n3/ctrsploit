package godl

import (
	"github.com/ssst0n3/awesome_libs/log"
	"testing"
)

func TestGetDynLibDirs(t *testing.T) {
	log.Logger.Info(GetDynLibDirs())
}

func TestGetELFDependencies(t *testing.T) {
	log.Logger.Info(MyGetELFDependencies("/bin/bash", GetDynLibDirs(), false))
}