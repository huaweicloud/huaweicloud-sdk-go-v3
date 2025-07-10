package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type JobDetailInfo struct {

	// 子任务ID。
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

	// 任务执行的服务器IP。
	Host *string `json:"host,omitempty"`

	// 项目ID。
	ProjectId *string `json:"project_id,omitempty"`

	// 任务ID。
	JobId *string `json:"job_id,omitempty"`

	// 任务进度。
	Process *int32 `json:"process,omitempty"`

	// 关联用户。
	AttachUser *string `json:"attach_user,omitempty"`

	// 操作对象。
	Entity *string `json:"entity,omitempty"`

	// ip地址。
	IpAddress *string `json:"ip_address,omitempty"`
}

func (o JobDetailInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobDetailInfo struct{}"
	}

	return strings.Join([]string{"JobDetailInfo", string(data)}, " ")
}
