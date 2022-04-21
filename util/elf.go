package util

import (
	"debug/elf"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/ssst0n3/awesome_libs/log"
)

func ReplaceELFDependencies(path string) (err error) {
	f, err := elf.Open(path)
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	symbols, err := f.ImportedLibraries()
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	log.Logger.Info(symbols)
	return
}