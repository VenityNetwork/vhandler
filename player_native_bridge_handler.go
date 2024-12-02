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

type PlayerNativeBridgeHandler struct {
	p *player.Player
	h *PlayerHandlers
}

// Compile-time check to ensure that the PlayerNativeBridgeHandler implements the PlayerHandler interface.
var _ player.Handler = &PlayerNativeBridgeHandler{}

func (h *PlayerNativeBridgeHandler) HandleMove(ctx *player.Context, newPos mgl64.Vec3, newRot cube.Rotation) {
	for _, handler := range h.h.moveHandlers {
		handler.h(ctx, newPos, newRot)
	}
}

func (h *PlayerNativeBridgeHandler) HandleJump(p *player.Player) {
	for _, handler := range h.h.jumpHandlers {
		handler.h(p)
	}
}

func (h *PlayerNativeBridgeHandler) HandleTeleport(ctx *player.Context, pos mgl64.Vec3) {
	for _, handler := range h.h.teleportHandlers {
		handler.h(ctx, pos)
	}
}

func (h *PlayerNativeBridgeHandler) HandleChangeWorld(before *world.World, after *world.World) {
	for _, handler := range h.h.changeWorldHandlers {
		handler.h(before, after)
	}
}

func (h *PlayerNativeBridgeHandler) HandleToggleSprint(ctx *player.Context, after bool) {
	for _, handler := range h.h.toggleSprintHandlers {
		handler.h(ctx, after)
	}
}

func (h *PlayerNativeBridgeHandler) HandleToggleSneak(ctx *player.Context, after bool) {
	for _, handler := range h.h.toggleSneakHandlers {
		handler.h(ctx, after)
	}
}

func (h *PlayerNativeBridgeHandler) HandleChat(ctx *player.Context, message *string) {
	for _, handler := range h.h.chatHandlers {
		handler.h(ctx, message)
	}
}

func (h *PlayerNativeBridgeHandler) HandleFoodLoss(ctx *player.Context, from int, to *int) {
	for _, handler := range h.h.foodLossHandlers {
		handler.h(ctx, from, to)
	}
}

func (h *PlayerNativeBridgeHandler) HandleHeal(ctx *player.Context, health *float64, src world.HealingSource) {
	for _, handler := range h.h.healHandlers {
		handler.h(ctx, health, src)
	}
}

func (h *PlayerNativeBridgeHandler) HandleHurt(ctx *player.Context, damage *float64, immune bool, attackImmunity *time.Duration, src world.DamageSource) {
	for _, handler := range h.h.hurtHandlers {
		handler.h(ctx, damage, immune, attackImmunity, src)
	}
}

func (h *PlayerNativeBridgeHandler) HandleDeath(p *player.Player, src world.DamageSource, keepInv *bool) {
	for _, handler := range h.h.deathHandlers {
		handler.h(p, src, keepInv)
	}
}

func (h *PlayerNativeBridgeHandler) HandleRespawn(p *player.Player, pos *mgl64.Vec3, w **world.World) {
	for _, handler := range h.h.respawnHandlers {
		handler.h(p, pos, w)
	}
}

func (h *PlayerNativeBridgeHandler) HandleSkinChange(ctx *player.Context, skin *skin.Skin) {
	for _, handler := range h.h.skinChangeHandlers {
		handler.h(ctx, skin)
	}
}

func (h *PlayerNativeBridgeHandler) HandleFireExtinguish(ctx *player.Context, pos cube.Pos) {
	for _, handler := range h.h.fireExtinguishHandlers {
		handler.h(ctx, pos)
	}
}

func (h *PlayerNativeBridgeHandler) HandleStartBreak(ctx *player.Context, pos cube.Pos) {
	for _, handler := range h.h.startBreakHandlers {
		handler.h(ctx, pos)
	}
}

func (h *PlayerNativeBridgeHandler) HandleBlockBreak(ctx *player.Context, pos cube.Pos, drops *[]item.Stack, xp *int) {
	for _, handler := range h.h.blockBreakHandlers {
		handler.h(ctx, pos, drops, xp)
	}
}

func (h *PlayerNativeBridgeHandler) HandleBlockPlace(ctx *player.Context, pos cube.Pos, b world.Block) {
	for _, handler := range h.h.blockPlaceHandlers {
		handler.h(ctx, pos, b)
	}
}

func (h *PlayerNativeBridgeHandler) HandleBlockPick(ctx *player.Context, pos cube.Pos, b world.Block) {
	for _, handler := range h.h.blockPickHandlers {
		handler.h(ctx, pos, b)
	}
}

func (h *PlayerNativeBridgeHandler) HandleItemUse(ctx *player.Context) {
	for _, handler := range h.h.itemUseHandlers {
		handler.h(ctx)
	}
}

func (h *PlayerNativeBridgeHandler) HandleItemUseOnBlock(ctx *player.Context, pos cube.Pos, face cube.Face, clickPos mgl64.Vec3) {
	for _, handler := range h.h.itemUseOnBlockHandlers {
		handler.h(ctx, pos, face, clickPos)
	}
}

func (h *PlayerNativeBridgeHandler) HandleItemUseOnEntity(ctx *player.Context, e world.Entity) {
	for _, handler := range h.h.itemUseOnEntityHandlers {
		handler.h(ctx, e)
	}
}

func (h *PlayerNativeBridgeHandler) HandleItemRelease(ctx *player.Context, item item.Stack, dur time.Duration) {
	for _, handler := range h.h.itemReleaseHandlers {
		handler.h(ctx, item, dur)
	}
}

func (h *PlayerNativeBridgeHandler) HandleItemConsume(ctx *player.Context, item item.Stack) {
	for _, handler := range h.h.itemConsumeHandlers {
		handler.h(ctx, item)
	}
}

func (h *PlayerNativeBridgeHandler) HandleAttackEntity(ctx *player.Context, e world.Entity, force *float64, height *float64, critical *bool) {
	for _, handler := range h.h.attackEntityHandlers {
		handler.h(ctx, e, force, height, critical)
	}
}

func (h *PlayerNativeBridgeHandler) HandleExperienceGain(ctx *player.Context, amount *int) {
	for _, handler := range h.h.experienceGainHandlers {
		handler.h(ctx, amount)
	}
}

func (h *PlayerNativeBridgeHandler) HandlePunchAir(ctx *player.Context) {
	for _, handler := range h.h.punchAirHandlers {
		handler.h(ctx)
	}
}

func (h *PlayerNativeBridgeHandler) HandleSignEdit(ctx *player.Context, pos cube.Pos, frontSide bool, oldText string, newText string) {
	for _, handler := range h.h.signEditHandlers {
		handler.h(ctx, pos, frontSide, oldText, newText)
	}
}

func (h *PlayerNativeBridgeHandler) HandleLecternPageTurn(ctx *player.Context, pos cube.Pos, oldPage int, newPage *int) {
	for _, handler := range h.h.lecternPageTurnHandlers {
		handler.h(ctx, pos, oldPage, newPage)
	}
}

func (h *PlayerNativeBridgeHandler) HandleItemDamage(ctx *player.Context, i item.Stack, damage int) {
	for _, handler := range h.h.itemDamageHandlers {
		handler.h(ctx, i, damage)
	}
}

func (h *PlayerNativeBridgeHandler) HandleItemPickup(ctx *player.Context, i *item.Stack) {
	for _, handler := range h.h.itemPickupHandlers {
		handler.h(ctx, i)
	}
}

func (h *PlayerNativeBridgeHandler) HandleHeldSlotChange(ctx *player.Context, from int, to int) {
	for _, handler := range h.h.heldSlotChangeHandlers {
		handler.h(ctx, from, to)
	}
}

func (h *PlayerNativeBridgeHandler) HandleItemDrop(ctx *player.Context, s item.Stack) {
	for _, handler := range h.h.itemDropHandlers {
		handler.h(ctx, s)
	}
}

func (h *PlayerNativeBridgeHandler) HandleTransfer(ctx *player.Context, addr *net.UDPAddr) {
	for _, handler := range h.h.transferHandlers {
		handler.h(ctx, addr)
	}
}

func (h *PlayerNativeBridgeHandler) HandleCommandExecution(ctx *player.Context, command cmd.Command, args []string) {
	for _, handler := range h.h.commandExecutionHandlers {
		handler.h(ctx, command, args)
	}
}

func (h *PlayerNativeBridgeHandler) HandleQuit(p *player.Player) {
	for _, handler := range h.h.quitHandlers {
		handler.h(p)
	}
}

func (h *PlayerNativeBridgeHandler) HandleDiagnostics(p *player.Player, d session.Diagnostics) {
	for _, handler := range h.h.diagnosticsHandlers {
		handler.h(p, d)
	}
}
