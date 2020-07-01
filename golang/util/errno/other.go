package errno

const (
	// 关系相关
	MGR_BAN   Errno = 2200
	strMgrBan       = "您已被房管禁言"
	ADM_BAN   Errno = 2201
	strAdmBan       = "您已被超管禁言"

	// 好友相关
	HAD_BEEN_MATE  Errno = 2400
	strHadBeenMate       = "该用户已经是您的好友"
	HAD_APPLIED    Errno = 2401
	strHadApplied        = "已经发起过好友申请"

	// 其他相关
	FEED_BACK_IMAGE  Errno = 2300
	strFeedBackImage       = "获取反馈头像失败"
)

func init() {
	// 关系相关
	Set(MGR_BAN, strMgrBan)
	Set(ADM_BAN, strAdmBan)

	// 好友相关
	Set(HAD_BEEN_MATE, strHadBeenMate)
	Set(HAD_APPLIED, strHadApplied)

	// 配置反馈相关
	Set(FEED_BACK_IMAGE, strFeedBackImage)
}
