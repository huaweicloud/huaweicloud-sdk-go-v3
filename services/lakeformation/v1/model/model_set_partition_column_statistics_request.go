package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetPartitionColumnStatisticsRequest Request Object
type SetPartitionColumnStatisticsRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// catalog名字
	CatalogName string `json:"catalog_name"`

	// 数据库名字
	DatabaseName string `json:"database_name"`

	// 表名称
	TableName string `json:"table_name"`

	Body *SetPartitionColumnStatisticsInput `json:"body,omitempty"`
}

func (o SetPartitionColumnStatisticsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetPartitionColumnStatisticsRequest struct{}"
	}

	return strings.Join([]string{"SetPartitionColumnStatisticsRequest", string(data)}, " ")
}
