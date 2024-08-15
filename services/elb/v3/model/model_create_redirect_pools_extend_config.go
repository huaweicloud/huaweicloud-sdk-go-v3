package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateRedirectPoolsExtendConfig 参数解释：转发到的后端主机组的配置。  约束限制：当action为REDIRECT_TO_POOL时生效。
type CreateRedirectPoolsExtendConfig struct {

	// 参数解释：是否开启url重定向。
	RewriteUrlEnable *bool `json:"rewrite_url_enable,omitempty"`

	RewriteUrlConfig *CreateRewriteUrlConfig `json:"rewrite_url_config,omitempty"`

	InsertHeadersConfig *CreateInsertHeadersConfig `json:"insert_headers_config,omitempty"`

	RemoveHeadersConfig *CreateRemoveHeadersConfig `json:"remove_headers_config,omitempty"`

	TrafficLimitConfig *CreateTrafficLimitConfig `json:"traffic_limit_config,omitempty"`
}

func (o CreateRedirectPoolsExtendConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRedirectPoolsExtendConfig struct{}"
	}

	return strings.Join([]string{"CreateRedirectPoolsExtendConfig", string(data)}, " ")
}
