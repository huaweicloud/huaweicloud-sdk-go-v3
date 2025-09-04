package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetStatsConfigBody 设置统计配置请求体
type SetStatsConfigBody struct {

	// 配置类别.0：热点统计类
	ConfigType *int32 `json:"config_type,omitempty"`

	// 资源类型。domain:resource_name为域名，account:resource_name为账号
	ResourceType string `json:"resource_type"`

	// 资源名称为账号或域名。多个域名以英文逗号分隔
	ResourceName string `json:"resource_name"`

	// 配置信息.top指标仅支持ua、refer、url、origin url
	ConfigInfo *interface{} `json:"config_info"`

	// 统计配置失效时间，秒时间戳。不能超过当前时间点往后一年
	ExpiredTime *int64 `json:"expired_time,omitempty"`
}

func (o SetStatsConfigBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetStatsConfigBody struct{}"
	}

	return strings.Join([]string{"SetStatsConfigBody", string(data)}, " ")
}
