package config

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
)

var (
	w = wrap{}

	TG           tg
	Pattern      string
	PatternClear string
	Monitor      map[string]*Threshold
	Clear        map[string]*Threshold
)

type wrap struct {
	TG           tg
	Pattern      string
	PatternClear string
	Monitor      map[string]*Threshold
	Clear        map[string]*Threshold
}

type tg struct {
	Token  string
	ChatId int64
}

type Threshold struct {
	Memory uint64

	Restart string
	Size    int64
	Files   []string
}

func init() {
	ctx := gctx.GetInitCtx()
	m := g.Cfg().MustGet(ctx, "my-linux").Map()
	err := gconv.Struct(m, &w)
	if err != nil {
		g.Log().Panicf(ctx, "Struct err: %s", err)
	}
	TG = w.TG
	Pattern = w.Pattern
	PatternClear = w.PatternClear
	Monitor = w.Monitor
	Clear = w.Clear
}
