package errno

// 系统级错误
const (
	SYS_ERR        Errno = 1000 // 系统调用异常
	strSysErr            = "攻城狮学俄语去了"
	SRV_ERR        Errno = 1001 // 内部服务异常
	strSrvErr            = "攻城狮学英语去了"
	OUT_ERR        Errno = 1002 // 外部服务异常
	strOutErr            = "攻城狮打猎去了"
	DB_ERR         Errno = 1003 // 内部数据库异常
	strDbErr             = "攻城狮学法语去了"
	DATA_INVALID   Errno = 1004 // 内部数据解析异常
	strDataInvalid       = "攻城狮学德语去了"
	ERR_ORI              = 1010 // 原始错误信息
	strErrOri            = "攻城狮学习去了"
)

func init() {
	Set(SYS_ERR, strSysErr)
	Set(SRV_ERR, strSrvErr)
	Set(DB_ERR, strDbErr)
	Set(OUT_ERR, strOutErr)
	Set(DATA_INVALID, strDataInvalid)
	Set(ERR_ORI, strErrOri)
}
