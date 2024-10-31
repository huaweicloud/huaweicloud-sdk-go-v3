package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HttpStatisticsItem struct {

	// 攻击类别
	AttackCategory *string `json:"attack_category,omitempty"`

	// 统计数量
	StatNum *int64 `json:"stat_num,omitempty"`
}

func (o HttpStatisticsItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HttpStatisticsItem struct{}"
	}

	return strings.Join([]string{"HttpStatisticsItem", string(data)}, " ")
}
