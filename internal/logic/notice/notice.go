package notice

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"

	"github.com/jdxj/my-linux/internal/config"
	"github.com/jdxj/my-linux/internal/model"
	"github.com/jdxj/my-linux/internal/service"
)

func init() {
	service.RegisterNotice(newSNotice())
}

func newSNotice() *sNotice {
	ba, err := tgbotapi.NewBotAPI(config.TG.Token)
	if err != nil {
		g.Log().Panicf(gctx.GetInitCtx(), "NewBotAPI err: %s", err)
	}
	return &sNotice{
		botAPI: ba,
	}
}

type sNotice struct {
	botAPI *tgbotapi.BotAPI
}

func (s *sNotice) SendNoticeByTG(ctx context.Context, in *model.SendNoticeInput) (*model.SendNoticeOutput, error) {
	mc := tgbotapi.NewMessage(config.TG.ChatId, in.Content)
	_, err := s.botAPI.Send(mc)
	return nil, err
}
