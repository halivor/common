package errno

const (
	// 房间相关
	LIVING_OFF   Errno = 2100
	strLivingOff       = "主播关播中"
	// 流相关
	NOT_LIVING   Errno = 2110
	strNotLiving       = "主播未开播"
)

func init() {
	// 房间相关
	Set(LIVING_OFF, strLivingOff)
	// CDN
	Set(NOT_LIVING, strNotLiving)
}
