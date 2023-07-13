package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DbMaskTaskInfo struct {

	// DB类型
	DbType *string `json:"db_type,omitempty"`

	// 任务结束时间
	EndTime *int64 `json:"end_time,omitempty"`

	// 执行行数
	ExecuteLine *int32 `json:"execute_line,omitempty"`

	// 任务ID
	Id *string `json:"id,omitempty"`

	// 执行进度
	Progress *int32 `json:"progress,omitempty"`

	// 任务运行状态
	RunStatus *string `json:"run_status,omitempty"`

	// 任务开始时间
	StartTime *int64 `json:"start_time,omitempty"`

	// 任务模板ID
	TaskTemplateId *string `json:"task_template_id,omitempty"`

	// 任务类型
	Type *string `json:"type,omitempty"`
}

func (o DbMaskTaskInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DbMaskTaskInfo struct{}"
	}

	return strings.Join([]string{"DbMaskTaskInfo", string(data)}, " ")
}
