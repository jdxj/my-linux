package cmd

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"github.com/jdxj/my-linux/internal/controller/hello"
	"github.com/jdxj/my-linux/internal/logic"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Bind(
					hello.NewV1(),
				)
			})
			s.Run()
			return nil
		},
	}

	Mem = &gcmd.Command{
		Name:        "mem",
		Usage:       "",
		Brief:       "",
		Description: "",
		Arguments: []gcmd.Argument{
			{
				Name: "pid",
			},
		},
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			pid := parser.GetOpt("pid")
			if pid.IsEmpty() {
				return gerror.Newf("pid not set")
			}

			logic.PrintVmRSS(ctx, pid.Int32())
			return nil
		},
	}
)
