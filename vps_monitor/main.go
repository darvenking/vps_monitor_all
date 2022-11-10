package main

import (
	"github.com/gogf/gf/v2/os/gctx"
	"vps_monitor/internal/cmd"
	_ "vps_monitor/internal/model"
	_ "vps_monitor/utility/cfg"
)

func main() {
	cmd.Main.Run(gctx.New())
}
