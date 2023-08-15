package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTableRequest Request Object
type DeleteTableRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// catalog名字
	CatalogName string `json:"catalog_name"`

	// 数据库名字
	DatabaseName string `json:"database_name"`

	// 表名称
	TableName string `json:"table_name"`

	// 是否删除表中的数据
	DeleteData *bool `json:"delete_data,omitempty"`
}

func (o DeleteTableRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTableRequest struct{}"
	}

	return strings.Join([]string{"DeleteTableRequest", string(data)}, " ")
}
