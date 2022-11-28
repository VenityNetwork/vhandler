package vhandler

import (
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/event"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/go-gl/mathgl/mgl64"
)

type WorldHandler struct {
	world.Handler

	wh *World
}

func newWorldHandler(wh *World) *WorldHandler {
	return &WorldHandler{wh: wh}
}

type WorldLiquidFlowHandler func(ctx *event.Context, from, into cube.Pos, liquid world.Liquid, replaced world.Block)

func (w *WorldHandler) HandleLiquidFlow(ctx *event.Context, from, into cube.Pos, liquid world.Liquid, replaced world.Block) {
	w.wh.handlers[WorldLiquidFlowId].handle(func(h Handler) {
		h.(WorldLiquidFlowHandler)(ctx, from, into, liquid, replaced)
	})
}

type WorldLiquidDecayHandler func(ctx *event.Context, pos cube.Pos, before, after world.Liquid)

func (w *WorldHandler) HandleLiquidDecay(ctx *event.Context, pos cube.Pos, before, after world.Liquid) {
	w.wh.handlers[WorldLiquidDecayId].handle(func(h Handler) {
		h.(WorldLiquidDecayHandler)(ctx, pos, before, after)
	})
}

type WorldLiquidHardenHandler func(ctx *event.Context, hardenedPos cube.Pos, liquidHardened, otherLiquid, newBlock world.Block)

func (w *WorldHandler) HandleLiquidHarden(ctx *event.Context, hardenedPos cube.Pos, liquidHardened, otherLiquid, newBlock world.Block) {
	w.wh.handlers[WorldLiquidHardenId].handle(func(h Handler) {
		h.(WorldLiquidHardenHandler)(ctx, hardenedPos, liquidHardened, otherLiquid, newBlock)
	})
}

type WorldSoundHandler func(ctx *event.Context, s world.Sound, pos mgl64.Vec3)

func (w *WorldHandler) HandleSound(ctx *event.Context, s world.Sound, pos mgl64.Vec3) {
	w.wh.handlers[WorldSoundId].handle(func(h Handler) {
		h.(WorldSoundHandler)(ctx, s, pos)
	})
}

type WorldFireSpreadHandler func(ctx *event.Context, from, to cube.Pos)

func (w *WorldHandler) HandleFireSpread(ctx *event.Context, from, to cube.Pos) {
	w.wh.handlers[WorldFireSpreadId].handle(func(h Handler) {
		h.(WorldFireSpreadHandler)(ctx, from, to)
	})
}

type WorldBlockBurnHandler func(ctx *event.Context, pos cube.Pos)

func (w *WorldHandler) HandleBlockBurn(ctx *event.Context, pos cube.Pos) {
	w.wh.handlers[WorldBlockBurnId].handle(func(h Handler) {
		h.(WorldBlockBurnHandler)(ctx, pos)
	})
}

type WorldEntitySpawnHandler func(e world.Entity)

func (w *WorldHandler) HandleEntitySpawn(e world.Entity) {
	w.wh.handlers[WorldEntitySpawnId].handle(func(h Handler) {
		h.(WorldEntitySpawnHandler)(e)
	})
}

type WorldEntityDespawnHandler func(e world.Entity)

func (w *WorldHandler) HandleEntityDespawn(e world.Entity) {
	w.wh.handlers[WorldEntityDespawnId].handle(func(h Handler) {
		h.(WorldEntityDespawnHandler)(e)
	})
}

type WorldCloseHandler func()

func (w *WorldHandler) HandleClose() {
	w.wh.handlers[WorldCloseId].handle(func(h Handler) {
		h.(WorldCloseHandler)()
	})
}
