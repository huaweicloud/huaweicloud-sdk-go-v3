package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 存储组信息
type StorageGroup struct {

	// 存储组名称
	Name *string `json:"name,omitempty" xml:"name"`

	// 描述
	Description *string `json:"description,omitempty" xml:"description"`

	// 温数据老化策略，单位只支持d（天），且只支持整数，如365天则可配置为“365d”，如“365h”或“360.5d”等均不被支持
	WarmDataRetentionPolicy *string `json:"warm_data_retention_policy,omitempty" xml:"warm_data_retention_policy"`

	// 冷数据老化策略，单位只支持d（天），且只支持整数，如365天则可配置为“365d”，如“365h”或“360.5d”等均不被支持
	ColdDataRetentionPolicy *string `json:"cold_data_retention_policy,omitempty" xml:"cold_data_retention_policy"`
}

func (o StorageGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StorageGroup struct{}"
	}

	return strings.Join([]string{"StorageGroup", string(data)}, " ")
}
