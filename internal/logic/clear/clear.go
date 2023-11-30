package clear

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/gogf/gf/v2/frame/g"

	"github.com/jdxj/my-linux/internal/config"
	"github.com/jdxj/my-linux/internal/model"
	"github.com/jdxj/my-linux/internal/service"
)

func Clear(ctx context.Context) {
	for _, threshold := range config.Clear {
		check(ctx, threshold)
	}
}

func check(ctx context.Context, threshold *config.Threshold) {
	var hasRm bool
	for _, name := range threshold.Files {
		rm, err := checkSize(threshold.Size, name)
		if err != nil {
			_, err = service.Notice().SendNoticeByTG(ctx, &model.SendNoticeInput{
				Content: fmt.Sprintf("rm %s err: %s", name, err),
			})
			if err != nil {
				g.Log().Errorf(ctx, "SendNoticeByTG err: %s", err)
			}
			continue
		}
		if rm {
			hasRm = true
		}
	}

	if threshold.Restart != "" && hasRm {
		restart(ctx, threshold.Restart)
	}
}

func checkSize(size int64, name string) (bool, error) {
	file, err := os.Stat(name)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false, nil
		}
		return false, err
	}
	if file.Size() < size {
		return false, nil
	}
	err = os.Remove(name)
	if err != nil {
		return false, err
	}
	return true, nil
}

func restart(ctx context.Context, cli string) {
	c := strings.Split(cli, " ")
	cmd := exec.Command(c[0], c[1:]...)
	res, err := cmd.CombinedOutput()
	if err != nil {
		_, err = service.Notice().SendNoticeByTG(ctx, &model.SendNoticeInput{
			Content: fmt.Sprintf("restart %s err: %s", c[0], err),
		})
		if err != nil {
			g.Log().Errorf(ctx, "SendNoticeByTG err: %s", err)
		}
	}
	g.Log().Infof(ctx, "cmd %s output: %s", c[0], res)
}
