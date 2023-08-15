package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateRedirectPoolsExtendConfig struct {

	// url重写的开关
	RewriteUrlEnable *bool `json:"rewrite_url_enable,omitempty"`

	RewriteUrlConfig *CreateRewriteUrlConfig `json:"rewrite_url_config,omitempty"`
}

func (o CreateRedirectPoolsExtendConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRedirectPoolsExtendConfig struct{}"
	}

	return strings.Join([]string{"CreateRedirectPoolsExtendConfig", string(data)}, " ")
}
