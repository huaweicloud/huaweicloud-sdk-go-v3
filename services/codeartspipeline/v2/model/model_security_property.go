package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SecurityProperty 漏洞规则详情
type SecurityProperty struct {

	// 是否启用
	Enable *bool `json:"enable,omitempty"`

	Rules *SecurityRule `json:"rules,omitempty"`
}

func (o SecurityProperty) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityProperty struct{}"
	}

	return strings.Join([]string{"SecurityProperty", string(data)}, " ")
}
