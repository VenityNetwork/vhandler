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

type PlayerHandler struct {
	player.Handler

	ph *Player
}

func newPlayerHandler(ph *Player) *PlayerHandler {
	return &PlayerHandler{ph: ph}
}

type PlayerItemDropHandler func(ctx *event.Context, e *entity.Item)

func (p *PlayerHandler) HandleItemDrop(ctx *event.Context, e *entity.Item) {
	p.ph.handlers[PlayerItemDropId].handle(func(h Handler) {
		h.(PlayerItemDropHandler)(ctx, e)
	})
}

type MoveHandler func(ctx *event.Context, newPos mgl64.Vec3, newYaw, newPitch float64)

func (p *PlayerHandler) HandleMove(ctx *event.Context, newPos mgl64.Vec3, newYaw, newPitch float64) {
	p.ph.handlers[PlayerMoveId].handle(func(h Handler) {
		h.(MoveHandler)(ctx, newPos, newYaw, newPitch)
	})
}

type PlayerJumpHandler func()

func (p *PlayerHandler) HandleJump() {
	p.ph.handlers[PlayerJumpId].handle(func(h Handler) {
		h.(PlayerJumpHandler)()
	})
}

type PlayerTeleportHandler func(ctx *event.Context, pos mgl64.Vec3)

func (p *PlayerHandler) HandleTeleport(ctx *event.Context, pos mgl64.Vec3) {
	p.ph.handlers[PlayerTeleportId].handle(func(h Handler) {
		h.(PlayerTeleportHandler)(ctx, pos)
	})
}

type PlayerChangeWorldHandler func(before, after *world.World)

func (p *PlayerHandler) HandleChangeWorld(before, after *world.World) {
	p.ph.handlers[PlayerChangeWorldId].handle(func(h Handler) {
		h.(PlayerChangeWorldHandler)(before, after)
	})
}

type PlayerToggleSprintHandler func(ctx *event.Context, after bool)

func (p *PlayerHandler) HandleToggleSprint(ctx *event.Context, after bool) {
	p.ph.handlers[PlayerToggleSprintId].handle(func(h Handler) {
		h.(PlayerToggleSprintHandler)(ctx, after)
	})
}

type PlayerToggleSneakHandler func(ctx *event.Context, after bool)

func (p *PlayerHandler) HandleToggleSneak(ctx *event.Context, after bool) {
	p.ph.handlers[PlayerToggleSneakId].handle(func(h Handler) {
		h.(PlayerToggleSneakHandler)(ctx, after)
	})
}

type PlayerCommandExecutionHandler func(ctx *event.Context, command cmd.Command, args []string)

func (p *PlayerHandler) HandleCommandExecution(ctx *event.Context, command cmd.Command, args []string) {
	p.ph.handlers[PlayerCommandExecutionId].handle(func(h Handler) {
		h.(PlayerCommandExecutionHandler)(ctx, command, args)
	})
}

type PlayerTransferHandler func(ctx *event.Context, addr *net.UDPAddr)

func (p *PlayerHandler) HandleTransfer(ctx *event.Context, addr *net.UDPAddr) {
	p.ph.handlers[PlayerTransferId].handle(func(h Handler) {
		h.(PlayerTransferHandler)(ctx, addr)
	})
}

type PlayerChatHandler func(ctx *event.Context, message *string)

func (p *PlayerHandler) HandleChat(ctx *event.Context, message *string) {
	p.ph.handlers[PlayerChatId].handle(func(h Handler) {
		h.(PlayerChatHandler)(ctx, message)
	})
}

type PlayerSkinChangeHandler func(ctx *event.Context, skin *skin.Skin)

func (p *PlayerHandler) HandleSkinChange(ctx *event.Context, skin *skin.Skin) {
	p.ph.handlers[PlayerSkinChangeId].handle(func(h Handler) {
		h.(PlayerSkinChangeHandler)(ctx, skin)
	})
}

type PlayerStartBreakHandler func(ctx *event.Context, pos cube.Pos)

func (p *PlayerHandler) HandleStartBreak(ctx *event.Context, pos cube.Pos) {
	p.ph.handlers[PlayerStartBreakId].handle(func(h Handler) {
		h.(PlayerStartBreakHandler)(ctx, pos)
	})
}

type PlayerBlockBreakHandler func(ctx *event.Context, pos cube.Pos, drops *[]item.Stack, xp *int)

func (p *PlayerHandler) HandleBlockBreak(ctx *event.Context, pos cube.Pos, drops *[]item.Stack, xp *int) {
	p.ph.handlers[PlayerBlockBreakId].handle(func(h Handler) {
		h.(PlayerBlockBreakHandler)(ctx, pos, drops, xp)
	})
}

type PlayerBlockPlaceHandler func(ctx *event.Context, pos cube.Pos, b world.Block)

func (p *PlayerHandler) HandleBlockPlace(ctx *event.Context, pos cube.Pos, b world.Block) {
	p.ph.handlers[PlayerBlockPlaceId].handle(func(h Handler) {
		h.(PlayerBlockPlaceHandler)(ctx, pos, b)
	})
}

type PlayerBlockPickHandler func(ctx *event.Context, pos cube.Pos, b world.Block)

func (p *PlayerHandler) HandleBlockPick(ctx *event.Context, pos cube.Pos, b world.Block) {
	p.ph.handlers[PlayerBlockPickId].handle(func(h Handler) {
		h.(PlayerBlockPickHandler)(ctx, pos, b)
	})
}

type PlayerSignEditHandler func(ctx *event.Context, oldText, newText string)

func (p *PlayerHandler) HandleSignEdit(ctx *event.Context, oldText, newText string) {
	p.ph.handlers[PlayerSignEditId].handle(func(h Handler) {
		h.(PlayerSignEditHandler)(ctx, oldText, newText)
	})
}

type PlayerItemPickupHandler func(ctx *event.Context, i item.Stack)

func (p *PlayerHandler) HandleItemPickup(ctx *event.Context, i item.Stack) {
	p.ph.handlers[PlayerItemPickupId].handle(func(h Handler) {
		h.(PlayerItemPickupHandler)(ctx, i)
	})
}

type PlayerItemUseHandler func(ctx *event.Context)

func (p *PlayerHandler) HandleItemUse(ctx *event.Context) {
	p.ph.handlers[PlayerItemUseId].handle(func(h Handler) {
		h.(PlayerItemUseHandler)(ctx)
	})
}

type PlayerItemUseOnBlockHandler func(ctx *event.Context, pos cube.Pos, face cube.Face, clickPos mgl64.Vec3)

func (p *PlayerHandler) HandleItemUseOnBlock(ctx *event.Context, pos cube.Pos, face cube.Face, clickPos mgl64.Vec3) {
	p.ph.handlers[PlayerItemUseOnBlockId].handle(func(h Handler) {
		h.(PlayerItemUseOnBlockHandler)(ctx, pos, face, clickPos)
	})
}

type PlayerItemUseOnEntityHandler func(ctx *event.Context, e world.Entity)

func (p *PlayerHandler) HandleItemUseOnEntity(ctx *event.Context, e world.Entity) {
	p.ph.handlers[PlayerItemUseOnEntityId].handle(func(h Handler) {
		h.(PlayerItemUseOnEntityHandler)(ctx, e)
	})
}

type PlayerItemConsumeHandler func(ctx *event.Context, item item.Stack)

func (p *PlayerHandler) HandleItemConsume(ctx *event.Context, item item.Stack) {
	p.ph.handlers[PlayerItemConsumeId].handle(func(h Handler) {
		h.(PlayerItemConsumeHandler)(ctx, item)
	})
}

type PlayerItemDamageHandler func(ctx *event.Context, i item.Stack, damage int)

func (p *PlayerHandler) HandleItemDamage(ctx *event.Context, i item.Stack, damage int) {
	p.ph.handlers[PlayerItemDamageId].handle(func(h Handler) {
		h.(PlayerItemDamageHandler)(ctx, i, damage)
	})
}

type PlayerAttackEntityHandler func(ctx *event.Context, e world.Entity, force, height *float64, critical *bool)

func (p *PlayerHandler) HandleAttackEntity(ctx *event.Context, e world.Entity, force, height *float64, critical *bool) {
	p.ph.handlers[PlayerAttackEntityId].handle(func(h Handler) {
		h.(PlayerAttackEntityHandler)(ctx, e, force, height, critical)
	})
}

type PlayerExperienceGainHandler func(ctx *event.Context, amount *int)

func (p *PlayerHandler) HandleExperienceGain(ctx *event.Context, amount *int) {
	p.ph.handlers[PlayerExperienceGainId].handle(func(h Handler) {
		h.(PlayerExperienceGainHandler)(ctx, amount)
	})
}

type PlayerPunchAirHandler func(ctx *event.Context)

func (p *PlayerHandler) HandlePunchAir(ctx *event.Context) {
	p.ph.handlers[PlayerPunchAirId].handle(func(h Handler) {
		h.(PlayerPunchAirHandler)(ctx)
	})
}

type PlayerHurtHandler func(ctx *event.Context, damage *float64, attackImmunity *time.Duration, src world.DamageSource)

func (p *PlayerHandler) HandleHurt(ctx *event.Context, damage *float64, attackImmunity *time.Duration, src world.DamageSource) {
	p.ph.handlers[PlayerHurtId].handle(func(h Handler) {
		h.(PlayerHurtHandler)(ctx, damage, attackImmunity, src)
	})
}

type PlayerHealHandler func(ctx *event.Context, health *float64, src world.HealingSource)

func (p *PlayerHandler) HandleHeal(ctx *event.Context, health *float64, src world.HealingSource) {
	p.ph.handlers[PlayerHealId].handle(func(h Handler) {
		h.(PlayerHealHandler)(ctx, health, src)
	})
}

type PlayerFoodLossHandler func(ctx *event.Context, from int, to *int)

func (p *PlayerHandler) HandleFoodLoss(ctx *event.Context, from int, to *int) {
	p.ph.handlers[PlayerFoodLossId].handle(func(h Handler) {
		h.(PlayerFoodLossHandler)(ctx, from, to)
	})
}

type PlayerDeathHandler func(src world.DamageSource, keepInv *bool)

func (p *PlayerHandler) HandleDeath(src world.DamageSource, keepInv *bool) {
	p.ph.handlers[PlayerDeathId].handle(func(h Handler) {
		h.(PlayerDeathHandler)(src, keepInv)
	})
}

type PlayerRespawnHandler func(pos *mgl64.Vec3, w **world.World)

func (p *PlayerHandler) HandleRespawn(pos *mgl64.Vec3, w **world.World) {
	p.ph.handlers[PlayerRespawnId].handle(func(h Handler) {
		h.(PlayerRespawnHandler)(pos, w)
	})
}

type PlayerQuitHandler func()

func (p *PlayerHandler) HandleQuit() {
	p.ph.handlers[PlayerQuitId].handle(func(h Handler) {
		h.(PlayerQuitHandler)()
	})
}
