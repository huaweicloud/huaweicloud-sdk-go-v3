package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 配置项。
type ConfigsGetBody struct {
	// 回源请求头配置

	OriginRequestHeader *[]OriginRequestHeader `json:"origin_request_header,omitempty"`
	// http header配置

	HttpResponseHeader *[]HttpResponseHeader `json:"http_response_header,omitempty"`

	UrlAuth *UrlAuthGetBody `json:"url_auth,omitempty"`

	Https *HttpGetBody `json:"https,omitempty"`
}

func (o ConfigsGetBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigsGetBody struct{}"
	}

	return strings.Join([]string{"ConfigsGetBody", string(data)}, " ")
}
