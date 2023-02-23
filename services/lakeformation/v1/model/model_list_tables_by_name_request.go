package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListTablesByNameRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// catalog名字
	CatalogName string `json:"catalog_name"`

	// 数据库名字
	DatabaseName string `json:"database_name"`

	Body *ListTableByNameInput `json:"body,omitempty"`
}

func (o ListTablesByNameRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTablesByNameRequest struct{}"
	}

	return strings.Join([]string{"ListTablesByNameRequest", string(data)}, " ")
}
