package main

import (
	_ "github.com/jdxj/my-linux/internal/packed"

	_ "github.com/jdxj/my-linux/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	"github.com/jdxj/my-linux/internal/cmd"
)

func main() {
	cmd.Main.AddCommand(cmd.Mem)
	cmd.Main.Run(gctx.GetInitCtx())
}
