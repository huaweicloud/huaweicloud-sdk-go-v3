package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateKillOpRuleRequestBody struct {

	// killOp规则ID列表。
	Ids []string `json:"ids"`

	// 启用/禁用 killOp规则。  - enable，启用killOp规则。 - disable，禁用killOp规则。
	Action string `json:"action"`
}

func (o UpdateKillOpRuleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateKillOpRuleRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateKillOpRuleRequestBody", string(data)}, " ")
}
