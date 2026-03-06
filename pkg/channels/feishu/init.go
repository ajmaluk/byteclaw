package feishu

import (
	"github.com/ajmaluk/byteclaw/pkg/bus"
	"github.com/ajmaluk/byteclaw/pkg/channels"
	"github.com/ajmaluk/byteclaw/pkg/config"
)

func init() {
	channels.RegisterFactory("feishu", func(cfg *config.Config, b *bus.MessageBus) (channels.Channel, error) {
		return NewFeishuChannel(cfg.Channels.Feishu, b)
	})
}
