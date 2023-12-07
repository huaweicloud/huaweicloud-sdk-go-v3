package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AutoScalingPolicyInfo struct {

	// 当前自动伸缩规则是否开启。
	AutoScalingEnable bool `json:"auto_scaling_enable"`

	// 指定该节点组的最小保留节点数。 取值范围：[0～500]
	MinCapacity int32 `json:"min_capacity"`

	// 指定该节点组的最大节点数。 取值范围：[0～500]
	MaxCapacity int32 `json:"max_capacity"`

	// 资源计划列表。若该参数为空表示不启用资源计划。  当启用弹性伸缩时，资源计划与自动伸缩规则需至少配置其中一种。
	ResourcesPlans *[]ResourcesPlan `json:"resources_plans,omitempty"`

	// 自动伸缩的规则列表。 当启用弹性伸缩时，资源计划与自动伸缩规则需至少配置其中一种。
	Rules *[]Rule `json:"rules,omitempty"`

	// 弹性伸缩标签列表
	Tags *[]Tag `json:"tags,omitempty"`
}

func (o AutoScalingPolicyInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AutoScalingPolicyInfo struct{}"
	}

	return strings.Join([]string{"AutoScalingPolicyInfo", string(data)}, " ")
}
