package vhandler

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"net"
	"time"

	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/player/skin"
	"github.com/df-mc/dragonfly/server/session"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/go-gl/mathgl/mgl64"
)

type PlayerHandleMoveFunc func(ctx *player.Context, newPos mgl64.Vec3, newRot cube.Rotation)
type PlayerHandleJumpFunc func(p *player.Player)
type PlayerHandleTeleportFunc func(ctx *player.Context, pos mgl64.Vec3)
type PlayerHandleChangeWorldFunc func(before *world.World, after *world.World)
type PlayerHandleToggleSprintFunc func(ctx *player.Context, after bool)
type PlayerHandleToggleSneakFunc func(ctx *player.Context, after bool)
type PlayerHandleChatFunc func(ctx *player.Context, message *string)
type PlayerHandleFoodLossFunc func(ctx *player.Context, from int, to *int)
type PlayerHandleHealFunc func(ctx *player.Context, health *float64, src world.HealingSource)
type PlayerHandleHurtFunc func(ctx *player.Context, damage *float64, immune bool, attackImmunity *time.Duration, src world.DamageSource)
type PlayerHandleDeathFunc func(p *player.Player, src world.DamageSource, keepInv *bool)
type PlayerHandleRespawnFunc func(p *player.Player, pos *mgl64.Vec3, w **world.World)
type PlayerHandleSkinChangeFunc func(ctx *player.Context, skin *skin.Skin)
type PlayerHandleFireExtinguishFunc func(ctx *player.Context, pos cube.Pos)
type PlayerHandleStartBreakFunc func(ctx *player.Context, pos cube.Pos)
type PlayerHandleBlockBreakFunc func(ctx *player.Context, pos cube.Pos, drops *[]item.Stack, xp *int)
type PlayerHandleBlockPlaceFunc func(ctx *player.Context, pos cube.Pos, b world.Block)
type PlayerHandleBlockPickFunc func(ctx *player.Context, pos cube.Pos, b world.Block)
type PlayerHandleItemUseFunc func(ctx *player.Context)
type PlayerHandleItemUseOnBlockFunc func(ctx *player.Context, pos cube.Pos, face cube.Face, clickPos mgl64.Vec3)
type PlayerHandleItemUseOnEntityFunc func(ctx *player.Context, e world.Entity)
type PlayerHandleItemReleaseFunc func(ctx *player.Context, item item.Stack, dur time.Duration)
type PlayerHandleItemConsumeFunc func(ctx *player.Context, item item.Stack)
type PlayerHandleAttackEntityFunc func(ctx *player.Context, e world.Entity, force *float64, height *float64, critical *bool)
type PlayerHandleExperienceGainFunc func(ctx *player.Context, amount *int)
type PlayerHandlePunchAirFunc func(ctx *player.Context)
type PlayerHandleSignEditFunc func(ctx *player.Context, pos cube.Pos, frontSide bool, oldText string, newText string)
type PlayerHandleLecternPageTurnFunc func(ctx *player.Context, pos cube.Pos, oldPage int, newPage *int)
type PlayerHandleItemDamageFunc func(ctx *player.Context, i item.Stack, damage int)
type PlayerHandleItemPickupFunc func(ctx *player.Context, i *item.Stack)
type PlayerHandleHeldSlotChangeFunc func(ctx *player.Context, from int, to int)
type PlayerHandleItemDropFunc func(ctx *player.Context, s item.Stack)
type PlayerHandleTransferFunc func(ctx *player.Context, addr *net.UDPAddr)
type PlayerHandleCommandExecutionFunc func(ctx *player.Context, command cmd.Command, args []string)
type PlayerHandleQuitFunc func(p *player.Player)
type PlayerHandleDiagnosticsFunc func(p *player.Player, d session.Diagnostics)
