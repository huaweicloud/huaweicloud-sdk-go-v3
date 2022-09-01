package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateGroupResponse struct {

	// 存储组 ID
	GroupId *string `json:"group_id,omitempty" xml:"group_id"`

	// 存储组名称
	Name *string `json:"name,omitempty" xml:"name"`

	// 描述
	Description *string `json:"description,omitempty" xml:"description"`

	// 温数据存储用量
	WarmDataUsage *int64 `json:"warm_data_usage,omitempty" xml:"warm_data_usage"`

	// 此存储组下存储实例的个数
	DataStoreCount *int64 `json:"data_store_count,omitempty" xml:"data_store_count"`

	// 冷数据存储用量
	ColdDataUsage *int64 `json:"cold_data_usage,omitempty" xml:"cold_data_usage"`

	// 温数据老化策略，单位只支持d（天），且只支持整数，如365天则可配置为“365d”，如“365h”或“360.5d”等均不被支持
	WarmDataRetentionPolicy *string `json:"warm_data_retention_policy,omitempty" xml:"warm_data_retention_policy"`

	// 冷数据老化策略，单位只支持d（天），且只支持整数，如365天则可配置为“365d”，如“365h”或“360.5d”等均不被支持
	ColdDataRetentionPolicy *string `json:"cold_data_retention_policy,omitempty" xml:"cold_data_retention_policy"`

	// 创建时间
	CreatedTime *string `json:"created_time,omitempty" xml:"created_time"`

	// 修改时间
	ModifiedTime *string `json:"modified_time,omitempty" xml:"modified_time"`

	// 存储类型，有资产存储(取值:AssetStorage)、设备存储(取值:DeviceStorage)两种类型
	Type           *string `json:"type,omitempty" xml:"type"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGroupResponse struct{}"
	}

	return strings.Join([]string{"CreateGroupResponse", string(data)}, " ")
}
