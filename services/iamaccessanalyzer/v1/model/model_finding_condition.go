package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FindingCondition 分析的策略语句中导致查找结果的条件。
type FindingCondition struct {

	// 标签\"键\"的标识符或名称。
	Key string `json:"key"`

	// 标签\"键\"对应的\"值\"。
	Value string `json:"value"`
}

func (o FindingCondition) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FindingCondition struct{}"
	}

	return strings.Join([]string{"FindingCondition", string(data)}, " ")
}
