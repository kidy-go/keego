package mux

import (
	"mime/multipart"
	"net/http"
	"net/url"
)

type AContext interface {
	// 保留原始的`net/http`对象
	// Request *http.Request
	// ResponseWriter http.ResponseWriter
	Http() *HttpContext

	// 获取封装后的Request对象, 实现对 *http.Request 的高级封装
	Request() Request

	// 写入到http响应中
	Response(interface{}, int) Response
	Get(string) interface{}
	// 获取客户端上传的文件
	// TODO:
	// 需要返回一个封装后的文件对象, 用以更高级的操作
	File(string) (multipart.File, *multipart.FileHeader, error)
}

type (
	Maps        map[string]interface{}
	HttpContext struct {
		Request        *http.Request
		ResponseWriter http.ResponseWriter
	}
	Request    struct{}
	Response   struct{}
	MuxContext struct {
		params   Maps
		query    url.Values
		request  *http.Request
		response http.ResponseWriter
	}

	Context interface {
		Params() Maps
		SetParam(string, string)
	}
)

func (c *MuxContext) SetParam(key, val string) {
	c.params[key] = val
}

func (c *MuxContext) Params() Maps {
	return c.params
}

func (c *MuxContext) Param(name string, defaultValue ...interface{}) interface{} {
	if val, ok := c.params[name]; ok {
		return val
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return nil
}
