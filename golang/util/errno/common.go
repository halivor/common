package errno

// 通用错误
const (
	BAD_REQ             Errno = 2000
	strBadReq                 = "攻城狮罚站去了" // 请求参数错误
	NOT_EXIST           Errno = 2001
	strNotExist               = "攻城狮消失了" // 数据不存在
	EXIST               Errno = 2002
	strExist                  = "攻城狮被抱走了" // 数据已存在
	UPGRADE_MANDATORY   Errno = 2003
	strUpgradeMandatory       = "攻城狮升级去了"

	// TOKEN 相关
	TOKEN_INVALID   Errno = 2010
	strTokenInvalid       = "攻城狮考证去了" // Token校验失败
	TOKEN_EXPIRE    Errno = 2011
	strTokenExpire        = "攻城狮被欠薪了" // Token过期
)

func init() {
	// 应用错误
	Set(BAD_REQ, strBadReq)
	Set(UPGRADE_MANDATORY, strUpgradeMandatory)
	Set(NOT_EXIST, strNotExist)
	Set(EXIST, strExist)

	// TOKEN异常
	Set(TOKEN_INVALID, strTokenInvalid)
	Set(TOKEN_EXPIRE, strTokenExpire)
}
