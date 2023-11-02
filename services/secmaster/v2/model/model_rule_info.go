package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RuleInfo 剧本触发规格信息
type RuleInfo struct {

	// 规则ID
	Id *string `json:"id,omitempty"`

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// 触发规则
	Rule *string `json:"rule,omitempty"`
}

func (o RuleInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RuleInfo struct{}"
	}

	return strings.Join([]string{"RuleInfo", string(data)}, " ")
}
