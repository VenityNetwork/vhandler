package vhandler

import (
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/go-gl/mathgl/mgl64"
)

type WorldNativeBridgeHandler struct {
	w *world.World
	h *WorldHandlers
}

// Compile-time check to ensure that the WorldNativeBridgeHandler implements the WorldHandler interface.
var _ world.Handler = &WorldNativeBridgeHandler{}

func (h *WorldNativeBridgeHandler) HandleLiquidFlow(ctx *world.Context, from cube.Pos, into cube.Pos, liquid world.Liquid, replaced world.Block) {
	for _, handler := range h.h.liquidFlowHandlers {
		handler.h(ctx, from, into, liquid, replaced)
	}
}

func (h *WorldNativeBridgeHandler) HandleLiquidDecay(ctx *world.Context, pos cube.Pos, before world.Liquid, after world.Liquid) {
	for _, handler := range h.h.liquidDecayHandlers {
		handler.h(ctx, pos, before, after)
	}
}

func (h *WorldNativeBridgeHandler) HandleLiquidHarden(ctx *world.Context, hardenedPos cube.Pos, liquidHardened world.Block, otherLiquid world.Block, newBlock world.Block) {
	for _, handler := range h.h.liquidHardenHandlers {
		handler.h(ctx, hardenedPos, liquidHardened, otherLiquid, newBlock)
	}
}

func (h *WorldNativeBridgeHandler) HandleSound(ctx *world.Context, s world.Sound, pos mgl64.Vec3) {
	for _, handler := range h.h.soundHandlers {
		handler.h(ctx, s, pos)
	}
}

func (h *WorldNativeBridgeHandler) HandleFireSpread(ctx *world.Context, from cube.Pos, to cube.Pos) {
	for _, handler := range h.h.fireSpreadHandlers {
		handler.h(ctx, from, to)
	}
}

func (h *WorldNativeBridgeHandler) HandleBlockBurn(ctx *world.Context, pos cube.Pos) {
	for _, handler := range h.h.blockBurnHandlers {
		handler.h(ctx, pos)
	}
}

func (h *WorldNativeBridgeHandler) HandleCropTrample(ctx *world.Context, pos cube.Pos) {
	for _, handler := range h.h.cropTrampleHandlers {
		handler.h(ctx, pos)
	}
}

func (h *WorldNativeBridgeHandler) HandleLeavesDecay(ctx *world.Context, pos cube.Pos) {
	for _, handler := range h.h.leavesDecayHandlers {
		handler.h(ctx, pos)
	}
}

func (h *WorldNativeBridgeHandler) HandleEntitySpawn(tx *world.Tx, e world.Entity) {
	for _, handler := range h.h.entitySpawnHandlers {
		handler.h(tx, e)
	}
}

func (h *WorldNativeBridgeHandler) HandleEntityDespawn(tx *world.Tx, e world.Entity) {
	for _, handler := range h.h.entityDespawnHandlers {
		handler.h(tx, e)
	}
}

func (h *WorldNativeBridgeHandler) HandleClose(tx *world.Tx) {
	for _, handler := range h.h.closeHandlers {
		handler.h(tx)
	}
}
