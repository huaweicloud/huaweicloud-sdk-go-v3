package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateGroupResponse struct {
	// 存储组 ID

	GroupId *string `json:"group_id,omitempty"`
	// 存储组名称

	Name *string `json:"name,omitempty"`
	// 描述

	Description *string `json:"description,omitempty"`
	// 温数据存储用量

	WarmDataUsage *int64 `json:"warm_data_usage,omitempty"`
	// 此存储组下存储实例的个数

	DataStoreCount *int64 `json:"data_store_count,omitempty"`
	// 冷数据存储用量

	ColdDataUsage *int64 `json:"cold_data_usage,omitempty"`
	// 温数据老化策略，单位只支持d（天），且只支持整数，如365天则可配置为“365d”，如“365h”或“360.5d”等均不被支持

	WarmDataRetentionPolicy *string `json:"warm_data_retention_policy,omitempty"`
	// 冷数据老化策略，单位只支持d（天），且只支持整数，如365天则可配置为“365d”，如“365h”或“360.5d”等均不被支持

	ColdDataRetentionPolicy *string `json:"cold_data_retention_policy,omitempty"`
	// 创建时间

	CreatedTime *string `json:"created_time,omitempty"`
	// 修改时间

	ModifiedTime *string `json:"modified_time,omitempty"`
	// 存储类型，有资产存储(取值:AssetStorage)、设备存储(取值:DeviceStorage)两种类型

	Type           *string `json:"type,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateGroupResponse struct{}"
	}

	return strings.Join([]string{"UpdateGroupResponse", string(data)}, " ")
}
