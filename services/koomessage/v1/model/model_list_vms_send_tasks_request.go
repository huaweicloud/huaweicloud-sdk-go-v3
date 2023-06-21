package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListVmsSendTasksRequest struct {

	// 智能信息基础版任务名称。
	TaskName *string `json:"task_name,omitempty"`

	// 智能信息基础版任务ID。
	TaskId *string `json:"task_id,omitempty"`

	// 智能信息基础版模板ID。
	TplId *string `json:"tpl_id,omitempty"`

	// 智能信息基础版模板名称。
	TplName *string `json:"tpl_name,omitempty"`

	// 智能信息基础版任务创建开始时间。 样例为：2019-10-12T07:20:50Z。
	BeginTime *string `json:"begin_time,omitempty"`

	// 智能信息基础版任务创建结束时间。 样例为：2019-10-12T07:20:50Z。
	EndTime *string `json:"end_time,omitempty"`

	// 智能信息基础版任务发送开始时间。 样例为：2019-10-12T07:20:50Z。
	SendBeginTime *string `json:"send_begin_time,omitempty"`

	// 智能信息基础版任务发送结束时间。 样例为：2019-10-12T07:20:50Z。
	SendEndTime *string `json:"send_end_time,omitempty"`

	// 操作员。
	Operator *string `json:"operator,omitempty"`

	// 发送状态类型。 - submit_success：提交成功 - submit_failed：提交失败 - processing：发送中 - sending_failed：发送失败 - re_submit_success：重试提交成功 - sending_complete：发送完成 - re_submit_failed：重试提交成功 - re_processing：重试进行中 - re_sending_complete：重试发送完成 - scheduled：定时任务
	TaskStatus *string `json:"task_status,omitempty"`

	// 偏移量，表示从此偏移量开始查询，offset大于等于0。
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的条目数量。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListVmsSendTasksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVmsSendTasksRequest struct{}"
	}

	return strings.Join([]string{"ListVmsSendTasksRequest", string(data)}, " ")
}
