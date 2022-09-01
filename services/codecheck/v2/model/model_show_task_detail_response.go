package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowTaskDetailResponse struct {

	// 任务id
	TaskId *string `json:"task_id,omitempty" xml:"task_id"`

	// 任务名字
	TaskName *string `json:"task_name,omitempty" xml:"task_name"`

	// 创建者id
	CreatorId *string `json:"creator_id,omitempty" xml:"creator_id"`

	// 代码仓地址
	GitUrl *string `json:"git_url,omitempty" xml:"git_url"`

	// 代码仓分支,如果是MR模式，为源分支
	GitBranch *string `json:"git_branch,omitempty" xml:"git_branch"`

	// 上一次检查时间
	LastCheckTime *string `json:"last_check_time,omitempty" xml:"last_check_time"`

	// 代码总行数
	CodeLineTotal *int32 `json:"code_line_total,omitempty" xml:"code_line_total"`

	// 代码有效行数
	CodeLine *int32 `json:"code_line,omitempty" xml:"code_line"`

	// 代码质量
	CodeQuality float32 `json:"code_quality,omitempty" xml:"code_quality"`

	// 问题数
	IssueCount *int32 `json:"issue_count,omitempty" xml:"issue_count"`

	// 危险系数
	RiskCoefficient float32 `json:"risk_coefficient,omitempty" xml:"risk_coefficient"`

	// 重复比例
	DuplicationRatio *string `json:"duplication_ratio,omitempty" xml:"duplication_ratio"`

	// 复杂度
	ComplexityCount *int32 `json:"complexity_count,omitempty" xml:"complexity_count"`

	// 重复行数
	DuplicatedLines *int32 `json:"duplicated_lines,omitempty" xml:"duplicated_lines"`

	// 注释行数
	CommentLines *int32 `json:"comment_lines,omitempty" xml:"comment_lines"`

	// 注释比例
	CommentRatio *string `json:"comment_ratio,omitempty" xml:"comment_ratio"`

	// 重复块
	DuplicatedBlocks *int32 `json:"duplicated_blocks,omitempty" xml:"duplicated_blocks"`

	// 上次执行时间
	LastExecTime *string `json:"last_exec_time,omitempty" xml:"last_exec_time"`

	// 检查类型
	CheckType *string `json:"check_type,omitempty" xml:"check_type"`

	// 创建时间
	CreatedAt *string `json:"created_at,omitempty" xml:"created_at"`

	// 代码平均复杂度
	CyclomaticComplexityPerMethod *string `json:"cyclomatic_complexity_per_method,omitempty" xml:"cyclomatic_complexity_per_method"`

	// 代码平均复杂度(文件)
	CyclomaticComplexityPerFile *string `json:"cyclomatic_complexity_per_file,omitempty" xml:"cyclomatic_complexity_per_file"`

	// 致命问题数
	CriticalCount *string `json:"critical_count,omitempty" xml:"critical_count"`

	// 严重问题数
	MajorCount *string `json:"major_count,omitempty" xml:"major_count"`

	// 一般问题数
	MinorCount *string `json:"minor_count,omitempty" xml:"minor_count"`

	// 提示问题数
	SuggestionCount *string `json:"suggestion_count,omitempty" xml:"suggestion_count"`

	// 门禁质量是否通过
	IsAccess *string `json:"is_access,omitempty" xml:"is_access"`

	// 任务触发方式
	TriggerType *string `json:"trigger_type,omitempty" xml:"trigger_type"`

	// 文件重复率
	FileDuplicationRatio *string `json:"file_duplication_ratio,omitempty" xml:"file_duplication_ratio"`
	HttpStatusCode       int     `json:"-"`
}

func (o ShowTaskDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTaskDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowTaskDetailResponse", string(data)}, " ")
}
