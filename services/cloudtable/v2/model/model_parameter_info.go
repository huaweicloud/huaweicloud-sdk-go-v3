package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ParameterInfo 集群配置信息详情
type ParameterInfo struct {

	// 参数ID
	Id *string `json:"id,omitempty"`

	// 配置名称
	Name string `json:"name"`

	// 默认值
	DefaultValue string `json:"default_value"`

	// 配置值类型
	ValueType string `json:"value_type"`

	// 集群当前运行的配置值
	RunningValue string `json:"running_value"`

	// 单位
	Unit string `json:"unit"`

	// 是否需要重启生效
	Reboot bool `json:"reboot"`

	// 配置值取值范围
	ValueRange string `json:"value_range"`

	// 配置描述信息
	Description string `json:"description"`
}

func (o ParameterInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ParameterInfo struct{}"
	}

	return strings.Join([]string{"ParameterInfo", string(data)}, " ")
}
