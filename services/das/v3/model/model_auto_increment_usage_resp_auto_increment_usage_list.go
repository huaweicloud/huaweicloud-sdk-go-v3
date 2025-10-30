package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AutoIncrementUsageRespAutoIncrementUsageList struct {

	// 数据库名
	Database string `json:"database"`

	// 表名
	Table string `json:"table"`

	// 列名
	Column string `json:"column"`

	// 自增ID当前值
	CurrentValue string `json:"current_value"`

	// 该数据类型的自增 ID 支持的最大值。
	MaxValue string `json:"max_value"`

	// 自增ID使用比例
	Ratio string `json:"ratio"`
}

func (o AutoIncrementUsageRespAutoIncrementUsageList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AutoIncrementUsageRespAutoIncrementUsageList struct{}"
	}

	return strings.Join([]string{"AutoIncrementUsageRespAutoIncrementUsageList", string(data)}, " ")
}
