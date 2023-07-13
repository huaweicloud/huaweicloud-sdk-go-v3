package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTableColumnStatisticsRequest Request Object
type ListTableColumnStatisticsRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// catalog名字
	CatalogName string `json:"catalog_name"`

	// 数据库名字
	DatabaseName string `json:"database_name"`

	// 表名称
	TableName string `json:"table_name"`

	Body *GetTableColumnStatisticsInput `json:"body,omitempty"`
}

func (o ListTableColumnStatisticsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTableColumnStatisticsRequest struct{}"
	}

	return strings.Join([]string{"ListTableColumnStatisticsRequest", string(data)}, " ")
}
