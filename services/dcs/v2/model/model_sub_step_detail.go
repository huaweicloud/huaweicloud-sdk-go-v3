package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SubStepDetail 子任务详情
type SubStepDetail struct {

	// 任务id
	SubStepId *string `json:"sub_step_id,omitempty"`

	// 任务名
	SubStepName *string `json:"sub_step_name,omitempty"`

	// 任务状态
	SubStepStatus *string `json:"sub_step_status,omitempty"`

	// 任务启动时间，格式为2020-06-17T07:38:42.503Z
	BeginTime *string `json:"begin_time,omitempty"`

	// 任务结束时间，格式为2020-06-17T07:38:42.503Z
	EndTime *string `json:"end_time,omitempty"`

	// 子任务的附加属性详情
	Detail *string `json:"detail,omitempty"`

	// 错误码
	ErrorCode *string `json:"error_code,omitempty"`
}

func (o SubStepDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubStepDetail struct{}"
	}

	return strings.Join([]string{"SubStepDetail", string(data)}, " ")
}
