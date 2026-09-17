package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RapidGrowthTableInfo RapidGrowthTableInfo对象
type RapidGrowthTableInfo struct {

	// 库名
	DbName *string `json:"db_name,omitempty"`

	// 表名
	TableName *string `json:"table_name,omitempty"`

	// 增长量，单位MB
	Growth *float64 `json:"growth,omitempty"`
}

func (o RapidGrowthTableInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RapidGrowthTableInfo struct{}"
	}

	return strings.Join([]string{"RapidGrowthTableInfo", string(data)}, " ")
}
