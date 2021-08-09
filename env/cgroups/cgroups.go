package cgroups

import (
	"github.com/ctrsploit/ctrsploit/log"
	"github.com/ctrsploit/ctrsploit/pkg/cgroup"
	"github.com/ctrsploit/ctrsploit/util"
	"fmt"
)

const CommandCgroupsName = "cgroups"

func Version() (err error) {
	info := fmt.Sprintf("===========Cgroups=========\n")
	info += fmt.Sprintf("is cgroupv1: %v\n", util.ColorfulTickOrBallot(cgroup.IsCgroupV1()))
	info += fmt.Sprintf("is cgroupv2: %v", util.ColorfulTickOrBallot(cgroup.IsCgroupV2()))
	log.Logger.Info(info)
	return
}
