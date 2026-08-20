package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StatsConfigDetails 配置详情
type StatsConfigDetails struct {

	// **参数解释：** 配置类型 **取值范围：** - 0：热点统计 - 1：ces上报
	ConfigType *int32 `json:"config_type,omitempty"`

	// **参数解释：** 资源类型 **取值范围：** - domain：域名，对应resource_name需配置为域名 - account：账号，对应resource_name需配置为账号
	ResourceType *string `json:"resource_type,omitempty"`

	// **参数解释：** 资源名称 > 账号或域名  **约束限制：** 不涉及 **取值范围：** 多个域名以英文逗号分隔
	ResourceName *string `json:"resource_name,omitempty"`

	ConfigInfo *ConfigInfo `json:"config_info,omitempty"`
}

func (o StatsConfigDetails) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StatsConfigDetails struct{}"
	}

	return strings.Join([]string{"StatsConfigDetails", string(data)}, " ")
}
