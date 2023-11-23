package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteKillOpRuleListRequestBody struct {

	// killOp规则ID列表。
	Ids []string `json:"ids"`
}

func (o DeleteKillOpRuleListRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteKillOpRuleListRequestBody struct{}"
	}

	return strings.Join([]string{"DeleteKillOpRuleListRequestBody", string(data)}, " ")
}
