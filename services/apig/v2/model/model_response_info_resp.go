package model

import (
	"encoding/json"

	"strings"
)

type ResponseInfoResp struct {
	// 响应的HTTP状态码

	Status *int32 `json:"status,omitempty"`
	// 响应的Body模板

	Body *string `json:"body,omitempty"`
	// 是否为默认响应

	Default *bool `json:"default,omitempty"`
}

func (o ResponseInfoResp) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ResponseInfoResp struct{}"
	}

	return strings.Join([]string{"ResponseInfoResp", string(data)}, " ")
}
