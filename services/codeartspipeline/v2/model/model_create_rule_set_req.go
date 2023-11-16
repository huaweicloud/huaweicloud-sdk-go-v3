package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateRuleSetReq struct {

	// 规则集名称
	Name string `json:"name"`

	// 规则集类型
	Type string `json:"type"`

	// 规则集合
	Rules []RequestRuleInstance `json:"rules"`
}

func (o CreateRuleSetReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRuleSetReq struct{}"
	}

	return strings.Join([]string{"CreateRuleSetReq", string(data)}, " ")
}
