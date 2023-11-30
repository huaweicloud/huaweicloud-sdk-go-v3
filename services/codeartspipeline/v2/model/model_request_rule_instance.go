package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RequestRuleInstance struct {

	// 规则实例ID
	Id *string `json:"id,omitempty"`

	// 规则实例状态
	IsValid *bool `json:"is_valid,omitempty"`
}

func (o RequestRuleInstance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RequestRuleInstance struct{}"
	}

	return strings.Join([]string{"RequestRuleInstance", string(data)}, " ")
}
