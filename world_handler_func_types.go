package vhandler

import (
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/event"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/go-gl/mathgl/mgl64"
)

type WorldHandleLiquidFlowFunc func(w *world.World, ctx *event.Context, from cube.Pos, into cube.Pos, liquid world.Liquid, replaced world.Block)
type WorldHandleLiquidDecayFunc func(w *world.World, ctx *event.Context, pos cube.Pos, before world.Liquid, after world.Liquid)
type WorldHandleLiquidHardenFunc func(w *world.World, ctx *event.Context, hardenedPos cube.Pos, liquidHardened world.Block, otherLiquid world.Block, newBlock world.Block)
type WorldHandleSoundFunc func(w *world.World, ctx *event.Context, s world.Sound, pos mgl64.Vec3)
type WorldHandleFireSpreadFunc func(w *world.World, ctx *event.Context, from cube.Pos, to cube.Pos)
type WorldHandleBlockBurnFunc func(w *world.World, ctx *event.Context, pos cube.Pos)
type WorldHandleCropTrampleFunc func(w *world.World, ctx *event.Context, pos cube.Pos)
type WorldHandleEntitySpawnFunc func(w *world.World, e world.Entity)
type WorldHandleEntityDespawnFunc func(w *world.World, e world.Entity)
type WorldHandleCloseFunc func(w *world.World)
