package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstancesNoIndexTablesRequest Request Object
type ListInstancesNoIndexTablesRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 指定查询是否应侧重于检索最新或最新的特殊表。
	Newest bool `json:"newest"`

	// 表格类型。
	TableType string `json:"table_type"`
}

func (o ListInstancesNoIndexTablesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstancesNoIndexTablesRequest struct{}"
	}

	return strings.Join([]string{"ListInstancesNoIndexTablesRequest", string(data)}, " ")
}
