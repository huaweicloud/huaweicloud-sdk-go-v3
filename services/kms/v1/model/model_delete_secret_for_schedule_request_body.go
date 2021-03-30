package model

import (
	"encoding/json"

	"strings"
)

type DeleteSecretForScheduleRequestBody struct {
	// 创建定时删除凭据的任务，且指定可恢复的天数。  约束：7~30。  默认值：30。

	RecoveryWindowInDays int32 `json:"recovery_window_in_days"`
}

func (o DeleteSecretForScheduleRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteSecretForScheduleRequestBody struct{}"
	}

	return strings.Join([]string{"DeleteSecretForScheduleRequestBody", string(data)}, " ")
}
