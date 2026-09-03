package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TimeoutAlert struct {
	AlertTemplate *AlertTemplate `json:"alert_template,omitempty"`

	// 超时告警开启 0关闭 1开启
	Enable *string `json:"enable,omitempty"`

	TaskTimeoutPolicy *TaskTimeoutPolicy `json:"task_timeout_policy,omitempty"`

	TestCaseTimeoutPolicy *TestCaseTimeoutPolicy `json:"testCaseTimeoutPolicy,omitempty"`

	// 超时重试次数
	TimeoutRetryTimes *int32 `json:"timeoutRetryTimes,omitempty"`
}

func (o TimeoutAlert) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TimeoutAlert struct{}"
	}

	return strings.Join([]string{"TimeoutAlert", string(data)}, " ")
}
