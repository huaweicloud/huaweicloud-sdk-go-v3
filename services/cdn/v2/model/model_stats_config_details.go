package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StatsConfigDetails 配置详情
type StatsConfigDetails struct {

	// 配置类别.0：热点统计类
	ConfigType *int32 `json:"config_type,omitempty"`

	// 资源类型。domain:resource_name为域名，account:resource_name为账号
	ResourceType *string `json:"resource_type,omitempty"`

	// 资源名称为账号或域名。多个域名以英文逗号分隔
	ResourceName *string `json:"resource_name,omitempty"`

	ConfigInfo *ConfigInfo `json:"config_info,omitempty"`

	// 统计配置失效时间，秒时间戳
	ExpiredTime *int64 `json:"expired_time,omitempty"`
}

func (o StatsConfigDetails) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StatsConfigDetails struct{}"
	}

	return strings.Join([]string{"StatsConfigDetails", string(data)}, " ")
}
