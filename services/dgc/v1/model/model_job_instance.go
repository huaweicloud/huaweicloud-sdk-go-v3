package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type JobInstance struct {
	JobName *string `json:"jobName,omitempty" xml:"jobName"`

	Status *string `json:"status,omitempty" xml:"status"`

	PlanTime *int32 `json:"planTime,omitempty" xml:"planTime"`

	StartTime *int32 `json:"startTime,omitempty" xml:"startTime"`

	EndTime *int32 `json:"endTime,omitempty" xml:"endTime"`

	ExecuteTime *int32 `json:"executeTime,omitempty" xml:"executeTime"`

	InstancesId *string `json:"instancesId,omitempty" xml:"instancesId"`
}

func (o JobInstance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobInstance struct{}"
	}

	return strings.Join([]string{"JobInstance", string(data)}, " ")
}
