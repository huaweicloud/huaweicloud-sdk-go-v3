package model

import (
	"encoding/json"

	"strings"
)

type JobInstance struct {
	JobName     *string `json:"jobName,omitempty"`
	Status      *string `json:"status,omitempty"`
	PlanTime    *int32  `json:"planTime,omitempty"`
	StartTime   *int32  `json:"startTime,omitempty"`
	EndTime     *int32  `json:"endTime,omitempty"`
	ExecuteTime *int32  `json:"executeTime,omitempty"`
	InstancesId *string `json:"instancesId,omitempty"`
}

func (o JobInstance) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "JobInstance struct{}"
	}

	return strings.Join([]string{"JobInstance", string(data)}, " ")
}
