package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InterfacesRequest 接口请求。
type InterfacesRequest struct {

	// 方法。
	Method *string `json:"method,omitempty"`

	// 请求参数。
	Params map[string]string `json:"params,omitempty"`

	// 请求头信息。
	Headers map[string]string `json:"headers,omitempty"`

	// 请求体。
	Body *string `json:"body,omitempty"`

	// URL。
	Url *string `json:"url,omitempty"`

	// 配置。
	Configs map[string]interface{} `json:"configs,omitempty"`
}

func (o InterfacesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InterfacesRequest struct{}"
	}

	return strings.Join([]string{"InterfacesRequest", string(data)}, " ")
}
