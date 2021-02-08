package router

type (
	ResponseHeader  map[string][]string
	ResponseBody    interface{}
	ResponseCookies map[string][]string

	Response struct {
		Header ResponseHeader
		Body   ResponseBody
	}
)
