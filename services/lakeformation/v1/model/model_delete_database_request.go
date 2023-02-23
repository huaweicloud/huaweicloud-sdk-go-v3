package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteDatabaseRequest struct {

	// 实例Id
	InstanceId string `json:"instance_id"`

	// catalog名字
	CatalogName string `json:"catalog_name"`

	// 数据库名
	DatabaseName string `json:"database_name"`

	// 是否删除数据库路径下的数据
	DeleteData *bool `json:"delete_data,omitempty"`

	// 是否级联删除数据库下的表、分区以及函数
	Cascade *bool `json:"cascade,omitempty"`
}

func (o DeleteDatabaseRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDatabaseRequest struct{}"
	}

	return strings.Join([]string{"DeleteDatabaseRequest", string(data)}, " ")
}
