package model

import (
	"encoding/json"

	"strings"
)

type PlanRecordTime struct {
	// 录制开始时间，格式：yyyy-mm-ddThh:mm:ssZ，UTC时间。

	StartTime string `json:"start_time"`
	// 录制结束时间，格式：yyyy-mm-ddThh:mm:ssZ，UTC时间。如果填写，填写的时间必须晚于当前时间。如果不填写，则在计划录制触发后不停止。

	EndTime *string `json:"end_time,omitempty"`
}

func (o PlanRecordTime) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "PlanRecordTime struct{}"
	}

	return strings.Join([]string{"PlanRecordTime", string(data)}, " ")
}
