package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetInstancesNoIndexTablesRequest Request Object
type GetInstancesNoIndexTablesRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 指定查询是否应侧重于检索最新或最新的特殊表。
	Newest bool `json:"newest"`

	// 表格类型。
	TableType string `json:"table_type"`

	// 索引位置，偏移量。
	Offset *string `json:"offset,omitempty"`

	// 查询记录数。
	Limit *string `json:"limit,omitempty"`
}

func (o GetInstancesNoIndexTablesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetInstancesNoIndexTablesRequest struct{}"
	}

	return strings.Join([]string{"GetInstancesNoIndexTablesRequest", string(data)}, " ")
}
