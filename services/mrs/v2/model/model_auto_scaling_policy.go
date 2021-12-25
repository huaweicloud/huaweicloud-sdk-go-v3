package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AutoScalingPolicy struct {
	// 当前自动伸缩规则是否开启。

	AutoScalingEnable bool `json:"auto_scaling_enable"`
	// 指定该节点组的最小保留节点数。  取值范围：[0～500]

	MinCapacity int32 `json:"min_capacity"`
	// 指定该节点组的最大节点数。  取值范围：[0～500]

	MaxCapacity int32 `json:"max_capacity"`
	// 资源计划列表。若该参数为空表示不启用资源计划。  当启用弹性伸缩时，资源计划与自动伸缩规则需至少配置其中一种。

	ResourcesPlans *[]ResourcesPlan `json:"resources_plans,omitempty"`
	// 自动伸缩的规则列表。  当启用弹性伸缩时，资源计划与自动伸缩规则需至少配置其中一种。

	Rules *[]Rules `json:"rules,omitempty"`
	// 弹性伸缩自定义自动化脚本列表。若该参数为空表示不启用自动化脚本。

	ExecScripts *[]ScaleScript `json:"exec_scripts,omitempty"`
}

func (o AutoScalingPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AutoScalingPolicy struct{}"
	}

	return strings.Join([]string{"AutoScalingPolicy", string(data)}, " ")
}
