package vhandler

type WorldHandlers struct {
	liquidFlowHandlers    []*handlerWrapper[WorldHandleLiquidFlowFunc]
	liquidDecayHandlers   []*handlerWrapper[WorldHandleLiquidDecayFunc]
	liquidHardenHandlers  []*handlerWrapper[WorldHandleLiquidHardenFunc]
	soundHandlers         []*handlerWrapper[WorldHandleSoundFunc]
	fireSpreadHandlers    []*handlerWrapper[WorldHandleFireSpreadFunc]
	blockBurnHandlers     []*handlerWrapper[WorldHandleBlockBurnFunc]
	cropTrampleHandlers   []*handlerWrapper[WorldHandleCropTrampleFunc]
	entitySpawnHandlers   []*handlerWrapper[WorldHandleEntitySpawnFunc]
	entityDespawnHandlers []*handlerWrapper[WorldHandleEntityDespawnFunc]
	closeHandlers         []*handlerWrapper[WorldHandleCloseFunc]
}

func NewWorldHandlers() *WorldHandlers {
	return &WorldHandlers{
		liquidFlowHandlers:    make([]*handlerWrapper[WorldHandleLiquidFlowFunc], 0),
		liquidDecayHandlers:   make([]*handlerWrapper[WorldHandleLiquidDecayFunc], 0),
		liquidHardenHandlers:  make([]*handlerWrapper[WorldHandleLiquidHardenFunc], 0),
		soundHandlers:         make([]*handlerWrapper[WorldHandleSoundFunc], 0),
		fireSpreadHandlers:    make([]*handlerWrapper[WorldHandleFireSpreadFunc], 0),
		blockBurnHandlers:     make([]*handlerWrapper[WorldHandleBlockBurnFunc], 0),
		cropTrampleHandlers:   make([]*handlerWrapper[WorldHandleCropTrampleFunc], 0),
		entitySpawnHandlers:   make([]*handlerWrapper[WorldHandleEntitySpawnFunc], 0),
		entityDespawnHandlers: make([]*handlerWrapper[WorldHandleEntityDespawnFunc], 0),
		closeHandlers:         make([]*handlerWrapper[WorldHandleCloseFunc], 0),
	}
}

func (h *WorldHandlers) OnLiquidFlow(handler WorldHandleLiquidFlowFunc, priority Priority) {
	h.liquidFlowHandlers = append(h.liquidFlowHandlers, &handlerWrapper[WorldHandleLiquidFlowFunc]{priority, handler})
	sortHandlers(h.liquidFlowHandlers)
}

func (h *WorldHandlers) OnLiquidDecay(handler WorldHandleLiquidDecayFunc, priority Priority) {
	h.liquidDecayHandlers = append(h.liquidDecayHandlers, &handlerWrapper[WorldHandleLiquidDecayFunc]{priority, handler})
	sortHandlers(h.liquidDecayHandlers)
}

func (h *WorldHandlers) OnLiquidHarden(handler WorldHandleLiquidHardenFunc, priority Priority) {
	h.liquidHardenHandlers = append(h.liquidHardenHandlers, &handlerWrapper[WorldHandleLiquidHardenFunc]{priority, handler})
	sortHandlers(h.liquidHardenHandlers)
}

func (h *WorldHandlers) OnSound(handler WorldHandleSoundFunc, priority Priority) {
	h.soundHandlers = append(h.soundHandlers, &handlerWrapper[WorldHandleSoundFunc]{priority, handler})
	sortHandlers(h.soundHandlers)
}

func (h *WorldHandlers) OnFireSpread(handler WorldHandleFireSpreadFunc, priority Priority) {
	h.fireSpreadHandlers = append(h.fireSpreadHandlers, &handlerWrapper[WorldHandleFireSpreadFunc]{priority, handler})
	sortHandlers(h.fireSpreadHandlers)
}

func (h *WorldHandlers) OnBlockBurn(handler WorldHandleBlockBurnFunc, priority Priority) {
	h.blockBurnHandlers = append(h.blockBurnHandlers, &handlerWrapper[WorldHandleBlockBurnFunc]{priority, handler})
	sortHandlers(h.blockBurnHandlers)
}

func (h *WorldHandlers) OnCropTrample(handler WorldHandleCropTrampleFunc, priority Priority) {
	h.cropTrampleHandlers = append(h.cropTrampleHandlers, &handlerWrapper[WorldHandleCropTrampleFunc]{priority, handler})
	sortHandlers(h.cropTrampleHandlers)
}

func (h *WorldHandlers) OnEntitySpawn(handler WorldHandleEntitySpawnFunc, priority Priority) {
	h.entitySpawnHandlers = append(h.entitySpawnHandlers, &handlerWrapper[WorldHandleEntitySpawnFunc]{priority, handler})
	sortHandlers(h.entitySpawnHandlers)
}

func (h *WorldHandlers) OnEntityDespawn(handler WorldHandleEntityDespawnFunc, priority Priority) {
	h.entityDespawnHandlers = append(h.entityDespawnHandlers, &handlerWrapper[WorldHandleEntityDespawnFunc]{priority, handler})
	sortHandlers(h.entityDespawnHandlers)
}

func (h *WorldHandlers) OnClose(handler WorldHandleCloseFunc, priority Priority) {
	h.closeHandlers = append(h.closeHandlers, &handlerWrapper[WorldHandleCloseFunc]{priority, handler})
	sortHandlers(h.closeHandlers)
}
