package errno

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
)

func init() {
	// 个人信息异常
	Set(NICK_EXIST, strNickExist)
	Set(NICK_INVALID, strNickInvalid)
	Set(NICK_LENGTH_INVALID, strNickLengthInvalid)
	Set(GENDER_ERROR, strGenderError)
	Set(AVATAR_FAILED, strAvatarFailed)
	// 账户相关
	Set(NO_COIN, strNoCoin)
}
