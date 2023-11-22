package logic

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/shirou/gopsutil/v3/process"
)

func PrintVmRSS(ctx context.Context, pid int32) {
	p, err := process.NewProcess(pid)
	if err != nil {
		g.Log().Panicf(ctx, "%s\n", err)
	}
	m, err := p.MemoryInfo()
	if err != nil {
		g.Log().Panicf(ctx, "%s\n", err)
	}
	fmt.Printf("rss: %d\n", m.RSS)
}
