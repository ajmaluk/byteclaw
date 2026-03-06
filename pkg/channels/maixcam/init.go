package maixcam

import (
	"github.com/ajmaluk/byteclaw/pkg/bus"
	"github.com/ajmaluk/byteclaw/pkg/channels"
	"github.com/ajmaluk/byteclaw/pkg/config"
)

func init() {
	channels.RegisterFactory("maixcam", func(cfg *config.Config, b *bus.MessageBus) (channels.Channel, error) {
		return NewMaixCamChannel(cfg.Channels.MaixCam, b)
	})
}
