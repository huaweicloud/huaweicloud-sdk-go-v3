package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DataFilteringAdvancedSetting 数据过滤高级设置信息体
type DataFilteringAdvancedSetting struct {

	// 库名
	DbName string `json:"db_name"`

	// 表名
	TableName string `json:"table_name"`

	// 列名
	ColNames string `json:"col_names"`

	// 主键或唯一索引
	PrimKeyOrIndexes string `json:"prim_key_or_indexes"`

	// 优化查询所需的索引
	Indexes string `json:"indexes"`

	// 过滤条件
	Values string `json:"values"`
}

func (o DataFilteringAdvancedSetting) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataFilteringAdvancedSetting struct{}"
	}

	return strings.Join([]string{"DataFilteringAdvancedSetting", string(data)}, " ")
}
