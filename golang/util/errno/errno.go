package errno

import (
	"fmt"
)

type Errno int

const (
	SUCC    Errno = 0
	strSucc       = "成功"
)

// 应用级错误
const (
	// 个人信息相关
	NICK_EXIST           Errno = 2020
	strNickExist               = "昵称已存在"
	NICK_INVALID         Errno = 2021
	strNickInvalid             = "昵称内容不合法"
	NICK_LENGTH_INVALID  Errno = 2022
	strNickLengthInvalid       = "昵称长度不合法"
	GENDER_ERROR         Errno = 2025
	strGenderError             = "性别不合法"
	AVATAR_FAILED        Errno = 2026
	strAvatarFailed            = "获取头像失败"
	// 账户相关
	NO_COIN   Errno = 2050
	strNoCoin       = "余额不足"

	// 房间相关
	LIVING_OFF   Errno = 2100
	strLivingOff       = "主播关播中"
	// 流相关
	NOT_LIVING   Errno = 2110
	strNotLiving       = "主播未开播"
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

var errmsg = make(map[Errno]string, 1024)

func Set(no Errno, msg string) {
	errmsg[no] = msg
}

func init() {
	Set(SUCC, strSucc)
	other()
}
func other() {
	// 个人信息异常
	Set(NICK_EXIST, strNickExist)
	Set(NICK_INVALID, strNickInvalid)
	Set(NICK_LENGTH_INVALID, strNickLengthInvalid)
	Set(GENDER_ERROR, strGenderError)
	Set(AVATAR_FAILED, strAvatarFailed)
	// 账户相关
	Set(NO_COIN, strNoCoin)

	// 房间相关
	Set(LIVING_OFF, strLivingOff)
	// 关系相关
	Set(MGR_BAN, strMgrBan)
	Set(ADM_BAN, strAdmBan)
	// CDN
	Set(NOT_LIVING, strNotLiving)

	// 好友相关
	Set(HAD_BEEN_MATE, strHadBeenMate)
	Set(HAD_APPLIED, strHadApplied)

	// 配置反馈相关
	Set(FEED_BACK_IMAGE, strFeedBackImage)
}

func (e Errno) Errno() int {
	return int(e)
}

func (e Errno) String() string {
	return fmt.Sprintf("%d->%s", e.Errno(), errmsg[e])
}

func (e Errno) Error() string {
	return errmsg[e]
}
