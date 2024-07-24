package errcode

var (
	ERR_MOVE                = CreateErrCode(1101, "移动错误")
	ERR_OPEN_WALL           = CreateErrCode(1102, "砸墙错误")
	ERR_STRENGTH_NOT_ENOUGH = CreateErrCode(1103, "体力不足")
	ERR_POINT               = CreateErrCode(1104, "目标错误")
	ERR_NO_DEAD             = CreateErrCode(1105, "玩家未死亡")
	ERR_GM_PARAM            = CreateErrCode(1106, "GM参数错误")
	ERR_DEAD                = CreateErrCode(1107, "玩家已死亡")
	ERR_GM_CLOSE            = CreateErrCode(1108, "GM未开启")
	ERR_GOLD_NOT_ENOUGH     = CreateErrCode(1109, "金币不足")
	ERR_LV_MAX              = CreateErrCode(1110, "已达最大等级")
	ERR_CONF_NIL            = CreateErrCode(1111, "配置未找到")
	ERR_DIAMOND_NOT_ENOUGH  = CreateErrCode(1112, "钻石不足")
)
