package model

import (
	"encoding/json"

	"strings"
)

type NodeInstance struct {
	NodeName       *string `json:"nodeName,omitempty"`
	Status         *string `json:"status,omitempty"`
	PlanTime       *int32  `json:"planTime,omitempty"`
	StartTime      *int32  `json:"startTime,omitempty"`
	EndTime        *int32  `json:"endTime,omitempty"`
	ExecuteTime    *int32  `json:"executeTime,omitempty"`
	NodeType       *string `json:"nodeType,omitempty"`
	RetryTimes     *int32  `json:"retryTimes,omitempty"`
	InstanceId     *int32  `json:"instanceId,omitempty"`
	InputRowCount  *int32  `json:"inputRowCount,omitempty"`
	OutputRowCount *int32  `json:"outputRowCount,omitempty"`
	LogPath        *string `json:"logPath,omitempty"`
}

func (o NodeInstance) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "NodeInstance struct{}"
	}

	return strings.Join([]string{"NodeInstance", string(data)}, " ")
}
