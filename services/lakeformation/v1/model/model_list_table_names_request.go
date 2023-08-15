package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTableNamesRequest Request Object
type ListTableNamesRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// catalog名字
	CatalogName string `json:"catalog_name"`

	// 数据库名字
	DatabaseName string `json:"database_name"`

	// 表名通配符
	TablePattern *string `json:"table_pattern,omitempty"`

	// 查询的表类型
	TableType *string `json:"table_type,omitempty"`
}

func (o ListTableNamesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTableNamesRequest struct{}"
	}

	return strings.Join([]string{"ListTableNamesRequest", string(data)}, " ")
}
