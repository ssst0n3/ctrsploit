/*
Package godl
borrowed from https://github.com/kontsevoy/godl
I don't want to copy the code, but these codes are under main package, I cannot import that.
*/
package godl

import (
	"bufio"
	"debug/elf"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func GetDynLibDirs() []string {
	dirs, err := ParseDynLibConf("/etc/ld.so.conf")
	if err != nil {
		awesome_error.CheckWarning(err)
	}
	return append(dirs, "/usr/lib", "/lib")
}

// ParseDynLibConf reads/parses DL config files defined as a pattern
// and returns a list of directories found in there (or an error).
func ParseDynLibConf(pattern string) (dirs []string, err error) {
	files := GlobMany([]string{pattern}, nil)

	for _, configFile := range files {
		fd, err := os.Open(configFile)
		if err != nil {
			return dirs, err
		}
		defer fd.Close()

		sc := bufio.NewScanner(fd)
		for sc.Scan() {
			line := strings.TrimSpace(sc.Text())
			// ignore comments and empty lines
			if len(line) == 0 || line[0] == '#' || line[0] == ';' {
				continue
			}
			// found "include" directive?
			words := strings.Fields(line)
			if strings.ToLower(words[0]) == "include" {
				subdirs, err := ParseDynLibConf(words[1])
				if err != nil && !os.IsNotExist(err) {
					return dirs, err
				}
				dirs = append(dirs, subdirs...)
			} else {
				dirs = append(dirs, line)
			}
		}
	}
	return dirs, err
}

// GlobMany takes a search pattern and returns absolute file paths that mach that
// pattern.
//	 - targets : list of paths to glob
//   - mask    : GlobDirs or GlobFiles
//   - onErr   : callback function to call when there's an error.
//                can be nil.
func GlobMany(targets []string, onErr func(string, error)) []string {
	rv := make([]string, 0, 20)
	addFile := func(path string, fi os.FileInfo, err error) error {
		if err != nil {
			log.Println(err.Error())
			return err
		}
		rv = append(rv, path)
		return err
	}

	for _, p := range targets {
		// "p" is a wildcard pattern? expand it:
		if strings.Contains(p, "*") {
			matches, err := filepath.Glob(p)
			if err == nil {
				// walk each match:
				for _, p := range matches {
					filepath.Walk(p, addFile)
				}
			}
			// path is not a wildcard, walk it:
		} else {
			filepath.Walk(p, addFile)
		}
	}
	return rv
}


func MyGetELFDependencies(target string, dlDirs []string, recursive bool) (retval []string) {
	var (
		onFile func(string)
		deps   map[string]bool = make(map[string]bool)
	)

	// gets called on every binary:
	onFile = func(fp string) {
		f, err := elf.Open(fp)
		if err != nil {
			return
		}
		defer f.Close()

		libs, err := f.ImportedLibraries()
		if err != nil {
			return
		}
		// check rpath and runpath
		rp1, _ := f.DynString(elf.DT_RPATH)
		rp2, _ := f.DynString(elf.DT_RUNPATH)

		// look for the lib in every location where dynamic linker would look:
		for _, lib := range libs {
			for _, dir := range append(append(rp1, rp2...), dlDirs...) {
				// does this .so file exist? if so, recursively treat it
				// as another binary:
				so := filepath.Join(dir, lib)
				fi, err := os.Stat(so)
				if err == nil && !fi.IsDir() {
					deps[so] = true
					if recursive {
						onFile(so)
					}
				}
			}
		}
	}

	// process command-line args (patterns/files):
	for _, p := range GlobMany([]string{target}, nil) {
		onFile(p)
	}

	// convert map values to slice of strings:
	retval = make([]string, 0, len(deps)/2)
	for p, _ := range deps {
		retval = append(retval, p)
	}
	return retval
}