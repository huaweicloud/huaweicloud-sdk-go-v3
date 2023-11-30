package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateRuleInstance struct {

	// 规则实例ID
	Id *string `json:"id,omitempty"`

	// 规则实例状态
	IsValid *bool `json:"is_valid,omitempty"`
}

func (o UpdateRuleInstance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRuleInstance struct{}"
	}

	return strings.Join([]string{"UpdateRuleInstance", string(data)}, " ")
}
