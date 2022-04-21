package util

import (
	"github.com/ctrsploit/ctrsploit/util/godl"
	"os"
	"path/filepath"
)

func GetLibraryPath(library string, rPath []string, runPath []string) (libPath string) {
	for _, dir := range append(append(rPath, runPath...), godl.GetDynLibDirs()...) {
		// does this .so file exist? if so, recursively treat it
		// as another binary:
		so := filepath.Join(dir, library)
		fi, err := os.Stat(so)
		if err == nil && !fi.IsDir() {
			libPath = so
		}
	}
	return
}
