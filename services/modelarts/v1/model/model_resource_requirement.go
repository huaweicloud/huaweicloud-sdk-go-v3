package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResourceRequirement 算法资源约束。可不设置。设置后，在算法使用于训练作业时，控制台会过滤可用的公共资源池。
type ResourceRequirement struct {

	// 资源约束，可选键值如： - flavor_type（资源类型），对应值可选择CPU、GPU[或Ascend](tag:hc,hk,fcs_super)； - device_distributed_mode（是否支持多卡训练），对应值可选择multiple（支持）、singular（不支持）； - host_distributed_mode（是否支持分布式训练），对应值可选择multiple（支持）、singular（不支持）。
	Key *string `json:"key,omitempty"`

	// 资源约束键对应值。
	Values *[]string `json:"values,omitempty"`

	// 键与值关系，当前只支持in。例如flavor_type in [CPU,GPU]。
	Operator *string `json:"operator,omitempty"`
}

func (o ResourceRequirement) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceRequirement struct{}"
	}

	return strings.Join([]string{"ResourceRequirement", string(data)}, " ")
}
