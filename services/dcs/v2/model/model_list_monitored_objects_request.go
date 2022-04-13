package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListMonitoredObjectsRequest struct {
	// 主维度ID，当前支持dcs_instance_id，dcs_memcached_instance_id。

	DimName string `json:"dim_name"`
	// 偏移量，表示从此偏移量开始查询，offset大于等于0

	Offset *int32 `json:"offset,omitempty"`
	// 每页显示的条目数量

	Limit *int32 `json:"limit,omitempty"`
}

func (o ListMonitoredObjectsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMonitoredObjectsRequest struct{}"
	}

	return strings.Join([]string{"ListMonitoredObjectsRequest", string(data)}, " ")
}
