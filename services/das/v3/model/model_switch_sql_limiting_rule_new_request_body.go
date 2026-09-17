package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchSqlLimitingRuleNewRequestBody 切换SQL限流规则请求体
type SwitchSqlLimitingRuleNewRequestBody struct {

	// 数据库引擎类型
	EngineType string `json:"engine_type"`

	// SQL限流规则ID，可组合，用逗号分隔
	ItemIds string `json:"item_ids"`

	// SQL限流动作标志
	Action string `json:"action"`
}

func (o SwitchSqlLimitingRuleNewRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchSqlLimitingRuleNewRequestBody struct{}"
	}

	return strings.Join([]string{"SwitchSqlLimitingRuleNewRequestBody", string(data)}, " ")
}
