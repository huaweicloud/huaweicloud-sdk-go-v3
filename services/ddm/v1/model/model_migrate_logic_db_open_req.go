package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MigrateLogicDbOpenReq struct {

	// 关联的后端DN信息。
	DataNodes *[]DataNode `json:"data_nodes,omitempty"`

	// 路由切换开始时间。
	SwitchRouteBeginTime *string `json:"switch_route_begin_time,omitempty"`

	// 路由切换结束时间。
	SwitchRouteEndTime *string `json:"switch_route_end_time,omitempty"`

	// 新分片数。
	NewShardNumber *string `json:"new_shard_number,omitempty"`

	// 是否独占。
	IsExclusive *bool `json:"is_exclusive,omitempty"`
}

func (o MigrateLogicDbOpenReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MigrateLogicDbOpenReq struct{}"
	}

	return strings.Join([]string{"MigrateLogicDbOpenReq", string(data)}, " ")
}
