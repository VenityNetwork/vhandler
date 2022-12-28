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

func (w *World) Remove(h Handler) error {
	switch h.(type) {
	case WorldLiquidFlowHandler:
		return w.handlers[WorldLiquidFlowId].remove(h)
	case WorldLiquidDecayHandler:
		return w.handlers[WorldLiquidDecayId].remove(h)
	case WorldLiquidHardenHandler:
		return w.handlers[WorldLiquidHardenId].remove(h)
	case WorldSoundHandler:
		return w.handlers[WorldSoundId].remove(h)
	case WorldFireSpreadHandler:
		return w.handlers[WorldFireSpreadId].remove(h)
	case WorldBlockBurnHandler:
		return w.handlers[WorldBlockBurnId].remove(h)
	case WorldEntitySpawnHandler:
		return w.handlers[WorldEntitySpawnId].remove(h)
	case WorldEntityDespawnHandler:
		return w.handlers[WorldEntityDespawnId].remove(h)
	case WorldCloseHandler:
		return w.handlers[WorldCloseId].remove(h)
	default:
		return errors.New("invalid handler")
	}
}
