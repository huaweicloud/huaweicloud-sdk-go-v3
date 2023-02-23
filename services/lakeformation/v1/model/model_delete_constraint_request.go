package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteConstraintRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// catalog名字
	CatalogName string `json:"catalog_name"`

	// 数据库名字
	DatabaseName string `json:"database_name"`

	// 表名称
	TableName string `json:"table_name"`

	// 列限制信息名字; 外键填写外键限制名
	ConstraintName string `json:"constraint_name"`
}

func (o DeleteConstraintRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteConstraintRequest struct{}"
	}

	return strings.Join([]string{"DeleteConstraintRequest", string(data)}, " ")
}
