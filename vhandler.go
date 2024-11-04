package vhandler

import (
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
)

func NewPlayerHandlers() *PlayerHandlers {
	return &PlayerHandlers{}
}

func NewWorldHandlers() *WorldHandlers {
	return &WorldHandlers{}
}

func HandlePlayer(p *player.Player, h *PlayerHandlers) {
	p.Handle(&PlayerNativeBridgeHandler{p: p, h: h})
}

func HandleWorld(w *world.World, h *WorldHandlers) {
	w.Handle(&WorldNativeBridgeHandler{w: w, h: h})
}

func Test() {

}
