package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateRuleSetReq struct {

	// 规则模版实例名称
	Name string `json:"name"`

	// 规则实例列表
	Rules *[]UpdateRuleInstance `json:"rules,omitempty"`
}

func (o UpdateRuleSetReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRuleSetReq struct{}"
	}

	return strings.Join([]string{"UpdateRuleSetReq", string(data)}, " ")
}
