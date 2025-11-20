package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MigrateResultsResponse Response Object
type MigrateResultsResponse struct {

	// 实例名称。
	InstanceName *string `json:"instance_name,omitempty"`

	// 原始分片数。
	OriginalShardNum *int32 `json:"original_shard_num,omitempty"`

	// 现分片数。
	AfterShardNum *int32 `json:"after_shard_num,omitempty"`

	// 分片变更前关联的后端DN信息。
	OriginalDnInfoList *[]MigrateDnInfoOpenResponse `json:"original_dn_info_list,omitempty"`

	// 分片变更后关联的后端DN信息。
	AfterDnInfoList *[]MigrateDnInfoOpenResponse `json:"after_dn_info_list,omitempty"`

	// 是否手动切换路由。
	SwitchRouteIsManual *bool `json:"switch_route_is_manual,omitempty"`

	// 自动切换路由开始时间。
	AutoSwitchBeginTime *string `json:"auto_switch_begin_time,omitempty"`

	// 自动切换路由结束时间。
	AutoSwitchEndTime *string `json:"auto_switch_end_time,omitempty"`

	// 分片变更的节点。
	NodeIdForMigrate *string `json:"node_id_for_migrate,omitempty"`

	// 是否独占式。
	IsExclusive    *bool `json:"is_exclusive,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o MigrateResultsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MigrateResultsResponse struct{}"
	}

	return strings.Join([]string{"MigrateResultsResponse", string(data)}, " ")
}
