package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RedirectPoolsExtendConfig **参数解释**：转发到的后端服务器组的配置。
type RedirectPoolsExtendConfig struct {

	// **参数解释**：是否开启url重定向。  **取值范围**：true 开启，false 未开启。
	RewriteUrlEnable *bool `json:"rewrite_url_enable,omitempty"`

	RewriteUrlConfig *RewriteUrlConfig `json:"rewrite_url_config,omitempty"`

	InsertHeadersConfig *InsertHeadersConfig `json:"insert_headers_config,omitempty"`

	RemoveHeadersConfig *RemoveHeadersConfig `json:"remove_headers_config,omitempty"`

	TrafficLimitConfig *TrafficLimitConfig `json:"traffic_limit_config,omitempty"`

	CorsConfig *CorsConfig `json:"cors_config,omitempty"`

	TrafficMirrorConfig *TrafficMirrorConfig `json:"traffic_mirror_config,omitempty"`
}

func (o RedirectPoolsExtendConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RedirectPoolsExtendConfig struct{}"
	}

	return strings.Join([]string{"RedirectPoolsExtendConfig", string(data)}, " ")
}
