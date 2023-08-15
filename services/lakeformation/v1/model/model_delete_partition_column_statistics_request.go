package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePartitionColumnStatisticsRequest Request Object
type DeletePartitionColumnStatisticsRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// catalog名字
	CatalogName string `json:"catalog_name"`

	// 数据库名字
	DatabaseName string `json:"database_name"`

	// 表名称
	TableName string `json:"table_name"`

	// 分区的值列表
	PartitionValues []string `json:"partition_values"`

	// 列名
	ColumnName *string `json:"column_name,omitempty"`
}

func (o DeletePartitionColumnStatisticsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePartitionColumnStatisticsRequest struct{}"
	}

	return strings.Join([]string{"DeletePartitionColumnStatisticsRequest", string(data)}, " ")
}
