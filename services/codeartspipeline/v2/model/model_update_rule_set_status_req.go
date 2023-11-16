package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateRuleSetStatusReq struct {

	// 规则模版实例状态
	IsValid bool `json:"is_valid"`
}

func (o UpdateRuleSetStatusReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRuleSetStatusReq struct{}"
	}

	return strings.Join([]string{"UpdateRuleSetStatusReq", string(data)}, " ")
}
