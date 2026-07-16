package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AffinityType struct {

	// 参数描述： 专属池场景下设置强亲和特性 参数约束： key、values、operator必填
	RequiredDuringSchedulingIgnoredDuringExecution *[]AffinityRule `json:"required_during_scheduling_ignored_during_execution,omitempty"`

	// 参数描述： 专属池场景下设置弱亲和特性 参数约束： key、values、operator必填，weight选填
	PreferredDuringSchedulingIgnoredDuringExecution *[]AffinityRule `json:"preferred_during_scheduling_ignored_during_execution,omitempty"`
}

func (o AffinityType) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AffinityType struct{}"
	}

	return strings.Join([]string{"AffinityType", string(data)}, " ")
}
