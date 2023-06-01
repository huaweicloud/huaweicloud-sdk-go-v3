package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 转发到的后端主机组的配置。当action为REDIRECT_TO_POOL时生效。
type RedirectPoolsExtendConfig struct {

	// 是否开启url重定向
	RewriteUrlEnable *bool `json:"rewrite_url_enable,omitempty"`

	RewriteUrlConfig *RewriteUrlConfig `json:"rewrite_url_config,omitempty"`
}

func (o RedirectPoolsExtendConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RedirectPoolsExtendConfig struct{}"
	}

	return strings.Join([]string{"RedirectPoolsExtendConfig", string(data)}, " ")
}
