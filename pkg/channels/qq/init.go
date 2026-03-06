package qq

import (
	"github.com/ajmaluk/byteclaw/pkg/bus"
	"github.com/ajmaluk/byteclaw/pkg/channels"
	"github.com/ajmaluk/byteclaw/pkg/config"
)

func init() {
	channels.RegisterFactory("qq", func(cfg *config.Config, b *bus.MessageBus) (channels.Channel, error) {
		return NewQQChannel(cfg.Channels.QQ, b)
	})
}
