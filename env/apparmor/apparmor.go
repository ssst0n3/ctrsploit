package apparmor

import (
	"fmt"
	"github.com/ctrsploit/ctrsploit/internal"
	"github.com/ctrsploit/ctrsploit/pkg/apparmor"
	"github.com/ctrsploit/ctrsploit/pkg/lsm"
	"github.com/ctrsploit/sploit-spec/pkg/result"
	"github.com/ctrsploit/sploit-spec/pkg/result/item"
)

type Result struct {
	Name      result.Title
	Kernel    item.Bool  `json:"kernel"`
	Container item.Bool  `json:"container"`
	Profile   item.Short `json:"profile"`
	Mode      item.Short `json:"mode"`
}

func (r Result) String() (s string) {
	s += internal.Print(r.Name, r.Kernel, r.Container)
	if r.Container.Result {
		s += internal.Print(r.Profile, r.Mode)
	}
	return s
}

func Apparmor() (err error) {
	enabled := apparmor.IsEnabled()
	r := Result{
		Name: result.Title{
			Name: "AppArmor",
		},
		Kernel: item.Bool{
			Name:        "Kernel Supported",
			Description: "Kernel enabled apparmor module",
			Result:      apparmor.IsSupport(),
		},
		Container: item.Bool{
			Name:        "Container Enabled",
			Description: "Current container enabled apparmor",
			Result:      enabled,
		},
	}
	if enabled {
		current, err := lsm.Current()
		if err != nil {
			return err
		}
		mode, err := apparmor.Mode()
		if err != nil {
			return err
		}
		r.Profile = item.Short{
			Name:        "Profile",
			Description: "",
			Result:      current,
		}
		r.Mode = item.Short{
			Name:        "Mode",
			Description: "",
			Result:      mode,
		}
	}
	fmt.Println(r)
	return
}
