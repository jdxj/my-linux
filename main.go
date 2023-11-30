package main

import (
	"context"
	"os"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gproc"

	"github.com/jdxj/my-linux/internal/config"
	_ "github.com/jdxj/my-linux/internal/logic"
	cl "github.com/jdxj/my-linux/internal/logic/clear"
	"github.com/jdxj/my-linux/internal/logic/monitor"
	_ "github.com/jdxj/my-linux/internal/packed"
)

func main() {
	ctx := gctx.GetInitCtx()
	_, err := gcron.AddSingleton(ctx, config.Pattern, func(ctx context.Context) {
		monitor.Monitor(ctx)
	}, "monitor")
	if err != nil {
		g.Log().Panicf(ctx, "Add err: %s", err)
	}

	_, err = gcron.AddSingleton(ctx, config.PatternClear, func(ctx context.Context) {
		cl.Clear(ctx)
	})
	if err != nil {
		g.Log().Panicf(ctx, "Add err: %s", err)
	}

	gproc.AddSigHandlerShutdown(func(sig os.Signal) {
		gcron.Stop()
	})
	gproc.Listen()
}
