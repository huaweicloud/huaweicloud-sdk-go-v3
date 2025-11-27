package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResourceBindingStatus struct {

	// 记录调度器当前观察到的资源生成版本代数
	SchedulerObservedGeneration *int32 `json:"schedulerObservedGeneration,omitempty"`

	// 资源的各种状态条件
	Conditions *[]ConditionStatus `json:"conditions,omitempty"`
}

func (o ResourceBindingStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceBindingStatus struct{}"
	}

	return strings.Join([]string{"ResourceBindingStatus", string(data)}, " ")
}
