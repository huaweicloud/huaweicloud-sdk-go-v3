package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PropagationSpec struct {

	// 资源选择器，用于选择要传播的资源
	ResourceSelectors *[]ResourceSelector `json:"resourceSelectors,omitempty"`

	// 是否自动传播引用的资源
	PropagateDeps *bool `json:"propagateDeps,omitempty"`

	Placement *Placement `json:"placement,omitempty"`

	// 策略的优先级，默认值为0
	Priority *int32 `json:"priority,omitempty"`

	// 调度器名称，默认值为“default-scheduler”
	SchedulerName *string `json:"schedulerName,omitempty"`
}

func (o PropagationSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PropagationSpec struct{}"
	}

	return strings.Join([]string{"PropagationSpec", string(data)}, " ")
}
