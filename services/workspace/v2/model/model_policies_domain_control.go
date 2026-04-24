package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PoliciesDomainControl 上网DNS管控策略
type PoliciesDomainControl struct {

	// 默认开关
	DefaultEnabled *bool `json:"default_enabled,omitempty"`

	// 域名
	DomainRules *string `json:"domain_rules,omitempty"`
}

func (o PoliciesDomainControl) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoliciesDomainControl struct{}"
	}

	return strings.Join([]string{"PoliciesDomainControl", string(data)}, " ")
}
