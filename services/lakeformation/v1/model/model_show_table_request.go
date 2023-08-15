package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTableRequest Request Object
type ShowTableRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// catalog名字
	CatalogName string `json:"catalog_name"`

	// 数据库名字
	DatabaseName string `json:"database_name"`

	// 表名称
	TableName string `json:"table_name"`
}

func (o ShowTableRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTableRequest struct{}"
	}

	return strings.Join([]string{"ShowTableRequest", string(data)}, " ")
}
