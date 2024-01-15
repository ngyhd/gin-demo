package response

var message map[Code]string

type Code int64

const SuccessCode Code = 0

const (
	ParamsErrCode Code = 40000
)

const (
	InternalErrCode Code = 50000
)

func init() {
	message = map[Code]string{}
	message[SuccessCode] = "Success"
	message[ParamsErrCode] = "参数错误"
	message[InternalErrCode] = "系统内部发生错误"
}

// 接口影响正常时code为0，错误时code为5位数
// 第1位,系统错误类型。5:系统内部错误 4:用户请求错误
// 第2~3位,业务类型错误。比如,00用户相关错误
// 第4~5位,具体错误。比如,00参数错误 01注册系统内部错误
// 例子:
// 40000:请求参数错误
// 50001:用户注册系统内部错误

func Success(d any) map[string]any {
	m := map[string]any{}
	m["code"] = 0
	m["msg"] = message[SuccessCode]
	m["data"] = d
	return m
}

func Fail(code Code) map[string]any {
	m := map[string]any{}
	m["code"] = code
	m["msg"] = message[code]
	m["data"] = ""
	return m
}
