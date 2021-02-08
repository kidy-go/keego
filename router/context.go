package router

type (
	Context interface {
		Request()
		Response(interface{}, int)
		Get(string) FileStore
		// 获取客户端上传的文件
		// TODO:
		// 需要返回一个封装后的文件系统对象
		File(string) interface{}
		Cookie(string) string
		Header(string) string
		Params()
	}
)
