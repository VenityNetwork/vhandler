package vhandler

import (
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/entity"
	"github.com/df-mc/dragonfly/server/event"
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/player/skin"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/go-gl/mathgl/mgl64"
	"net"
	"time"
)

type Handler interface {
}

type MainHandler struct {
	player.Handler

	vh *VHandler
}

func newMainHandler(vh *VHandler) *MainHandler {
	return &MainHandler{vh: vh}
}

type ItemDropHandler func(ctx *event.Context, e *entity.Item)

func (m *MainHandler) HandleItemDrop(ctx *event.Context, e *entity.Item) {
	m.vh.handlers[ItemDropId].handle(func(h Handler) {
		h.(ItemDropHandler)(ctx, e)
	})
}

type MoveHandler func(ctx *event.Context, newPos mgl64.Vec3, newYaw, newPitch float64)

func (m *MainHandler) HandleMove(ctx *event.Context, newPos mgl64.Vec3, newYaw, newPitch float64) {
	m.vh.handlers[MoveId].handle(func(h Handler) {
		h.(MoveHandler)(ctx, newPos, newYaw, newPitch)
	})
}

type JumpHandler func()

func (m *MainHandler) HandleJump() {
	m.vh.handlers[JumpId].handle(func(h Handler) {
		h.(JumpHandler)()
	})
}

type TeleportHandler func(ctx *event.Context, pos mgl64.Vec3)

func (m *MainHandler) HandleTeleport(ctx *event.Context, pos mgl64.Vec3) {
	m.vh.handlers[TeleportId].handle(func(h Handler) {
		h.(TeleportHandler)(ctx, pos)
	})
}

type ChangeWorldHandler func(before, after *world.World)

func (m *MainHandler) HandleChangeWorld(before, after *world.World) {
	m.vh.handlers[ChangeWorldId].handle(func(h Handler) {
		h.(ChangeWorldHandler)(before, after)
	})
}

type ToggleSprintHandler func(ctx *event.Context, after bool)

func (m *MainHandler) HandleToggleSprint(ctx *event.Context, after bool) {
	m.vh.handlers[ToggleSprintId].handle(func(h Handler) {
		h.(ToggleSprintHandler)(ctx, after)
	})
}

type ToggleSneakHandler func(ctx *event.Context, after bool)

func (m *MainHandler) HandleToggleSneak(ctx *event.Context, after bool) {
	m.vh.handlers[ToggleSneakId].handle(func(h Handler) {
		h.(ToggleSneakHandler)(ctx, after)
	})
}

type CommandExecutionHandler func(ctx *event.Context, command cmd.Command, args []string)

func (m *MainHandler) HandleCommandExecution(ctx *event.Context, command cmd.Command, args []string) {
	m.vh.handlers[CommandExecutionId].handle(func(h Handler) {
		h.(CommandExecutionHandler)(ctx, command, args)
	})
}

type TransferHandler func(ctx *event.Context, addr *net.UDPAddr)

func (m *MainHandler) HandleTransfer(ctx *event.Context, addr *net.UDPAddr) {
	m.vh.handlers[TransferId].handle(func(h Handler) {
		h.(TransferHandler)(ctx, addr)
	})
}

type ChatHandler func(ctx *event.Context, message *string)

func (m *MainHandler) HandleChat(ctx *event.Context, message *string) {
	m.vh.handlers[ChatId].handle(func(h Handler) {
		h.(ChatHandler)(ctx, message)
	})
}

type SkinChangeHandler func(ctx *event.Context, skin *skin.Skin)

func (m *MainHandler) HandleSkinChange(ctx *event.Context, skin *skin.Skin) {
	m.vh.handlers[SkinChangeId].handle(func(h Handler) {
		h.(SkinChangeHandler)(ctx, skin)
	})
}

type StartBreakHandler func(ctx *event.Context, pos cube.Pos)

func (m *MainHandler) HandleStartBreak(ctx *event.Context, pos cube.Pos) {
	m.vh.handlers[StartBreakId].handle(func(h Handler) {
		h.(StartBreakHandler)(ctx, pos)
	})
}

type BlockBreakHandler func(ctx *event.Context, pos cube.Pos, drops *[]item.Stack)

func (m *MainHandler) HandleBlockBreak(ctx *event.Context, pos cube.Pos, drops *[]item.Stack) {
	m.vh.handlers[BlockBreakId].handle(func(h Handler) {
		h.(BlockBreakHandler)(ctx, pos, drops)
	})
}

type BlockPlaceHandler func(ctx *event.Context, pos cube.Pos, b world.Block)

func (m *MainHandler) HandleBlockPlace(ctx *event.Context, pos cube.Pos, b world.Block) {
	m.vh.handlers[BlockPlaceId].handle(func(h Handler) {
		h.(BlockPlaceHandler)(ctx, pos, b)
	})
}

type BlockPickHandler func(ctx *event.Context, pos cube.Pos, b world.Block)

func (m *MainHandler) HandleBlockPick(ctx *event.Context, pos cube.Pos, b world.Block) {
	m.vh.handlers[BlockPickId].handle(func(h Handler) {
		h.(BlockPickHandler)(ctx, pos, b)
	})
}

type SignEditHandler func(ctx *event.Context, oldText, newText string)

func (m *MainHandler) HandleSignEdit(ctx *event.Context, oldText, newText string) {
	m.vh.handlers[SignEditId].handle(func(h Handler) {
		h.(SignEditHandler)(ctx, oldText, newText)
	})
}

type ItemPickupHandler func(ctx *event.Context, i item.Stack)

func (m *MainHandler) HandleItemPickup(ctx *event.Context, i item.Stack) {
	m.vh.handlers[ItemPickupId].handle(func(h Handler) {
		h.(ItemPickupHandler)(ctx, i)
	})
}

type ItemUseHandler func(ctx *event.Context)

func (m *MainHandler) HandleItemUse(ctx *event.Context) {
	m.vh.handlers[ItemUseId].handle(func(h Handler) {
		h.(ItemUseHandler)(ctx)
	})
}

type ItemUseOnBlockHandler func(ctx *event.Context, pos cube.Pos, face cube.Face, clickPos mgl64.Vec3)

func (m *MainHandler) HandleItemUseOnBlock(ctx *event.Context, pos cube.Pos, face cube.Face, clickPos mgl64.Vec3) {
	m.vh.handlers[ItemUseOnBlockId].handle(func(h Handler) {
		h.(ItemUseOnBlockHandler)(ctx, pos, face, clickPos)
	})
}

type ItemUseOnEntityHandler func(ctx *event.Context, e world.Entity)

func (m *MainHandler) HandleItemUseOnEntity(ctx *event.Context, e world.Entity) {
	m.vh.handlers[ItemUseOnEntityId].handle(func(h Handler) {
		h.(ItemUseOnEntityHandler)(ctx, e)
	})
}

type ItemConsumeHandler func(ctx *event.Context, item item.Stack)

func (m *MainHandler) HandleItemConsume(ctx *event.Context, item item.Stack) {
	m.vh.handlers[ItemConsumeId].handle(func(h Handler) {
		h.(ItemConsumeHandler)(ctx, item)
	})
}

type ItemDamageHandler func(ctx *event.Context, i item.Stack, damage int)

func (m *MainHandler) HandleItemDamage(ctx *event.Context, i item.Stack, damage int) {
	m.vh.handlers[ItemDamageId].handle(func(h Handler) {
		h.(ItemDamageHandler)(ctx, i, damage)
	})
}

type AttackEntityHandler func(ctx *event.Context, e world.Entity, force, height *float64, critical *bool)

func (m *MainHandler) HandleAttackEntity(ctx *event.Context, e world.Entity, force, height *float64, critical *bool) {
	m.vh.handlers[AttackEntityId].handle(func(h Handler) {
		h.(AttackEntityHandler)(ctx, e, force, height, critical)
	})
}

type ExperienceGainHandler func(ctx *event.Context, amount *int)

func (m *MainHandler) HandleExperienceGain(ctx *event.Context, amount *int) {
	m.vh.handlers[ExperienceGainId].handle(func(h Handler) {
		h.(ExperienceGainHandler)(ctx, amount)
	})
}

type PunchAirHandler func(ctx *event.Context)

func (m *MainHandler) HandlePunchAir(ctx *event.Context) {
	m.vh.handlers[PunchAirId].handle(func(h Handler) {
		h.(PunchAirHandler)(ctx)
	})
}

type HurtHandler func(ctx *event.Context, damage *float64, attackImmunity *time.Duration, src world.DamageSource)

func (m *MainHandler) HandleHurt(ctx *event.Context, damage *float64, attackImmunity *time.Duration, src world.DamageSource) {
	m.vh.handlers[HurtId].handle(func(h Handler) {
		h.(HurtHandler)(ctx, damage, attackImmunity, src)
	})
}

type HealHandler func(ctx *event.Context, health *float64, src world.HealingSource)

func (m *MainHandler) HandleHeal(ctx *event.Context, health *float64, src world.HealingSource) {
	m.vh.handlers[HealId].handle(func(h Handler) {
		h.(HealHandler)(ctx, health, src)
	})
}

type FoodLossHandler func(ctx *event.Context, from, to int)

func (m *MainHandler) HandleFoodLoss(ctx *event.Context, from, to int) {
	m.vh.handlers[FoodLossId].handle(func(h Handler) {
		h.(FoodLossHandler)(ctx, from, to)
	})
}

type DeathHandler func(src world.DamageSource, keepInv *bool)

func (m *MainHandler) HandleDeath(src world.DamageSource, keepInv *bool) {
	m.vh.handlers[DeathId].handle(func(h Handler) {
		h.(DeathHandler)(src, keepInv)
	})
}

type RespawnHandler func(pos *mgl64.Vec3, w **world.World)

func (m *MainHandler) HandleRespawn(pos *mgl64.Vec3, w **world.World) {
	m.vh.handlers[RespawnId].handle(func(h Handler) {
		h.(RespawnHandler)(pos, w)
	})
}

type QuitHandler func()

func (m *MainHandler) HandleQuit() {
	m.vh.handlers[QuitId].handle(func(h Handler) {
		h.(QuitHandler)()
	})
}
