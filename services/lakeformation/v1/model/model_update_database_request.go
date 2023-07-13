package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDatabaseRequest Request Object
type UpdateDatabaseRequest struct {

	// 实例Id
	InstanceId string `json:"instance_id"`

	// catalog名字
	CatalogName string `json:"catalog_name"`

	// 数据库名字
	DatabaseName string `json:"database_name"`

	Body *DatabaseInput `json:"body,omitempty"`
}

func (o UpdateDatabaseRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDatabaseRequest struct{}"
	}

	return strings.Join([]string{"UpdateDatabaseRequest", string(data)}, " ")
}
