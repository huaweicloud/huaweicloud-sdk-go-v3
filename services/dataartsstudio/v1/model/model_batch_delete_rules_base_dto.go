package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchDeleteRulesBaseDto struct {

	// 识别规则id列表
	RuleIds []string `json:"rule_ids"`
}

func (o BatchDeleteRulesBaseDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteRulesBaseDto struct{}"
	}

	return strings.Join([]string{"BatchDeleteRulesBaseDto", string(data)}, " ")
}
