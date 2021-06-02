package message

const (
	//AESKEY 秘钥长度需要为16 24 32
	AESKEY = "1jsh1d23u8720102ja22012d"
)

// Notice 通知结构体
type Notice struct {
	Title string `json:"title"`
	Info  string `json:"info"`
	Time  string `json:"time"`
}

// RequestMsg API ajax 请求结构体
type RequestMsg struct {
	Code int         `json:"code"`
	Res  interface{} `json:"res"`
	Msg  string      `json:"msg"`
}
