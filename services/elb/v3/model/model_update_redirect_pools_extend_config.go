package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateRedirectPoolsExtendConfig 参数解释：转发到的后端主机组的配置。  约束限制：当action为REDIRECT_TO_POOL时生效。
type UpdateRedirectPoolsExtendConfig struct {

	// 参数解释：是否开启url重定向
	RewriteUrlEnable *bool `json:"rewrite_url_enable,omitempty"`

	RewriteUrlConfig *UpdateRewriteUrlConfig `json:"rewrite_url_config,omitempty"`

	InsertHeadersConfig *UpdateInsertHeadersConfig `json:"insert_headers_config,omitempty"`

	RemoveHeadersConfig *UpdateRemoveHeadersConfig `json:"remove_headers_config,omitempty"`

	TrafficLimitConfig *UpdateTrafficLimitConfig `json:"traffic_limit_config,omitempty"`

	CorsConfig *CreateCorsConfig `json:"cors_config,omitempty"`

	TrafficMirrorConfig *CreateTrafficMirrorConfig `json:"traffic_mirror_config,omitempty"`
}

func (o UpdateRedirectPoolsExtendConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRedirectPoolsExtendConfig struct{}"
	}

	return strings.Join([]string{"UpdateRedirectPoolsExtendConfig", string(data)}, " ")
}
