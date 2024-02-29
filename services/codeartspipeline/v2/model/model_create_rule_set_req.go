package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateRuleSetReq struct {

	// 策略名称
	Name string `json:"name"`

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
