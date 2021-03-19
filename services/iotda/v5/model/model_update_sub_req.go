package model

import (
	"encoding/json"

	"strings"
)

// 修改订阅主题请求结构体
type UpdateSubReq struct {
	// 订阅的回调地址，用于接收对应资源事件的通知消息，例如：https://10.10.10.10:443/callbackurltest。

	Callbackurl *string `json:"callbackurl,omitempty"`
}

func (o UpdateSubReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateSubReq struct{}"
	}

	return strings.Join([]string{"UpdateSubReq", string(data)}, " ")
}
