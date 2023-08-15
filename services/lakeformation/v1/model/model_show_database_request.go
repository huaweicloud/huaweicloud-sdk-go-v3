package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDatabaseRequest Request Object
type ShowDatabaseRequest struct {

	// 实例Id
	InstanceId string `json:"instance_id"`

	// catalog名字
	CatalogName string `json:"catalog_name"`

	// 数据库名字
	DatabaseName string `json:"database_name"`
}

func (o ShowDatabaseRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDatabaseRequest struct{}"
	}

	return strings.Join([]string{"ShowDatabaseRequest", string(data)}, " ")
}
