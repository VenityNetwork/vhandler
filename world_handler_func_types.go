package vhandler

import (
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/go-gl/mathgl/mgl64"
)

type WorldHandleLiquidFlowFunc func(ctx *world.Context, from cube.Pos, into cube.Pos, liquid world.Liquid, replaced world.Block)
type WorldHandleLiquidDecayFunc func(ctx *world.Context, pos cube.Pos, before world.Liquid, after world.Liquid)
type WorldHandleLiquidHardenFunc func(ctx *world.Context, hardenedPos cube.Pos, liquidHardened world.Block, otherLiquid world.Block, newBlock world.Block)
type WorldHandleSoundFunc func(ctx *world.Context, s world.Sound, pos mgl64.Vec3)
type WorldHandleFireSpreadFunc func(ctx *world.Context, from cube.Pos, to cube.Pos)
type WorldHandleBlockBurnFunc func(ctx *world.Context, pos cube.Pos)
type WorldHandleCropTrampleFunc func(ctx *world.Context, pos cube.Pos)
type WorldHandleLeavesDecayFunc func(ctx *world.Context, pos cube.Pos)
type WorldHandleEntitySpawnFunc func(tx *world.Tx, e world.Entity)
type WorldHandleEntityDespawnFunc func(tx *world.Tx, e world.Entity)
type WorldHandleCloseFunc func(tx *world.Tx)
