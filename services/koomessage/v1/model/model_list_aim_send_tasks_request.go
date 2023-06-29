package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAimSendTasksRequest Request Object
type ListAimSendTasksRequest struct {

	// 智能信息发送任务ID。
	TaskId *string `json:"task_id,omitempty"`

	// 智能信息发送任务名称。
	TaskName *string `json:"task_name,omitempty"`

	// 智能信息模板ID。
	TplId *string `json:"tpl_id,omitempty"`

	// 智能信息模板名称。
	TplName *string `json:"tpl_name,omitempty"`

	// 智能信息发送任务创建开始时间。样例：2019-10-12T07:20:50.522Z。  > 需同时传入end_time才能生效，单独传begin_time不会作为过滤条件。缺省：查询最近7天（168小时）数据。
	BeginTime *string `json:"begin_time,omitempty"`

	// 智能信息发送任务创建结束时间。样例：2019-10-12T07:20:50.522Z。  > 需同时传入begin_time才能生效，单独传end_time不会作为过滤条件。缺省：查询最近7天（168小时）数据。
	EndTime *string `json:"end_time,omitempty"`

	// 智能信息发送任务状态。  - Success：创建成功  - Fail：创建失败
	TaskStatus *string `json:"task_status,omitempty"`

	// 偏移量，表示从此偏移量开始查询，offset大于等于0。  >为提高查询效率，offset+limit须小于等于10000，超出范围查询为空。
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的条目数量。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListAimSendTasksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAimSendTasksRequest struct{}"
	}

	return strings.Join([]string{"ListAimSendTasksRequest", string(data)}, " ")
}
