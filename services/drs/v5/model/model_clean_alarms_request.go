package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CleanAlarmsRequest Request Object
type CleanAlarmsRequest struct {

	// 请求语言类型。
	XLanguage *string `json:"X-Language,omitempty"`

	// 任务ID。
	JobId string `json:"job_id"`

	// DDL序列号。
	AlarmId string `json:"alarm_id"`
}

func (o CleanAlarmsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CleanAlarmsRequest struct{}"
	}

	return strings.Join([]string{"CleanAlarmsRequest", string(data)}, " ")
}
