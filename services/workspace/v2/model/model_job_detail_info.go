package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type JobDetailInfo struct {

	// 任务id。
	Id *string `json:"id,omitempty"`

	// 任务类型。
	JobType *string `json:"job_type,omitempty"`

	Entities *JobEntities `json:"entities,omitempty"`

	// 任务创建时间。
	BeginTime *string `json:"begin_time,omitempty"`

	// 任务结束时间。
	EndTime *string `json:"end_time,omitempty"`

	// 任务状态。
	Status *string `json:"status,omitempty"`

	// 任务执行失败时的错误码。
	ErrorCode *string `json:"error_code,omitempty"`

	// 任务失败原因。
	FailReason *string `json:"fail_reason,omitempty"`

	// 任务失败原因信息。
	Message *string `json:"message,omitempty"`

	// 任务ID。
	JobId *string `json:"job_id,omitempty"`

	// 桌面名称。
	DesktopName *string `json:"desktop_name,omitempty"`

	// ip地址。
	IpAddress *string `json:"ip_address,omitempty"`

	// mac地址。
	MacAddress *string `json:"mac_address,omitempty"`

	// 任务进度。
	Process *int32 `json:"process,omitempty"`
}

func (o JobDetailInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobDetailInfo struct{}"
	}

	return strings.Join([]string{"JobDetailInfo", string(data)}, " ")
}
