package notice

import (
	"testing"

	"github.com/gogf/gf/v2/os/gctx"

	"github.com/jdxj/my-linux/internal/model"
)

func TestSNotice_SendNoticeByTG(t *testing.T) {
	s := newSNotice()
	ctx := gctx.New()
	_, err := s.SendNoticeByTG(ctx, &model.SendNoticeInput{Content: "hello"})
	if err != nil {
		t.Fatalf("%s\n", err)
	}
}
