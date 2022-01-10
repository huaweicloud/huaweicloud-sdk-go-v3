package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 配置项。
type Configs struct {
	// 回源请求头配置

	OriginRequestHeader *[]OriginRequestHeader `json:"origin_request_header,omitempty"`
	// http header配置

	HttpResponseHeader *[]HttpResponseHeader `json:"http_response_header,omitempty"`

	UrlAuth *UrlAuth `json:"url_auth,omitempty"`
}

func (o Configs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Configs struct{}"
	}

	return strings.Join([]string{"Configs", string(data)}, " ")
}
