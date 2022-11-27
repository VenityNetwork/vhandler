package vhandler

type handlerId uint8

const (
	ItemDropId handlerId = iota
	MoveId
	JumpId
	TeleportId
	ChangeWorldId
	ToggleSprintId
	ToggleSneakId
	CommandExecutionId
	TransferId
	ChatId
	SkinChangeId
	StartBreakId
	BlockBreakId
	BlockPlaceId
	BlockPickId
	SignEditId
	ItemPickupId
	ItemUseId
	ItemUseOnBlockId
	ItemUseOnEntityId
	ItemConsumeId
	ItemDamageId
	AttackEntityId
	ExperienceGainId
	PunchAirId
	HurtId
	HealId
	FoodLossId
	DeathId
	RespawnId
	QuitId
)
