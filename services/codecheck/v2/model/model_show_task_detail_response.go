package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowTaskDetailResponse struct {
	// 任务id

	TaskId *string `json:"task_id,omitempty"`
	// 任务名字

	TaskName *string `json:"task_name,omitempty"`
	// 创建者id

	CreatorId *string `json:"creator_id,omitempty"`
	// 代码仓地址

	GitUrl *string `json:"git_url,omitempty"`
	// 代码仓分支,如果是MR模式，为源分支

	GitBranch *string `json:"git_branch,omitempty"`
	// 上一次检查时间

	LastCheckTime *string `json:"last_check_time,omitempty"`
	// 代码总行数

	CodeLineTotal *int32 `json:"code_line_total,omitempty"`
	// 代码有效行数

	CodeLine *int32 `json:"code_line,omitempty"`
	// 代码质量

	CodeQuality float32 `json:"code_quality,omitempty"`
	// 问题数

	IssueCount *int32 `json:"issue_count,omitempty"`
	// 危险系数

	RiskCoefficient float32 `json:"risk_coefficient,omitempty"`
	// 重复比例

	DuplicationRatio *string `json:"duplication_ratio,omitempty"`
	// 复杂度

	ComplexityCount *int32 `json:"complexity_count,omitempty"`
	// 重复行数

	DuplicatedLines *int32 `json:"duplicated_lines,omitempty"`
	// 注释行数

	CommentLines *int32 `json:"comment_lines,omitempty"`
	// 注释比例

	CommentRatio *string `json:"comment_ratio,omitempty"`
	// 重复块

	DuplicatedBlocks *int32 `json:"duplicated_blocks,omitempty"`
	// 上次执行时间

	LastExecTime *string `json:"last_exec_time,omitempty"`
	// 检查类型

	CheckType *string `json:"check_type,omitempty"`
	// 创建时间

	CreatedAt *string `json:"created_at,omitempty"`
	// 代码平均复杂度

	CyclomaticComplexityPerMethod *string `json:"cyclomatic_complexity_per_method,omitempty"`
	// 致命问题数

	CriticalCount *string `json:"critical_count,omitempty"`
	// 严重问题数

	MajorCount *string `json:"major_count,omitempty"`
	// 一般问题数

	MinorCount *string `json:"minor_count,omitempty"`
	// 提示问题数

	SuggestionCount *string `json:"suggestion_count,omitempty"`
	// 门禁质量是否通过

	IsAccess *string `json:"is_access,omitempty"`
	// 任务触发方式

	TriggerType    *string `json:"trigger_type,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowTaskDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTaskDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowTaskDetailResponse", string(data)}, " ")
}
