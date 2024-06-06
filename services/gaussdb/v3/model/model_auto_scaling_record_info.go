package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AutoScalingRecordInfo struct {

	// 记录ID。
	Id *string `json:"id,omitempty"`

	// 实例ID。
	InstanceId *string `json:"instance_id,omitempty"`

	// 实例名称。
	InstanceName *string `json:"instance_name,omitempty"`

	// 变配类型。
	ScalingType *string `json:"scaling_type,omitempty"`

	// 原始值。
	OriginalValue *string `json:"original_value,omitempty"`

	// 目标值。
	TargetValue *string `json:"target_value,omitempty"`

	// 变配结果。
	Result *string `json:"result,omitempty"`

	// 变配时间。
	CreateAt *string `json:"create_at,omitempty"`
}

func (o AutoScalingRecordInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AutoScalingRecordInfo struct{}"
	}

	return strings.Join([]string{"AutoScalingRecordInfo", string(data)}, " ")
}
