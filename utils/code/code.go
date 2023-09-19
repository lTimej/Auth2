package code

const (
	Success uint32 = iota + 10000
	ApiError
)

var (
	msg = map[uint32]string{
		Success:  "成功",
		ApiError: "接口错误",
	}
)

func GetMeg(code uint32) string {

	return msg[code]
}
