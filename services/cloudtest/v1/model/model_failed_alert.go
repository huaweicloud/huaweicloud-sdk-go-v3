package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FailedAlert struct {
	CloudAlarm *CloudAlarmDto `json:"cloudAlarm,omitempty"`

	TaskPolicy *TaskPolicy `json:"taskPolicy,omitempty"`

	TestCasePolicy *TestCasePolicy `json:"testCasePolicy,omitempty"`

	WiseEye *WiseEye `json:"wiseEye,omitempty"`
}

func (o FailedAlert) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FailedAlert struct{}"
	}

	return strings.Join([]string{"FailedAlert", string(data)}, " ")
}
