package vhandler

type PlayerHandlers struct {
	moveHandlers             []*handlerWrapper[PlayerHandleMoveFunc]
	jumpHandlers             []*handlerWrapper[PlayerHandleJumpFunc]
	teleportHandlers         []*handlerWrapper[PlayerHandleTeleportFunc]
	changeWorldHandlers      []*handlerWrapper[PlayerHandleChangeWorldFunc]
	toggleSprintHandlers     []*handlerWrapper[PlayerHandleToggleSprintFunc]
	toggleSneakHandlers      []*handlerWrapper[PlayerHandleToggleSneakFunc]
	chatHandlers             []*handlerWrapper[PlayerHandleChatFunc]
	foodLossHandlers         []*handlerWrapper[PlayerHandleFoodLossFunc]
	healHandlers             []*handlerWrapper[PlayerHandleHealFunc]
	hurtHandlers             []*handlerWrapper[PlayerHandleHurtFunc]
	deathHandlers            []*handlerWrapper[PlayerHandleDeathFunc]
	respawnHandlers          []*handlerWrapper[PlayerHandleRespawnFunc]
	skinChangeHandlers       []*handlerWrapper[PlayerHandleSkinChangeFunc]
	fireExtinguishHandlers   []*handlerWrapper[PlayerHandleFireExtinguishFunc]
	startBreakHandlers       []*handlerWrapper[PlayerHandleStartBreakFunc]
	blockBreakHandlers       []*handlerWrapper[PlayerHandleBlockBreakFunc]
	blockPlaceHandlers       []*handlerWrapper[PlayerHandleBlockPlaceFunc]
	blockPickHandlers        []*handlerWrapper[PlayerHandleBlockPickFunc]
	itemUseHandlers          []*handlerWrapper[PlayerHandleItemUseFunc]
	itemUseOnBlockHandlers   []*handlerWrapper[PlayerHandleItemUseOnBlockFunc]
	itemUseOnEntityHandlers  []*handlerWrapper[PlayerHandleItemUseOnEntityFunc]
	itemConsumeHandlers      []*handlerWrapper[PlayerHandleItemConsumeFunc]
	attackEntityHandlers     []*handlerWrapper[PlayerHandleAttackEntityFunc]
	experienceGainHandlers   []*handlerWrapper[PlayerHandleExperienceGainFunc]
	punchAirHandlers         []*handlerWrapper[PlayerHandlePunchAirFunc]
	signEditHandlers         []*handlerWrapper[PlayerHandleSignEditFunc]
	lecternPageTurnHandlers  []*handlerWrapper[PlayerHandleLecternPageTurnFunc]
	itemDamageHandlers       []*handlerWrapper[PlayerHandleItemDamageFunc]
	itemPickupHandlers       []*handlerWrapper[PlayerHandleItemPickupFunc]
	itemDropHandlers         []*handlerWrapper[PlayerHandleItemDropFunc]
	transferHandlers         []*handlerWrapper[PlayerHandleTransferFunc]
	commandExecutionHandlers []*handlerWrapper[PlayerHandleCommandExecutionFunc]
	quitHandlers             []*handlerWrapper[PlayerHandleQuitFunc]
	diagnosticsHandlers      []*handlerWrapper[PlayerHandleDiagnosticsFunc]
}

func (h *PlayerHandlers) OnMove(handler PlayerHandleMoveFunc, priority Priority) {
	h.moveHandlers = append(h.moveHandlers, &handlerWrapper[PlayerHandleMoveFunc]{priority, handler})
	sortHandlers(h.moveHandlers)
}

func (h *PlayerHandlers) OnJump(handler PlayerHandleJumpFunc, priority Priority) {
	h.jumpHandlers = append(h.jumpHandlers, &handlerWrapper[PlayerHandleJumpFunc]{priority, handler})
	sortHandlers(h.jumpHandlers)
}

func (h *PlayerHandlers) OnTeleport(handler PlayerHandleTeleportFunc, priority Priority) {
	h.teleportHandlers = append(h.teleportHandlers, &handlerWrapper[PlayerHandleTeleportFunc]{priority, handler})
	sortHandlers(h.teleportHandlers)
}

func (h *PlayerHandlers) OnChangeWorld(handler PlayerHandleChangeWorldFunc, priority Priority) {
	h.changeWorldHandlers = append(h.changeWorldHandlers, &handlerWrapper[PlayerHandleChangeWorldFunc]{priority, handler})
	sortHandlers(h.changeWorldHandlers)
}

func (h *PlayerHandlers) OnToggleSprint(handler PlayerHandleToggleSprintFunc, priority Priority) {
	h.toggleSprintHandlers = append(h.toggleSprintHandlers, &handlerWrapper[PlayerHandleToggleSprintFunc]{priority, handler})
	sortHandlers(h.toggleSprintHandlers)
}

func (h *PlayerHandlers) OnToggleSneak(handler PlayerHandleToggleSneakFunc, priority Priority) {
	h.toggleSneakHandlers = append(h.toggleSneakHandlers, &handlerWrapper[PlayerHandleToggleSneakFunc]{priority, handler})
	sortHandlers(h.toggleSneakHandlers)
}

func (h *PlayerHandlers) OnChat(handler PlayerHandleChatFunc, priority Priority) {
	h.chatHandlers = append(h.chatHandlers, &handlerWrapper[PlayerHandleChatFunc]{priority, handler})
	sortHandlers(h.chatHandlers)
}

func (h *PlayerHandlers) OnFoodLoss(handler PlayerHandleFoodLossFunc, priority Priority) {
	h.foodLossHandlers = append(h.foodLossHandlers, &handlerWrapper[PlayerHandleFoodLossFunc]{priority, handler})
	sortHandlers(h.foodLossHandlers)
}

func (h *PlayerHandlers) OnHeal(handler PlayerHandleHealFunc, priority Priority) {
	h.healHandlers = append(h.healHandlers, &handlerWrapper[PlayerHandleHealFunc]{priority, handler})
	sortHandlers(h.healHandlers)
}

func (h *PlayerHandlers) OnHurt(handler PlayerHandleHurtFunc, priority Priority) {
	h.hurtHandlers = append(h.hurtHandlers, &handlerWrapper[PlayerHandleHurtFunc]{priority, handler})
	sortHandlers(h.hurtHandlers)
}

func (h *PlayerHandlers) OnDeath(handler PlayerHandleDeathFunc, priority Priority) {
	h.deathHandlers = append(h.deathHandlers, &handlerWrapper[PlayerHandleDeathFunc]{priority, handler})
	sortHandlers(h.deathHandlers)
}

func (h *PlayerHandlers) OnRespawn(handler PlayerHandleRespawnFunc, priority Priority) {
	h.respawnHandlers = append(h.respawnHandlers, &handlerWrapper[PlayerHandleRespawnFunc]{priority, handler})
	sortHandlers(h.respawnHandlers)
}

func (h *PlayerHandlers) OnSkinChange(handler PlayerHandleSkinChangeFunc, priority Priority) {
	h.skinChangeHandlers = append(h.skinChangeHandlers, &handlerWrapper[PlayerHandleSkinChangeFunc]{priority, handler})
	sortHandlers(h.skinChangeHandlers)
}

func (h *PlayerHandlers) OnFireExtinguish(handler PlayerHandleFireExtinguishFunc, priority Priority) {
	h.fireExtinguishHandlers = append(h.fireExtinguishHandlers, &handlerWrapper[PlayerHandleFireExtinguishFunc]{priority, handler})
	sortHandlers(h.fireExtinguishHandlers)
}

func (h *PlayerHandlers) OnStartBreak(handler PlayerHandleStartBreakFunc, priority Priority) {
	h.startBreakHandlers = append(h.startBreakHandlers, &handlerWrapper[PlayerHandleStartBreakFunc]{priority, handler})
	sortHandlers(h.startBreakHandlers)
}

func (h *PlayerHandlers) OnBlockBreak(handler PlayerHandleBlockBreakFunc, priority Priority) {
	h.blockBreakHandlers = append(h.blockBreakHandlers, &handlerWrapper[PlayerHandleBlockBreakFunc]{priority, handler})
	sortHandlers(h.blockBreakHandlers)
}

func (h *PlayerHandlers) OnBlockPlace(handler PlayerHandleBlockPlaceFunc, priority Priority) {
	h.blockPlaceHandlers = append(h.blockPlaceHandlers, &handlerWrapper[PlayerHandleBlockPlaceFunc]{priority, handler})
	sortHandlers(h.blockPlaceHandlers)
}

func (h *PlayerHandlers) OnBlockPick(handler PlayerHandleBlockPickFunc, priority Priority) {
	h.blockPickHandlers = append(h.blockPickHandlers, &handlerWrapper[PlayerHandleBlockPickFunc]{priority, handler})
	sortHandlers(h.blockPickHandlers)
}

func (h *PlayerHandlers) OnItemUse(handler PlayerHandleItemUseFunc, priority Priority) {
	h.itemUseHandlers = append(h.itemUseHandlers, &handlerWrapper[PlayerHandleItemUseFunc]{priority, handler})
	sortHandlers(h.itemUseHandlers)
}

func (h *PlayerHandlers) OnItemUseOnBlock(handler PlayerHandleItemUseOnBlockFunc, priority Priority) {
	h.itemUseOnBlockHandlers = append(h.itemUseOnBlockHandlers, &handlerWrapper[PlayerHandleItemUseOnBlockFunc]{priority, handler})
	sortHandlers(h.itemUseOnBlockHandlers)
}

func (h *PlayerHandlers) OnItemUseOnEntity(handler PlayerHandleItemUseOnEntityFunc, priority Priority) {
	h.itemUseOnEntityHandlers = append(h.itemUseOnEntityHandlers, &handlerWrapper[PlayerHandleItemUseOnEntityFunc]{priority, handler})
	sortHandlers(h.itemUseOnEntityHandlers)
}

func (h *PlayerHandlers) OnItemConsume(handler PlayerHandleItemConsumeFunc, priority Priority) {
	h.itemConsumeHandlers = append(h.itemConsumeHandlers, &handlerWrapper[PlayerHandleItemConsumeFunc]{priority, handler})
	sortHandlers(h.itemConsumeHandlers)
}

func (h *PlayerHandlers) OnAttackEntity(handler PlayerHandleAttackEntityFunc, priority Priority) {
	h.attackEntityHandlers = append(h.attackEntityHandlers, &handlerWrapper[PlayerHandleAttackEntityFunc]{priority, handler})
	sortHandlers(h.attackEntityHandlers)
}

func (h *PlayerHandlers) OnExperienceGain(handler PlayerHandleExperienceGainFunc, priority Priority) {
	h.experienceGainHandlers = append(h.experienceGainHandlers, &handlerWrapper[PlayerHandleExperienceGainFunc]{priority, handler})
	sortHandlers(h.experienceGainHandlers)
}

func (h *PlayerHandlers) OnPunchAir(handler PlayerHandlePunchAirFunc, priority Priority) {
	h.punchAirHandlers = append(h.punchAirHandlers, &handlerWrapper[PlayerHandlePunchAirFunc]{priority, handler})
	sortHandlers(h.punchAirHandlers)
}

func (h *PlayerHandlers) OnSignEdit(handler PlayerHandleSignEditFunc, priority Priority) {
	h.signEditHandlers = append(h.signEditHandlers, &handlerWrapper[PlayerHandleSignEditFunc]{priority, handler})
	sortHandlers(h.signEditHandlers)
}

func (h *PlayerHandlers) OnLecternPageTurn(handler PlayerHandleLecternPageTurnFunc, priority Priority) {
	h.lecternPageTurnHandlers = append(h.lecternPageTurnHandlers, &handlerWrapper[PlayerHandleLecternPageTurnFunc]{priority, handler})
	sortHandlers(h.lecternPageTurnHandlers)
}

func (h *PlayerHandlers) OnItemDamage(handler PlayerHandleItemDamageFunc, priority Priority) {
	h.itemDamageHandlers = append(h.itemDamageHandlers, &handlerWrapper[PlayerHandleItemDamageFunc]{priority, handler})
	sortHandlers(h.itemDamageHandlers)
}

func (h *PlayerHandlers) OnItemPickup(handler PlayerHandleItemPickupFunc, priority Priority) {
	h.itemPickupHandlers = append(h.itemPickupHandlers, &handlerWrapper[PlayerHandleItemPickupFunc]{priority, handler})
	sortHandlers(h.itemPickupHandlers)
}

func (h *PlayerHandlers) OnItemDrop(handler PlayerHandleItemDropFunc, priority Priority) {
	h.itemDropHandlers = append(h.itemDropHandlers, &handlerWrapper[PlayerHandleItemDropFunc]{priority, handler})
	sortHandlers(h.itemDropHandlers)
}

func (h *PlayerHandlers) OnTransfer(handler PlayerHandleTransferFunc, priority Priority) {
	h.transferHandlers = append(h.transferHandlers, &handlerWrapper[PlayerHandleTransferFunc]{priority, handler})
	sortHandlers(h.transferHandlers)
}

func (h *PlayerHandlers) OnCommandExecution(handler PlayerHandleCommandExecutionFunc, priority Priority) {
	h.commandExecutionHandlers = append(h.commandExecutionHandlers, &handlerWrapper[PlayerHandleCommandExecutionFunc]{priority, handler})
	sortHandlers(h.commandExecutionHandlers)
}

func (h *PlayerHandlers) OnQuit(handler PlayerHandleQuitFunc, priority Priority) {
	h.quitHandlers = append(h.quitHandlers, &handlerWrapper[PlayerHandleQuitFunc]{priority, handler})
	sortHandlers(h.quitHandlers)
}

func (h *PlayerHandlers) OnDiagnostics(handler PlayerHandleDiagnosticsFunc, priority Priority) {
	h.diagnosticsHandlers = append(h.diagnosticsHandlers, &handlerWrapper[PlayerHandleDiagnosticsFunc]{priority, handler})
	sortHandlers(h.diagnosticsHandlers)
}
