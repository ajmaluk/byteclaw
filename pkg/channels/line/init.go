package line

import (
	"github.com/ajmaluk/byteclaw/pkg/bus"
	"github.com/ajmaluk/byteclaw/pkg/channels"
	"github.com/ajmaluk/byteclaw/pkg/config"
)

func init() {
	channels.RegisterFactory("line", func(cfg *config.Config, b *bus.MessageBus) (channels.Channel, error) {
		return NewLINEChannel(cfg.Channels.LINE, b)
	})
}
