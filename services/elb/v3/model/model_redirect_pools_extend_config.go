package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RedirectPoolsExtendConfig 参数解释：转发到的后端主机组的配置。  约束限制：当action为REDIRECT_TO_POOL时生效。
type RedirectPoolsExtendConfig struct {

	// 参数解释：是否开启url重定向。
	RewriteUrlEnable *bool `json:"rewrite_url_enable,omitempty"`

	RewriteUrlConfig *RewriteUrlConfig `json:"rewrite_url_config,omitempty"`

	InsertHeadersConfig *InsertHeadersConfig `json:"insert_headers_config,omitempty"`

	RemoveHeadersConfig *RemoveHeadersConfig `json:"remove_headers_config,omitempty"`

	TrafficLimitConfig *TrafficLimitConfig `json:"traffic_limit_config,omitempty"`
}

func (o RedirectPoolsExtendConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RedirectPoolsExtendConfig struct{}"
	}

	return strings.Join([]string{"RedirectPoolsExtendConfig", string(data)}, " ")
}
