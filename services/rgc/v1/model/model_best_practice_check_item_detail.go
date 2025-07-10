package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BestPracticeCheckItemDetail struct {

	// 检查项编号
	CheckItem *int32 `json:"check_item,omitempty"`

	// 检查项描述
	CheckItemName *string `json:"check_item_name,omitempty"`

	// 检查项风险描述
	RiskDescription *string `json:"risk_description,omitempty"`

	// 检测结果
	Result *string `json:"result,omitempty"`

	// 八大领域的细分场景
	Scene *string `json:"scene,omitempty"`

	// 风险等级
	RiskLevel *string `json:"risk_level,omitempty"`
}

func (o BestPracticeCheckItemDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BestPracticeCheckItemDetail struct{}"
	}

	return strings.Join([]string{"BestPracticeCheckItemDetail", string(data)}, " ")
}
