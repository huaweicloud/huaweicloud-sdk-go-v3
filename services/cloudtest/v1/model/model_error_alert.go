package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ErrorAlert struct {
	AlertTemplate *AlertTemplate `json:"alert_template,omitempty"`

	// 异常告警是否开启：0关闭，1开启，默认关闭
	Enable *string `json:"enable,omitempty"`

	TaskErrorPolicy *TaskErrorPolicy `json:"taskErrorPolicy,omitempty"`
}

func (o ErrorAlert) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ErrorAlert struct{}"
	}

	return strings.Join([]string{"ErrorAlert", string(data)}, " ")
}
