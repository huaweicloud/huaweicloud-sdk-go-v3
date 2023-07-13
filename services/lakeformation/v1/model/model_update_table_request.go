package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTableRequest Request Object
type UpdateTableRequest struct {

	// 实例Id
	InstanceId string `json:"instance_id"`

	// catalog名字
	CatalogName string `json:"catalog_name"`

	// 数据库名字
	DatabaseName string `json:"database_name"`

	// 表名称
	TableName string `json:"table_name"`

	Body *AlterTableInput `json:"body,omitempty"`
}

func (o UpdateTableRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTableRequest struct{}"
	}

	return strings.Join([]string{"UpdateTableRequest", string(data)}, " ")
}
