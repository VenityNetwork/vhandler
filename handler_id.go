package vhandler

type handlerId uint8

const (
	PlayerItemDropId handlerId = iota
	PlayerMoveId
	PlayerJumpId
	PlayerTeleportId
	PlayerChangeWorldId
	PlayerToggleSprintId
	PlayerToggleSneakId
	PlayerCommandExecutionId
	PlayerTransferId
	PlayerChatId
	PlayerSkinChangeId
	PlayerStartBreakId
	PlayerBlockBreakId
	PlayerBlockPlaceId
	PlayerBlockPickId
	PlayerSignEditId
	PlayerItemPickupId
	PlayerItemUseId
	PlayerItemUseOnBlockId
	PlayerItemUseOnEntityId
	PlayerItemConsumeId
	PlayerItemDamageId
	PlayerAttackEntityId
	PlayerExperienceGainId
	PlayerPunchAirId
	PlayerHurtId
	PlayerHealId
	PlayerFoodLossId
	PlayerDeathId
	PlayerRespawnId
	PlayerQuitId
)

const (
	WorldLiquidFlowId handlerId = iota
	WorldLiquidDecayId
	WorldLiquidHardenId
	WorldSoundId
	WorldFireSpreadId
	WorldBlockBurnId
	WorldEntitySpawnId
	WorldEntityDespawnId
	WorldCloseId
)
