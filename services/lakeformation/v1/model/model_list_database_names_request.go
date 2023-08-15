package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDatabaseNamesRequest Request Object
type ListDatabaseNamesRequest struct {

	// 实例Id
	InstanceId string `json:"instance_id"`

	// catalog名字
	CatalogName string `json:"catalog_name"`

	// 数据库名字通配符
	DatabasePattern *string `json:"database_pattern,omitempty"`
}

func (o ListDatabaseNamesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDatabaseNamesRequest struct{}"
	}

	return strings.Join([]string{"ListDatabaseNamesRequest", string(data)}, " ")
}
