package config

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
)

var (
	w = wrap{
		TG: &TG,
	}
	TG tg
)

type wrap struct {
	TG *tg
}

type tg struct {
	Token  string
	ChatId int64
}

func init() {
	ctx := gctx.GetInitCtx()
	m := g.Cfg().MustGet(ctx, "my-linux").Map()
	err := gconv.Struct(m, &w)
	if err != nil {
		g.Log().Panicf(ctx, "Struct err: %s", err)
	}
}
