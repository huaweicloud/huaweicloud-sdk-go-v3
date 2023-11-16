package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SecurityRuleCve 漏洞编号
type SecurityRuleCve struct {

	// 是否启用
	Enable *bool `json:"enable,omitempty"`

	// 漏洞编号
	Values *[]string `json:"values,omitempty"`
}

func (o SecurityRuleCve) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityRuleCve struct{}"
	}

	return strings.Join([]string{"SecurityRuleCve", string(data)}, " ")
}
