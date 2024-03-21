package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowJob struct {

	// Job的ID
	Id *string `json:"id,omitempty"`

	// 处理规则
	Action *string `json:"action,omitempty"`

	// 状态
	Status *string `json:"status,omitempty"`

	// error_task
	ErrorTask *string `json:"error_task,omitempty"`

	// error_code
	ErrorCode *string `json:"error_code,omitempty"`

	// error_message
	ErrorMessage *string `json:"error_message,omitempty"`

	// 起始时间
	StartTime *sdktime.SdkTime `json:"start_time,omitempty"`

	// 结束时间
	EndTime *sdktime.SdkTime `json:"end_time,omitempty"`
}

func (o ShowJob) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJob struct{}"
	}

	return strings.Join([]string{"ShowJob", string(data)}, " ")
}
