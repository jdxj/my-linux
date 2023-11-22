package main

import (
	_ "github.com/jdxj/my-linux/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"github.com/jdxj/my-linux/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
