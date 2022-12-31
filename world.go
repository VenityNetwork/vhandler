package vhandler

import (
	"errors"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/venitynetwork/vhandler/priority"
)

type World struct {
	handlers map[handlerId]*subHandler
	h        *WorldHandler
}

func NewWorld() *World {
	v := &World{handlers: map[handlerId]*subHandler{}}
	v.h = newWorldHandler(v)

	v.handlers[WorldLiquidFlowId] = newSubHandler()
	v.handlers[WorldLiquidDecayId] = newSubHandler()
	v.handlers[WorldLiquidHardenId] = newSubHandler()
	v.handlers[WorldSoundId] = newSubHandler()
	v.handlers[WorldFireSpreadId] = newSubHandler()
	v.handlers[WorldBlockBurnId] = newSubHandler()
	v.handlers[WorldEntitySpawnId] = newSubHandler()
	v.handlers[WorldEntityDespawnId] = newSubHandler()
	v.handlers[WorldCloseId] = newSubHandler()

	return v
}

func (v *World) OnLiquidFlow(p priority.Priority, h WorldLiquidFlowHandler) {
	v.handlers[WorldLiquidFlowId].add(p, h)
}

func (v *World) OnLiquidDecay(p priority.Priority, h WorldLiquidDecayHandler) {
	v.handlers[WorldLiquidDecayId].add(p, h)
}

func (v *World) OnLiquidHarden(p priority.Priority, h WorldLiquidHardenHandler) {
	v.handlers[WorldLiquidHardenId].add(p, h)
}

func (v *World) OnSound(p priority.Priority, h WorldSoundHandler) {
	v.handlers[WorldSoundId].add(p, h)
}

func (v *World) OnFireSpread(p priority.Priority, h WorldFireSpreadHandler) {
	v.handlers[WorldFireSpreadId].add(p, h)
}

func (v *World) OnBlockBurn(p priority.Priority, h WorldBlockBurnHandler) {
	v.handlers[WorldBlockBurnId].add(p, h)
}

func (v *World) OnEntitySpawn(p priority.Priority, h WorldEntitySpawnHandler) {
	v.handlers[WorldEntitySpawnId].add(p, h)
}

func (v *World) OnEntityDespawn(p priority.Priority, h WorldEntityDespawnHandler) {
	v.handlers[WorldEntityDespawnId].add(p, h)
}

func (v *World) OnClose(p priority.Priority, h WorldCloseHandler) {
	v.handlers[WorldCloseId].add(p, h)
}

func (v *World) Set(wo *world.World) {
	wo.Handle(v.h)
}

func (v *World) Attach(p priority.Priority, wh world.Handler) {
	nop := world.NopHandler{}
	nopHandlers := v.getHandlers(nop)

	handlers := v.getHandlers(wh)
	for hId, handler := range handlers {
		if handler == nopHandlers[hId] {
			continue // ignore nop handler
		}
		v.handlers[hId].add(p, handler)
	}
}

func (v *World) Detach(wh world.Handler) error {
	handlers := v.getHandlers(wh)
	for hId, handler := range handlers {
		if err := v.handlers[hId].remove(handler); err != nil {
			return err
		}
	}
	return nil
}

func (v *World) Remove(h Handler) error {
	hId, ok := v.getHandlerId(h)
	if !ok {
		return errors.New("invalid handler")
	}
	return v.handlers[hId].remove(h)
}

func (*World) getHandlers(h world.Handler) map[handlerId]Handler {
	var handlers map[handlerId]Handler

	handlers[WorldLiquidFlowId] = h.HandleLiquidFlow
	handlers[WorldLiquidDecayId] = h.HandleLiquidDecay
	handlers[WorldLiquidHardenId] = h.HandleLiquidHarden
	handlers[WorldSoundId] = h.HandleSound
	handlers[WorldFireSpreadId] = h.HandleFireSpread
	handlers[WorldBlockBurnId] = h.HandleBlockBurn
	handlers[WorldEntitySpawnId] = h.HandleEntitySpawn
	handlers[WorldEntityDespawnId] = h.HandleEntityDespawn
	handlers[WorldCloseId] = h.HandleClose

	return handlers
}

func (*World) getHandlerId(h Handler) (handlerId, bool) {
	switch h.(type) {
	case WorldLiquidFlowHandler:
		return WorldLiquidFlowId, true
	case WorldLiquidDecayHandler:
		return WorldLiquidDecayId, true
	case WorldLiquidHardenHandler:
		return WorldLiquidHardenId, true
	case WorldSoundHandler:
		return WorldSoundId, true
	case WorldFireSpreadHandler:
		return WorldFireSpreadId, true
	case WorldBlockBurnHandler:
		return WorldBlockBurnId, true
	case WorldEntitySpawnHandler:
		return WorldEntitySpawnId, true
	case WorldEntityDespawnHandler:
		return WorldEntityDespawnId, true
	case WorldCloseHandler:
		return WorldCloseId, true
	default:
		return 0, false
	}
}
