package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NodeInstance struct {
	NodeName *string `json:"nodeName,omitempty" xml:"nodeName"`

	Status *string `json:"status,omitempty" xml:"status"`

	PlanTime *int32 `json:"planTime,omitempty" xml:"planTime"`

	StartTime *int32 `json:"startTime,omitempty" xml:"startTime"`

	EndTime *int32 `json:"endTime,omitempty" xml:"endTime"`

	ExecuteTime *int32 `json:"executeTime,omitempty" xml:"executeTime"`

	NodeType *string `json:"nodeType,omitempty" xml:"nodeType"`

	RetryTimes *int32 `json:"retryTimes,omitempty" xml:"retryTimes"`

	InstanceId *int32 `json:"instanceId,omitempty" xml:"instanceId"`

	InputRowCount *int32 `json:"inputRowCount,omitempty" xml:"inputRowCount"`

	OutputRowCount *int32 `json:"outputRowCount,omitempty" xml:"outputRowCount"`

	LogPath *string `json:"logPath,omitempty" xml:"logPath"`
}

func (o NodeInstance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeInstance struct{}"
	}

	return strings.Join([]string{"NodeInstance", string(data)}, " ")
}
