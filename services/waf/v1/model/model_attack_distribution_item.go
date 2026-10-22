package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AttackDistributionItem 攻击类型分布项
type AttackDistributionItem struct {

	// 统计指标类型： - compliance_total_matched：内容合规检测命中次数 - injection_total_matched：提示词注入攻击检测命中次数
	Key string `json:"key"`

	// 该类型总命中次数
	Num int64 `json:"num"`

	// 子类型分布列表
	Items []AttackSubItem `json:"items"`
}

func (o AttackDistributionItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttackDistributionItem struct{}"
	}

	return strings.Join([]string{"AttackDistributionItem", string(data)}, " ")
}
