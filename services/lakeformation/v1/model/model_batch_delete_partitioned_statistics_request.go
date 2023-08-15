package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeletePartitionedStatisticsRequest Request Object
type BatchDeletePartitionedStatisticsRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// catalog名字
	CatalogName string `json:"catalog_name"`

	// 数据库名字
	DatabaseName string `json:"database_name"`

	// 表名称
	TableName string `json:"table_name"`

	Body *TruncatePartitionInput `json:"body,omitempty"`
}

func (o BatchDeletePartitionedStatisticsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeletePartitionedStatisticsRequest struct{}"
	}

	return strings.Join([]string{"BatchDeletePartitionedStatisticsRequest", string(data)}, " ")
}
