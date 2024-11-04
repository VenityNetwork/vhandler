package vhandler

import (
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/event"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/go-gl/mathgl/mgl64"
)

type WorldNativeBridgeHandler struct {
	w *world.World
	h *WorldHandlers
}

// Compile-time check to ensure that the WorldNativeBridgeHandler implements the WorldHandler interface.
var _ world.Handler = &WorldNativeBridgeHandler{}

func (h *WorldNativeBridgeHandler) HandleLiquidFlow(ctx *event.Context, from cube.Pos, into cube.Pos, liquid world.Liquid, replaced world.Block) {
	for _, handler := range h.h.liquidFlowHandlers {
		handler.h(h.w, ctx, from, into, liquid, replaced)
	}
}

func (h *WorldNativeBridgeHandler) HandleLiquidDecay(ctx *event.Context, pos cube.Pos, before world.Liquid, after world.Liquid) {
	for _, handler := range h.h.liquidDecayHandlers {
		handler.h(h.w, ctx, pos, before, after)
	}
}

func (h *WorldNativeBridgeHandler) HandleLiquidHarden(ctx *event.Context, hardenedPos cube.Pos, liquidHardened world.Block, otherLiquid world.Block, newBlock world.Block) {
	for _, handler := range h.h.liquidHardenHandlers {
		handler.h(h.w, ctx, hardenedPos, liquidHardened, otherLiquid, newBlock)
	}
}

func (h *WorldNativeBridgeHandler) HandleSound(ctx *event.Context, s world.Sound, pos mgl64.Vec3) {
	for _, handler := range h.h.soundHandlers {
		handler.h(h.w, ctx, s, pos)
	}
}

func (h *WorldNativeBridgeHandler) HandleFireSpread(ctx *event.Context, from cube.Pos, to cube.Pos) {
	for _, handler := range h.h.fireSpreadHandlers {
		handler.h(h.w, ctx, from, to)
	}
}

func (h *WorldNativeBridgeHandler) HandleBlockBurn(ctx *event.Context, pos cube.Pos) {
	for _, handler := range h.h.blockBurnHandlers {
		handler.h(h.w, ctx, pos)
	}
}

func (h *WorldNativeBridgeHandler) HandleCropTrample(ctx *event.Context, pos cube.Pos) {
	for _, handler := range h.h.cropTrampleHandlers {
		handler.h(h.w, ctx, pos)
	}
}

func (h *WorldNativeBridgeHandler) HandleEntitySpawn(e world.Entity) {
	for _, handler := range h.h.entitySpawnHandlers {
		handler.h(h.w, e)
	}
}

func (h *WorldNativeBridgeHandler) HandleEntityDespawn(e world.Entity) {
	for _, handler := range h.h.entityDespawnHandlers {
		handler.h(h.w, e)
	}
}

func (h *WorldNativeBridgeHandler) HandleClose() {
	for _, handler := range h.h.closeHandlers {
		handler.h(h.w)
	}
}
