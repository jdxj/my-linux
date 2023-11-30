package clear

import (
	"testing"

	"github.com/gogf/gf/v2/os/gctx"

	"github.com/jdxj/my-linux/internal/config"
)

func TestCheck(t *testing.T) {
	ctx := gctx.New()
	check(ctx, config.Clear["test"])
}
