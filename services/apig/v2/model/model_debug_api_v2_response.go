package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DebugApiV2Response struct {
	// 调试请求报文内容

	Request *string `json:"request,omitempty"`
	// 调试响应报文内容，响应消息体最大支持2097152字节，超过部分会被截断 > 响应消息体超过限制长度时，超过部分会被截断，并追加\"[TRUNCATED]\"信息。

	Response *string `json:"response,omitempty"`
	// 调试耗时，单位：毫秒

	Latency *int32 `json:"latency,omitempty"`
	// 调试过程日志

	Log            *string `json:"log,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DebugApiV2Response) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DebugApiV2Response struct{}"
	}

	return strings.Join([]string{"DebugApiV2Response", string(data)}, " ")
}
