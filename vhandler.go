package vhandler

import "github.com/df-mc/dragonfly/server/player"

type VHandler struct {
	handlers map[handlerId]*subHandler
	main     *MainHandler
}

func New() *VHandler {
	v := &VHandler{handlers: map[handlerId]*subHandler{}}
	v.main = newMainHandler(v)

	v.handlers[ItemDropId] = newSubHandler()
	v.handlers[MoveId] = newSubHandler()
	v.handlers[JumpId] = newSubHandler()
	v.handlers[TeleportId] = newSubHandler()
	v.handlers[ChangeWorldId] = newSubHandler()
	v.handlers[ToggleSprintId] = newSubHandler()
	v.handlers[ToggleSneakId] = newSubHandler()
	v.handlers[CommandExecutionId] = newSubHandler()
	v.handlers[TransferId] = newSubHandler()
	v.handlers[ChatId] = newSubHandler()
	v.handlers[SkinChangeId] = newSubHandler()
	v.handlers[StartBreakId] = newSubHandler()
	v.handlers[BlockBreakId] = newSubHandler()
	v.handlers[BlockPlaceId] = newSubHandler()
	v.handlers[BlockPickId] = newSubHandler()
	v.handlers[SignEditId] = newSubHandler()
	v.handlers[ItemPickupId] = newSubHandler()
	v.handlers[ItemUseId] = newSubHandler()
	v.handlers[ItemUseOnBlockId] = newSubHandler()
	v.handlers[ItemUseOnEntityId] = newSubHandler()
	v.handlers[ItemConsumeId] = newSubHandler()
	v.handlers[ItemDamageId] = newSubHandler()
	v.handlers[AttackEntityId] = newSubHandler()
	v.handlers[ExperienceGainId] = newSubHandler()
	v.handlers[PunchAirId] = newSubHandler()
	v.handlers[HurtId] = newSubHandler()
	v.handlers[HealId] = newSubHandler()
	v.handlers[FoodLossId] = newSubHandler()
	v.handlers[DeathId] = newSubHandler()
	v.handlers[RespawnId] = newSubHandler()
	v.handlers[QuitId] = newSubHandler()

	return v
}

func (v *VHandler) OnItemDrop(p priorityId, h ItemDropHandler) {
	v.handlers[ItemDropId].add(p, h)
}

func (v *VHandler) OnMove(p priorityId, h MoveHandler) {
	v.handlers[MoveId].add(p, h)
}

func (v *VHandler) OnJump(p priorityId, h JumpHandler) {
	v.handlers[JumpId].add(p, h)
}

func (v *VHandler) OnTeleport(p priorityId, h TeleportHandler) {
	v.handlers[TeleportId].add(p, h)
}

func (v *VHandler) OnChangeWorld(p priorityId, h ChangeWorldHandler) {
	v.handlers[ChangeWorldId].add(p, h)
}

func (v *VHandler) OnToggleSprint(p priorityId, h ToggleSprintHandler) {
	v.handlers[ToggleSprintId].add(p, h)
}

func (v *VHandler) OnToggleSneak(p priorityId, h ToggleSneakHandler) {
	v.handlers[ToggleSneakId].add(p, h)
}

func (v *VHandler) OnCommandExecution(p priorityId, h CommandExecutionHandler) {
	v.handlers[CommandExecutionId].add(p, h)
}

func (v *VHandler) OnTransfer(p priorityId, h TransferHandler) {
	v.handlers[TransferId].add(p, h)
}

func (v *VHandler) OnChat(p priorityId, h ChatHandler) {
	v.handlers[ChatId].add(p, h)
}

func (v *VHandler) OnSkinChange(p priorityId, h SkinChangeHandler) {
	v.handlers[SkinChangeId].add(p, h)
}

func (v *VHandler) OnStartBreak(p priorityId, h StartBreakHandler) {
	v.handlers[StartBreakId].add(p, h)
}

func (v *VHandler) OnBlockBreak(p priorityId, h BlockBreakHandler) {
	v.handlers[BlockBreakId].add(p, h)
}

func (v *VHandler) OnBlockPlace(p priorityId, h BlockPlaceHandler) {
	v.handlers[BlockPlaceId].add(p, h)
}

func (v *VHandler) OnBlockPick(p priorityId, h BlockPickHandler) {
	v.handlers[BlockPickId].add(p, h)
}

func (v *VHandler) OnSignEdit(p priorityId, h SignEditHandler) {
	v.handlers[SignEditId].add(p, h)
}

func (v *VHandler) OnItemPickup(p priorityId, h ItemPickupHandler) {
	v.handlers[ItemPickupId].add(p, h)
}

func (v *VHandler) OnItemUse(p priorityId, h ItemUseHandler) {
	v.handlers[ItemUseId].add(p, h)
}

func (v *VHandler) OnItemUseOnBlock(p priorityId, h ItemUseOnBlockHandler) {
	v.handlers[ItemUseOnBlockId].add(p, h)
}

func (v *VHandler) OnItemUseOnEntity(p priorityId, h ItemUseOnEntityHandler) {
	v.handlers[ItemUseOnEntityId].add(p, h)
}

func (v *VHandler) OnItemConsume(p priorityId, h ItemConsumeHandler) {
	v.handlers[ItemConsumeId].add(p, h)
}

func (v *VHandler) OnItemDamage(p priorityId, h ItemDamageHandler) {
	v.handlers[ItemDamageId].add(p, h)
}

func (v *VHandler) OnAttackEntity(p priorityId, h AttackEntityHandler) {
	v.handlers[AttackEntityId].add(p, h)
}

func (v *VHandler) OnExperienceGain(p priorityId, h ExperienceGainHandler) {
	v.handlers[ExperienceGainId].add(p, h)
}

func (v *VHandler) OnPunchAir(p priorityId, h PunchAirHandler) {
	v.handlers[PunchAirId].add(p, h)
}

func (v *VHandler) OnHurt(p priorityId, h HurtHandler) {
	v.handlers[HurtId].add(p, h)
}

func (v *VHandler) OnHeal(p priorityId, h HealHandler) {
	v.handlers[HealId].add(p, h)
}

func (v *VHandler) OnFoodLoss(p priorityId, h FoodLossHandler) {
	v.handlers[FoodLossId].add(p, h)
}

func (v *VHandler) OnDeath(p priorityId, h DeathHandler) {
	v.handlers[DeathId].add(p, h)
}

func (v *VHandler) OnRespawn(p priorityId, h RespawnHandler) {
	v.handlers[RespawnId].add(p, h)
}

func (v *VHandler) OnQuit(p priorityId, h QuitHandler) {
	v.handlers[QuitId].add(p, h)
}

func (v *VHandler) Accept(p *player.Player) {
	p.Handle(v.main)
}
