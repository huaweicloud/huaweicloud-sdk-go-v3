package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchListPartitionByValuesRequest Request Object
type BatchListPartitionByValuesRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// catalog名字
	CatalogName string `json:"catalog_name"`

	// 数据库名字
	DatabaseName string `json:"database_name"`

	// 表名称
	TableName string `json:"table_name"`

	Body *GetPartitionsByValuesInput `json:"body,omitempty"`
}

func (o BatchListPartitionByValuesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchListPartitionByValuesRequest struct{}"
	}

	return strings.Join([]string{"BatchListPartitionByValuesRequest", string(data)}, " ")
}
