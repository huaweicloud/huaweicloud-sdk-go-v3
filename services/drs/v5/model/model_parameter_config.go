package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ParameterConfig 任务参数配置信息。
type ParameterConfig struct {

	// 参数名称
	Name string `json:"name"`

	// 参数值
	Value string `json:"value"`

	// 参数默认值
	DefaultValue string `json:"default_value"`

	// 参数值范围，如Integer取值0-1、Boolean取值true|false等。
	ValueRange string `json:"value_range"`

	// 是否需要重启。默认为true, “false”表示否。“true”表示是。
	IsNeedRestart bool `json:"is_need_restart"`

	// 参数描述。
	Description string `json:"description"`

	// 创建时间，例如：2023-01-20T07:18:26Z
	CreatedAt *string `json:"created_at,omitempty"`

	// 更新时间，例如：2023-03-01T09:42:02Z
	UpdatedAt *string `json:"updated_at,omitempty"`
}

func (o ParameterConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ParameterConfig struct{}"
	}

	return strings.Join([]string{"ParameterConfig", string(data)}, " ")
}
