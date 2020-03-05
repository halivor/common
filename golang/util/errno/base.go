package errno

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

func init() {
	Set(SYS_ERR, strSysErr)
	Set(SRV_ERR, strSrvErr)
	Set(DB_ERR, strDbErr)
	Set(OUT_ERR, strOutErr)
	Set(DATA_INVALID, strDataInvalid)
	Set(ERR_ORI, strErrOri)
}
