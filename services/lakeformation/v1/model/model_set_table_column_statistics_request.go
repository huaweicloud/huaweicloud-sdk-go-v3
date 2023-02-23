package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type SetTableColumnStatisticsRequest struct {

	// 实例Id
	InstanceId string `json:"instance_id"`

	// catalog名字
	CatalogName string `json:"catalog_name"`

	// 数据库名字
	DatabaseName string `json:"database_name"`

	// 表名称
	TableName string `json:"table_name"`

	Body *MergeTableColumnStatisticInput `json:"body,omitempty"`
}

func (o SetTableColumnStatisticsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetTableColumnStatisticsRequest struct{}"
	}

	return strings.Join([]string{"SetTableColumnStatisticsRequest", string(data)}, " ")
}
