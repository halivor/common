package errno

type Errno int

const (
	SUCC    Errno = 0
	strSucc       = "成功"
)

// 系统级错误
const (
	SYS_ERR        Errno = 1000
	strSysErr            = "攻城狮学俄语去了" // 系统调用异常
	SRV_ERR        Errno = 1001
	strSrvErr            = "攻城狮学英语去了" // 内部服务异常
	OUT_ERR        Errno = 1002
	strOutErr            = "攻城狮打猎去了" // 外部服务异常
	DB_ERR         Errno = 1003
	strDbErr             = "攻城狮学法语去了" // 内部数据库异常
	DATA_INVALID   Errno = 1004
	strDataInvalid       = "攻城狮学德语去了" // 内部数据解析异常
	ERR_ORI              = 1010
	strErrOri            = "攻城狮学习去了" // 原始错误信息
)

// 应用级错误
const (
	// 通用错误
	BAD_REQ     Errno = 2000
	strBadReq         = "攻城狮罚站去了" // 请求参数错误
	NOT_EXIST   Errno = 2001
	strNotExist       = "攻城狮消失了" // 数据不存在
	// TOKEN 相关
	TOKEN_INVALID   Errno = 2010
	strTokenInvalid       = "攻城狮考证去了" // Token校验失败
	TOKEN_EXPIRE    Errno = 2011
	strTokenExpire        = "攻城狮被欠薪了" // Token过期
)
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

	// 其他相关
	FEED_BACK_IMAGE  Errno = 2300
	strFeedBackImage       = "获取反馈头像失败"
)

var errmsg map[Errno]string = map[Errno]string{}

func Set(no Errno, msg string) {
	errmsg[no] = msg
}

func init() {
	errmsg[SUCC] = strSucc
	// 服务错误
	errmsg[SYS_ERR] = strSysErr
	errmsg[SRV_ERR] = strSrvErr
	errmsg[DB_ERR] = strDbErr
	errmsg[OUT_ERR] = strOutErr
	errmsg[DATA_INVALID] = strDataInvalid
	errmsg[ERR_ORI] = strErrOri
	// 应用错误
	errmsg[BAD_REQ] = strBadReq
	errmsg[NOT_EXIST] = strNotExist
	// TOKEN异常
	errmsg[TOKEN_INVALID] = strTokenInvalid
	errmsg[TOKEN_EXPIRE] = strTokenExpire
	// 个人信息异常
	errmsg[NICK_EXIST] = strNickExist
	errmsg[NICK_INVALID] = strNickInvalid
	errmsg[NICK_LENGTH_INVALID] = strNickLengthInvalid
	errmsg[GENDER_ERROR] = strGenderError
	errmsg[AVATAR_FAILED] = strAvatarFailed
	// 账户相关
	errmsg[NO_COIN] = strNoCoin

	// 房间相关
	errmsg[LIVING_OFF] = strLivingOff
	// 关系相关
	errmsg[MGR_BAN] = strMgrBan
	errmsg[ADM_BAN] = strAdmBan
	// CDN
	errmsg[NOT_LIVING] = strNotLiving

	// 好友相关
	errmsg[HAD_BEEN_MATE] = strHadBeenMate

	// 配置反馈相关
	errmsg[FEED_BACK_IMAGE] = strFeedBackImage
}

func (e Errno) Errno() int {
	return int(e)
}

func (e Errno) String() string {
	return errmsg[e]
}

func (e Errno) Error() string {
	return errmsg[e]
}
