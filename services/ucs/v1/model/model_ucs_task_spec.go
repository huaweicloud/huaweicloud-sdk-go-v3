package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UcsTaskSpec struct {

	// 所属job的ID
	JobID *string `json:"jobID,omitempty"`

	// 执行目标
	Target *string `json:"target,omitempty"`

	// 执行目标类型
	TargetType *string `json:"targetType,omitempty"`

	// Task执行等待时间，单位：秒
	WaitTimeOut *int32 `json:"waitTimeOut,omitempty"`

	// Task类别
	Type *string `json:"type,omitempty"`

	// 执行方式：parallel和serial二选一
	Runway *string `json:"runway,omitempty"`
}

func (o UcsTaskSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UcsTaskSpec struct{}"
	}

	return strings.Join([]string{"UcsTaskSpec", string(data)}, " ")
}
