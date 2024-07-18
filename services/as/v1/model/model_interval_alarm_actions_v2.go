package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IntervalAlarmActionsV2 操作告警区间
type IntervalAlarmActionsV2 struct {

	// 操作选项，默认为ADD。 当scaling_resource_type为SCALING_GROUP，支持如下操作： ADD：增加 REMOVE/REDUCE：减少 SET：设置为 当scaling_resource_type为BANDWIDTH，支持如下操作： ADD：增加 REDUCE：减少
	Operation *string `json:"operation,omitempty"`

	// 操作限制。当scaling_resource_type为BANDWIDTH，且operation不为SET时，limits参数生效，单位为Mbit/s。此时，当operation为ADD时，limits表示带宽可调整的上限；当operation为REDUCE时，limits表示带宽可调整的下限。
	Limits *int32 `json:"limits,omitempty"`

	// 操作大小，取值范围为0到300的整数，默认为1。当scaling_resource_type为SCALING_GROUP时，size为实例个数,取值范围为0-300的整数，默认为1。当scaling_resource_type为BANDWIDTH时，size表示带宽大小，单位为Mbit/s，取值范围为1到300的整数，默认为1。当scaling_resource_type为SCALING_GROUP时，size和percentage参数只能选其中一个进行配置。
	Size *int32 `json:"size,omitempty"`

	LowerBound *float64 `json:"lower_bound,omitempty"`

	UpperBound *float64 `json:"upper_bound,omitempty"`

	// 操作百分比，取值为0到20000的整数。当scaling_resource_type为SCALING_GROUP时，size和instance_percentage参数均无配置，则size默认为1。当scaling_resource_type为BANDWIDTH时，不支持配置instance_percentage参数。
	Percentage *int32 `json:"percentage,omitempty"`
}

func (o IntervalAlarmActionsV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IntervalAlarmActionsV2 struct{}"
	}

	return strings.Join([]string{"IntervalAlarmActionsV2", string(data)}, " ")
}
