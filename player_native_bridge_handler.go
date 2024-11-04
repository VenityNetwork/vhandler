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

type PlayerNativeBridgeHandler struct {
	p *player.Player
	h *PlayerHandlers
}

// Compile-time check to ensure that the PlayerNativeBridgeHandler implements the PlayerHandler interface.
var _ player.Handler = &PlayerNativeBridgeHandler{}

func (h *PlayerNativeBridgeHandler) HandleMove(ctx *event.Context, newPos mgl64.Vec3, newYaw float64, newPitch float64) {
	for _, handler := range h.h.moveHandlers {
		handler.h(h.p, ctx, newPos, newYaw, newPitch)
	}
}

func (h *PlayerNativeBridgeHandler) HandleJump() {
	for _, handler := range h.h.jumpHandlers {
		handler.h(h.p)
	}
}

func (h *PlayerNativeBridgeHandler) HandleTeleport(ctx *event.Context, pos mgl64.Vec3) {
	for _, handler := range h.h.teleportHandlers {
		handler.h(h.p, ctx, pos)
	}
}

func (h *PlayerNativeBridgeHandler) HandleChangeWorld(before *world.World, after *world.World) {
	for _, handler := range h.h.changeWorldHandlers {
		handler.h(h.p, before, after)
	}
}

func (h *PlayerNativeBridgeHandler) HandleToggleSprint(ctx *event.Context, after bool) {
	for _, handler := range h.h.toggleSprintHandlers {
		handler.h(h.p, ctx, after)
	}
}

func (h *PlayerNativeBridgeHandler) HandleToggleSneak(ctx *event.Context, after bool) {
	for _, handler := range h.h.toggleSneakHandlers {
		handler.h(h.p, ctx, after)
	}
}

func (h *PlayerNativeBridgeHandler) HandleChat(ctx *event.Context, message *string) {
	for _, handler := range h.h.chatHandlers {
		handler.h(h.p, ctx, message)
	}
}

func (h *PlayerNativeBridgeHandler) HandleFoodLoss(ctx *event.Context, from int, to *int) {
	for _, handler := range h.h.foodLossHandlers {
		handler.h(h.p, ctx, from, to)
	}
}

func (h *PlayerNativeBridgeHandler) HandleHeal(ctx *event.Context, health *float64, src world.HealingSource) {
	for _, handler := range h.h.healHandlers {
		handler.h(h.p, ctx, health, src)
	}
}

func (h *PlayerNativeBridgeHandler) HandleHurt(ctx *event.Context, damage *float64, attackImmunity *time.Duration, src world.DamageSource) {
	for _, handler := range h.h.hurtHandlers {
		handler.h(h.p, ctx, damage, attackImmunity, src)
	}
}

func (h *PlayerNativeBridgeHandler) HandleDeath(src world.DamageSource, keepInv *bool) {
	for _, handler := range h.h.deathHandlers {
		handler.h(h.p, src, keepInv)
	}
}

func (h *PlayerNativeBridgeHandler) HandleRespawn(pos *mgl64.Vec3, w **world.World) {
	for _, handler := range h.h.respawnHandlers {
		handler.h(h.p, pos, w)
	}
}

func (h *PlayerNativeBridgeHandler) HandleSkinChange(ctx *event.Context, skin *skin.Skin) {
	for _, handler := range h.h.skinChangeHandlers {
		handler.h(h.p, ctx, skin)
	}
}

func (h *PlayerNativeBridgeHandler) HandleFireExtinguish(ctx *event.Context, pos cube.Pos) {
	for _, handler := range h.h.fireExtinguishHandlers {
		handler.h(h.p, ctx, pos)
	}
}

func (h *PlayerNativeBridgeHandler) HandleStartBreak(ctx *event.Context, pos cube.Pos) {
	for _, handler := range h.h.startBreakHandlers {
		handler.h(h.p, ctx, pos)
	}
}

func (h *PlayerNativeBridgeHandler) HandleBlockBreak(ctx *event.Context, pos cube.Pos, drops *[]item.Stack, xp *int) {
	for _, handler := range h.h.blockBreakHandlers {
		handler.h(h.p, ctx, pos, drops, xp)
	}
}

func (h *PlayerNativeBridgeHandler) HandleBlockPlace(ctx *event.Context, pos cube.Pos, b world.Block) {
	for _, handler := range h.h.blockPlaceHandlers {
		handler.h(h.p, ctx, pos, b)
	}
}

func (h *PlayerNativeBridgeHandler) HandleBlockPick(ctx *event.Context, pos cube.Pos, b world.Block) {
	for _, handler := range h.h.blockPickHandlers {
		handler.h(h.p, ctx, pos, b)
	}
}

func (h *PlayerNativeBridgeHandler) HandleItemUse(ctx *event.Context) {
	for _, handler := range h.h.itemUseHandlers {
		handler.h(h.p, ctx)
	}
}

func (h *PlayerNativeBridgeHandler) HandleItemUseOnBlock(ctx *event.Context, pos cube.Pos, face cube.Face, clickPos mgl64.Vec3) {
	for _, handler := range h.h.itemUseOnBlockHandlers {
		handler.h(h.p, ctx, pos, face, clickPos)
	}
}

func (h *PlayerNativeBridgeHandler) HandleItemUseOnEntity(ctx *event.Context, e world.Entity) {
	for _, handler := range h.h.itemUseOnEntityHandlers {
		handler.h(h.p, ctx, e)
	}
}

func (h *PlayerNativeBridgeHandler) HandleItemConsume(ctx *event.Context, item item.Stack) {
	for _, handler := range h.h.itemConsumeHandlers {
		handler.h(h.p, ctx, item)
	}
}

func (h *PlayerNativeBridgeHandler) HandleAttackEntity(ctx *event.Context, e world.Entity, force *float64, height *float64, critical *bool) {
	for _, handler := range h.h.attackEntityHandlers {
		handler.h(h.p, ctx, e, force, height, critical)
	}
}

func (h *PlayerNativeBridgeHandler) HandleExperienceGain(ctx *event.Context, amount *int) {
	for _, handler := range h.h.experienceGainHandlers {
		handler.h(h.p, ctx, amount)
	}
}

func (h *PlayerNativeBridgeHandler) HandlePunchAir(ctx *event.Context) {
	for _, handler := range h.h.punchAirHandlers {
		handler.h(h.p, ctx)
	}
}

func (h *PlayerNativeBridgeHandler) HandleSignEdit(ctx *event.Context, pos cube.Pos, frontSide bool, oldText string, newText string) {
	for _, handler := range h.h.signEditHandlers {
		handler.h(h.p, ctx, pos, frontSide, oldText, newText)
	}
}

func (h *PlayerNativeBridgeHandler) HandleLecternPageTurn(ctx *event.Context, pos cube.Pos, oldPage int, newPage *int) {
	for _, handler := range h.h.lecternPageTurnHandlers {
		handler.h(h.p, ctx, pos, oldPage, newPage)
	}
}

func (h *PlayerNativeBridgeHandler) HandleItemDamage(ctx *event.Context, i item.Stack, damage int) {
	for _, handler := range h.h.itemDamageHandlers {
		handler.h(h.p, ctx, i, damage)
	}
}

func (h *PlayerNativeBridgeHandler) HandleItemPickup(ctx *event.Context, i *item.Stack) {
	for _, handler := range h.h.itemPickupHandlers {
		handler.h(h.p, ctx, i)
	}
}

func (h *PlayerNativeBridgeHandler) HandleItemDrop(ctx *event.Context, e world.Entity) {
	for _, handler := range h.h.itemDropHandlers {
		handler.h(h.p, ctx, e)
	}
}

func (h *PlayerNativeBridgeHandler) HandleTransfer(ctx *event.Context, addr *net.UDPAddr) {
	for _, handler := range h.h.transferHandlers {
		handler.h(h.p, ctx, addr)
	}
}

func (h *PlayerNativeBridgeHandler) HandleCommandExecution(ctx *event.Context, command cmd.Command, args []string) {
	for _, handler := range h.h.commandExecutionHandlers {
		handler.h(h.p, ctx, command, args)
	}
}

func (h *PlayerNativeBridgeHandler) HandleQuit() {
	for _, handler := range h.h.quitHandlers {
		handler.h(h.p)
	}
}

func (h *PlayerNativeBridgeHandler) HandleDiagnostics(d diagnostics.Diagnostics) {
	for _, handler := range h.h.diagnosticsHandlers {
		handler.h(h.p, d)
	}
}
