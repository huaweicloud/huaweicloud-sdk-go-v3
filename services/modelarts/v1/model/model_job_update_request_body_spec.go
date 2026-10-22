package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// JobUpdateRequestBodySpec **参数解释**：训练作业规格参数。 **约束限制**：不涉及。
type JobUpdateRequestBodySpec struct {
	Notification *Notification `json:"notification,omitempty"`

	SchedulePolicy *JobUpdateRequestBodySpecSchedulePolicy `json:"schedule_policy,omitempty"`
}

func (o JobUpdateRequestBodySpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobUpdateRequestBodySpec struct{}"
	}

	return strings.Join([]string{"JobUpdateRequestBodySpec", string(data)}, " ")
}
