package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTableColumnStatisticsRequest Request Object
type DeleteTableColumnStatisticsRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// catalog名字
	CatalogName string `json:"catalog_name"`

	// 数据库名字
	DatabaseName string `json:"database_name"`

	// 表名称
	TableName string `json:"table_name"`

	// 被删除的列名，不填写则删除表所有统计信息
	ColumnName *string `json:"column_name,omitempty"`
}

func (o DeleteTableColumnStatisticsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTableColumnStatisticsRequest struct{}"
	}

	return strings.Join([]string{"DeleteTableColumnStatisticsRequest", string(data)}, " ")
}
