package pipe_primitive

import (
	"github.com/ctrsploit/ctrsploit/util"
	"os"
)

type Library struct {
	Name    string
	Path    string
	Payload []byte
}

type LibraryMap map[string]Library



func (l *Library) Init() (err error) {
	if l.Name != "" {
		l.Path = util.GetLibraryPath(l.Name, nil, nil)
	}
	_, err = os.Stat(l.Path)
	if os.IsNotExist(err) {
		err = nil
		l.Path, err = findProperLib()
		if err != nil {
			return
		}
	}
	return
}

func findProperLib() (path string, err error) {
	return
}
