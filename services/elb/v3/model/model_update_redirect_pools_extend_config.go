package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateRedirectPoolsExtendConfig 转发到的后端主机组的URL配置。
type UpdateRedirectPoolsExtendConfig struct {

	// 参数解释：是否开启url重定向
	RewriteUrlEnable *bool `json:"rewrite_url_enable,omitempty"`

	RewriteUrlConfig *UpdateRewriteUrlConfig `json:"rewrite_url_config,omitempty"`

	InsertHeadersConfig *UpdateInsertHeadersConfig `json:"insert_headers_config,omitempty"`

	RemoveHeadersConfig *UpdateRemoveHeadersConfig `json:"remove_headers_config,omitempty"`

	TrafficLimitConfig *UpdateTrafficLimitConfig `json:"traffic_limit_config,omitempty"`
}

func (o UpdateRedirectPoolsExtendConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRedirectPoolsExtendConfig struct{}"
	}

	return strings.Join([]string{"UpdateRedirectPoolsExtendConfig", string(data)}, " ")
}
