package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchShowPartitionColumnStatisticsRequest Request Object
type BatchShowPartitionColumnStatisticsRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// catalog名字
	CatalogName string `json:"catalog_name"`

	// 数据库名字
	DatabaseName string `json:"database_name"`

	// 表名称
	TableName string `json:"table_name"`

	Body *GetPartitionColumnStatisticsInput `json:"body,omitempty"`
}

func (o BatchShowPartitionColumnStatisticsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchShowPartitionColumnStatisticsRequest struct{}"
	}

	return strings.Join([]string{"BatchShowPartitionColumnStatisticsRequest", string(data)}, " ")
}
