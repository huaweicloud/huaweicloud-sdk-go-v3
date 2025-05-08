package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTaskDetailResponse Response Object
type ShowTaskDetailResponse struct {

	// 任务Id
	TaskId *string `json:"task_id,omitempty"`

	// 任务类型
	TaskType *string `json:"task_type,omitempty"`

	// 任务状态，包含就绪，处理中，成功，失败，已取消
	Status *string `json:"status,omitempty"`

	// 错误码
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误信息
	ErrorMsg *string `json:"error_msg,omitempty"`

	// 任务开始时间
	CreateTime *string `json:"create_time,omitempty"`

	// 任务开始处理时间，多个任务则是第一个任务的开始时间
	StartTime *string `json:"start_time,omitempty"`

	// 任务处理结束时间，多个任务则是最后一个任务的结束时间
	EndTime *string `json:"end_time,omitempty"`

	WorkflowTask *WorkflowTask `json:"workflow_task,omitempty"`

	EditMediaTask *EditMediaTask `json:"edit_media_task,omitempty"`

	// 用户自定义数据
	SessionContext *string `json:"session_context,omitempty"`

	// 客户回调地址
	CallbackUrl    *string `json:"callback_url,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowTaskDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTaskDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowTaskDetailResponse", string(data)}, " ")
}
