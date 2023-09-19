package code


const (
	Success uint32 = iota + 10000
	LoginSuccess
)

var (
	msg = map[uint32]string{
		LoginSuccess: "登录成功",
		Success: "成功",
	}
)

func GetMeg(code uint32)string{

	return msg[code]
}