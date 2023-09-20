package code

const (
	Success uint32 = iota + 10000
	ApiError
	ParamsError
	DBError
)

var (
	msg = map[uint32]string{
		Success:     "成功",
		ApiError:    "接口错误",
		ParamsError: "参数错误",
		DBError:     "数据库异常",
	}
)

func GetMeg(code uint32) string {
	return msg[code]
}
