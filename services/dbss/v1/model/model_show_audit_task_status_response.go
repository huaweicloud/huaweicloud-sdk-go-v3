package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAuditTaskStatusResponse Response Object
type ShowAuditTaskStatusResponse struct {

	// 开始时间
	BeginTime *int64 `json:"begin_time,omitempty"`

	// 业务类型  - audit：审计  - risk：风险
	BusiType *string `json:"busi_type,omitempty"`

	// 已统计实例数
	CompletedNum *int32 `json:"completed_num,omitempty"`

	// 结束时间
	EndTime *int64 `json:"end_time,omitempty"`

	// Task任务ID
	Id *int64 `json:"id,omitempty"`

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// 查询条件：开始时间
	QueryBeginTime *int64 `json:"query_begin_time,omitempty"`

	// 查询条件：结束时间
	QueryEndTime *int64 `json:"query_end_time,omitempty"`

	// 任务状态  - 0：已就绪  - 1：运行中  - 2：成功  - 3：失败
	TaskStatus *string `json:"task_status,omitempty"`

	// 任务开关
	TaskSwitch *string `json:"task_switch,omitempty"`

	// 总实例数
	TotalNum       *int32 `json:"total_num,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowAuditTaskStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAuditTaskStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowAuditTaskStatusResponse", string(data)}, " ")
}
