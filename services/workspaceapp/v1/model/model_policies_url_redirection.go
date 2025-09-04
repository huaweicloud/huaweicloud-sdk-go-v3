package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PoliciesUrlRedirection 配置url重定向。
type PoliciesUrlRedirection struct {

	// 配置url重定向开关： false: 表示关闭 true: 表示开启
	UrlRedirectionEnable *bool `json:"url_redirection_enable,omitempty"`

	Options *UrlRedirectionOptions `json:"options,omitempty"`
}

func (o PoliciesUrlRedirection) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoliciesUrlRedirection struct{}"
	}

	return strings.Join([]string{"PoliciesUrlRedirection", string(data)}, " ")
}
