package vhandler

import (
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

func (w *World) OnLiquidFlow(p priority.Priority, h WorldLiquidFlowHandler) {
	w.handlers[WorldLiquidFlowId].add(p, h)
}

func (w *World) OnLiquidDecay(p priority.Priority, h WorldLiquidDecayHandler) {
	w.handlers[WorldLiquidDecayId].add(p, h)
}

func (w *World) OnLiquidHarden(p priority.Priority, h WorldLiquidHardenHandler) {
	w.handlers[WorldLiquidHardenId].add(p, h)
}

func (w *World) OnSound(p priority.Priority, h WorldSoundHandler) {
	w.handlers[WorldSoundId].add(p, h)
}

func (w *World) OnFireSpread(p priority.Priority, h WorldFireSpreadHandler) {
	w.handlers[WorldFireSpreadId].add(p, h)
}

func (w *World) OnBlockBurn(p priority.Priority, h WorldBlockBurnHandler) {
	w.handlers[WorldBlockBurnId].add(p, h)
}

func (w *World) OnEntitySpawn(p priority.Priority, h WorldEntitySpawnHandler) {
	w.handlers[WorldEntitySpawnId].add(p, h)
}

func (w *World) OnEntityDespawn(p priority.Priority, h WorldEntityDespawnHandler) {
	w.handlers[WorldEntityDespawnId].add(p, h)
}

func (w *World) OnClose(p priority.Priority, h WorldCloseHandler) {
	w.handlers[WorldCloseId].add(p, h)
}

func (w *World) Set(wo *world.World) {
	wo.Handle(w.h)
}

func (w *World) Attach(p priority.Priority, wh world.Handler) {
	nop := world.NopHandler{}
	nopHandlers := w.getHandlers(nop)

	handlers := w.getHandlers(wh)
	for hId, handler := range handlers {
		if handler == nopHandlers[hId] {
			continue // ignore nop handler
		}
		w.handlers[hId].add(p, handler)
	}
}

func (w *World) Detach(wh world.Handler) error {
	handlers := w.getHandlers(wh)
	for hId, handler := range handlers {
		if err := w.handlers[hId].remove(handler); err != nil {
			return err
		}
	}
	return nil
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
