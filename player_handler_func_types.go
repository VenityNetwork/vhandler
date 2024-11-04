package vhandler

import (
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/event"
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/player/diagnostics"
	"github.com/df-mc/dragonfly/server/player/skin"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/go-gl/mathgl/mgl64"
	"net"
	"time"
)

type PlayerHandleMoveFunc func(p *player.Player, ctx *event.Context, newPos mgl64.Vec3, newYaw float64, newPitch float64)
type PlayerHandleJumpFunc func(p *player.Player)
type PlayerHandleTeleportFunc func(p *player.Player, ctx *event.Context, pos mgl64.Vec3)
type PlayerHandleChangeWorldFunc func(p *player.Player, before *world.World, after *world.World)
type PlayerHandleToggleSprintFunc func(p *player.Player, ctx *event.Context, after bool)
type PlayerHandleToggleSneakFunc func(p *player.Player, ctx *event.Context, after bool)
type PlayerHandleChatFunc func(p *player.Player, ctx *event.Context, message *string)
type PlayerHandleFoodLossFunc func(p *player.Player, ctx *event.Context, from int, to *int)
type PlayerHandleHealFunc func(p *player.Player, ctx *event.Context, health *float64, src world.HealingSource)
type PlayerHandleHurtFunc func(p *player.Player, ctx *event.Context, damage *float64, attackImmunity *time.Duration, src world.DamageSource)
type PlayerHandleDeathFunc func(p *player.Player, src world.DamageSource, keepInv *bool)
type PlayerHandleRespawnFunc func(p *player.Player, pos *mgl64.Vec3, w **world.World)
type PlayerHandleSkinChangeFunc func(p *player.Player, ctx *event.Context, skin *skin.Skin)
type PlayerHandleFireExtinguishFunc func(p *player.Player, ctx *event.Context, pos cube.Pos)
type PlayerHandleStartBreakFunc func(p *player.Player, ctx *event.Context, pos cube.Pos)
type PlayerHandleBlockBreakFunc func(p *player.Player, ctx *event.Context, pos cube.Pos, drops *[]item.Stack, xp *int)
type PlayerHandleBlockPlaceFunc func(p *player.Player, ctx *event.Context, pos cube.Pos, b world.Block)
type PlayerHandleBlockPickFunc func(p *player.Player, ctx *event.Context, pos cube.Pos, b world.Block)
type PlayerHandleItemUseFunc func(p *player.Player, ctx *event.Context)
type PlayerHandleItemUseOnBlockFunc func(p *player.Player, ctx *event.Context, pos cube.Pos, face cube.Face, clickPos mgl64.Vec3)
type PlayerHandleItemUseOnEntityFunc func(p *player.Player, ctx *event.Context, e world.Entity)
type PlayerHandleItemConsumeFunc func(p *player.Player, ctx *event.Context, item item.Stack)
type PlayerHandleAttackEntityFunc func(p *player.Player, ctx *event.Context, e world.Entity, force *float64, height *float64, critical *bool)
type PlayerHandleExperienceGainFunc func(p *player.Player, ctx *event.Context, amount *int)
type PlayerHandlePunchAirFunc func(p *player.Player, ctx *event.Context)
type PlayerHandleSignEditFunc func(p *player.Player, ctx *event.Context, pos cube.Pos, frontSide bool, oldText string, newText string)
type PlayerHandleLecternPageTurnFunc func(p *player.Player, ctx *event.Context, pos cube.Pos, oldPage int, newPage *int)
type PlayerHandleItemDamageFunc func(p *player.Player, ctx *event.Context, i item.Stack, damage int)
type PlayerHandleItemPickupFunc func(p *player.Player, ctx *event.Context, i *item.Stack)
type PlayerHandleItemDropFunc func(p *player.Player, ctx *event.Context, e world.Entity)
type PlayerHandleTransferFunc func(p *player.Player, ctx *event.Context, addr *net.UDPAddr)
type PlayerHandleCommandExecutionFunc func(p *player.Player, ctx *event.Context, command cmd.Command, args []string)
type PlayerHandleQuitFunc func(p *player.Player)
type PlayerHandleDiagnosticsFunc func(p *player.Player, d diagnostics.Diagnostics)
