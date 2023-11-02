package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateRuleInfo 创建剧本触发规则参数
type CreateRuleInfo struct {
	Rule *ConditionInfo `json:"rule"`
}

func (o CreateRuleInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRuleInfo struct{}"
	}

	return strings.Join([]string{"CreateRuleInfo", string(data)}, " ")
}
