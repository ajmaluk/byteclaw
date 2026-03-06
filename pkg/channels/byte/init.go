package byte

import (
	"github.com/ajmaluk/byteclaw/pkg/bus"
	"github.com/ajmaluk/byteclaw/pkg/channels"
	"github.com/ajmaluk/byteclaw/pkg/config"
)

func init() {
	channels.RegisterFactory("byte", func(cfg *config.Config, b *bus.MessageBus) (channels.Channel, error) {
		return NewByteChannel(cfg.Channels.Byte, b)
	})
}
