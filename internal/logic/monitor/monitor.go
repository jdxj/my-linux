package monitor

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/shirou/gopsutil/v3/process"

	"github.com/jdxj/my-linux/internal/config"
	"github.com/jdxj/my-linux/internal/model"
	"github.com/jdxj/my-linux/internal/service"
)

func processes(ctx context.Context) (map[string]*process.Process, error) {
	ps, err := process.ProcessesWithContext(ctx)
	if err != nil {
		return nil, err
	}

	psMap := make(map[string]*process.Process)
	for _, p := range ps {
		name, err := p.Name()
		if err != nil {
			g.Log().Warningf(ctx, "Name err: %s", err)
			continue
		}
		psMap[name] = p
	}
	return psMap, nil
}

func findProcesses(ctx context.Context) (map[string]*process.Process, error) {
	psMap, err := processes(ctx)
	if err != nil {
		return nil, err
	}

	targetPsMap := make(map[string]*process.Process)
	for name := range config.Monitor {
		p, ok := psMap[name]
		if !ok {
			return nil, gerror.Newf("process %s not found", name)
		}
		targetPsMap[name] = p
	}
	return targetPsMap, nil
}

func check(ctx context.Context, name string, p *process.Process, threshold *config.Threshold) error {
	err := checkMemory(ctx, name, p, threshold)
	if err != nil {
		return err
	}
	return nil
}

func checkMemory(ctx context.Context, name string, p *process.Process, threshold *config.Threshold) error {
	if threshold.Memory <= 0 {
		return nil
	}

	mi, err := p.MemoryInfo()
	if err != nil {
		return err
	}
	g.Log().Debugf(ctx, "%s, RSS: %d", name, mi.RSS)

	if mi.RSS > threshold.Memory {
		return gerror.Newf("%s RSS exceeds threshold: %d > %d",
			name, mi.RSS, threshold.Memory)
	}
	return nil
}

func Monitor(ctx context.Context) {
	targetPsMap, err := findProcesses(ctx)
	if err != nil {
		g.Log().Errorf(ctx, "findProcesses err: %s", err)
		return
	}

	for name, p := range targetPsMap {
		threshold := config.Monitor[name]
		notice := check(ctx, name, p, threshold)
		if notice == nil {
			continue
		}

		_, err := service.Notice().SendNoticeByTG(ctx, &model.SendNoticeInput{
			Content: notice.Error(),
		})
		if err != nil {
			g.Log().Errorf(ctx, "SendNoticeByTG err: %s", err)
		}
	}
}
