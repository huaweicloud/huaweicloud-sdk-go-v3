package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTaskDetailResponse Response Object
type ShowTaskDetailResponse struct {

	// **参数解释：** 旧版任务ID。即新建检查任务接口的返回值，每个代码检查任务及其下的分支任务、通过API创建的增量任务都有唯一的任务ID。 **取值范围：** 不涉及。
	TaskId *string `json:"task_id,omitempty"`

	// **参数解释：** 任务名称。 **取值范围：** 不涉及。
	TaskName *string `json:"task_name,omitempty"`

	// **参数解释：** 任务创建者用户ID。 **取值范围：** 不涉及。
	CreatorId *string `json:"creator_id,omitempty"`

	// **参数解释：** 代码仓地址。 **取值范围：** 不涉及。
	GitUrl *string `json:"git_url,omitempty"`

	// **参数解释：** 代码仓分支，如果是MR模式，为源分支。 **取值范围：** 不涉及。
	GitBranch *string `json:"git_branch,omitempty"`

	// **参数解释：** 上一次检查时间。使用UTC时间，格式为yyyy-MM-ddTHH:mm:ssZ。例如：2020-09-25T12:05:00Z。 **取值范围：** 不涉及。
	LastCheckTime *string `json:"last_check_time,omitempty"`

	// **参数解释：** 原始代码行数，包含空行和注释行。 **取值范围：** 不涉及。
	CodeLineTotal *int32 `json:"code_line_total,omitempty"`

	// **参数解释：** 有效代码行数，不包含空行和注释行。 **取值范围：** 不涉及。
	CodeLine *int32 `json:"code_line,omitempty"`

	// **参数解释：** 代码质量。 **取值范围：** 不涉及。
	CodeQuality float32 `json:"code_quality,omitempty"`

	// **参数解释：** [待处理问题数。](tag:hws,hws_hk)[待处理和重新打开（需打开忽略审核开关）问题数。](tag:fcs,hcs,hcs_sm,hcs_site,hcs_sitesm) **取值范围：** 不涉及。
	IssueCount *int32 `json:"issue_count,omitempty"`

	// **参数解释：** 危险系数。 **取值范围：** 不涉及。
	RiskCoefficient float32 `json:"risk_coefficient,omitempty"`

	// **参数解释：** 代码重复率。 **取值范围：** 不涉及。
	DuplicationRatio *string `json:"duplication_ratio,omitempty"`

	// **参数解释：** 圈复杂度总数。 **取值范围：** 不涉及。
	ComplexityCount *int32 `json:"complexity_count,omitempty"`

	// **参数解释：** 重复行数总数。 **取值范围：** 不涉及。
	DuplicatedLines *int32 `json:"duplicated_lines,omitempty"`

	// **参数解释：** 注释行数。 **取值范围：** 不涉及。
	CommentLines *int32 `json:"comment_lines,omitempty"`

	// **参数解释：** 注释率。 **取值范围：** 不涉及。
	CommentRatio *string `json:"comment_ratio,omitempty"`

	// **参数解释：** 重复块个数。 **取值范围：** 不涉及。
	DuplicatedBlocks *int32 `json:"duplicated_blocks,omitempty"`

	// **参数解释：** 上一次执行时间。使用UTC时间，格式为yyyy-MM-ddTHH:mm:ssZ。例如：2020-09-25T12:05:00Z。 **取值范围：** 不涉及。
	LastExecTime *string `json:"last_exec_time,omitempty"`

	// **参数解释：** 废弃字段。 **约束限制：** 不涉及。 **取值范围：** 不涉及。
	CheckType *string `json:"check_type,omitempty"`

	// **参数解释：** 创建时间。使用UTC时间，格式为yyyy-MM-ddTHH:mm:ssZ。例如：2020-09-25T12:05:00Z。 **取值范围：** 不涉及。
	CreatedAt *string `json:"created_at,omitempty"`

	// **参数解释：** 平均圈复杂度（方法）。 **取值范围：** 不涉及。
	CyclomaticComplexityPerMethod *string `json:"cyclomatic_complexity_per_method,omitempty"`

	// **参数解释：** 平均圈复杂度（文件）。 **取值范围：** 不涉及。
	CyclomaticComplexityPerFile *string `json:"cyclomatic_complexity_per_file,omitempty"`

	// **参数解释：** 致命问题数。不包括已忽略和已修复问题。 **取值范围：** 不涉及。
	CriticalCount *string `json:"critical_count,omitempty"`

	// **参数解释：** 严重问题数。不包括已忽略和已修复问题。 **取值范围：** 不涉及。
	MajorCount *string `json:"major_count,omitempty"`

	// **参数解释：** 一般问题数。不包括已忽略和已修复问题。 **取值范围：** 不涉及。
	MinorCount *string `json:"minor_count,omitempty"`

	// **参数解释：** 提示问题数。不包括已忽略和已修复问题。 **取值范围：** 不涉及。
	SuggestionCount *string `json:"suggestion_count,omitempty"`

	// **参数解释：** 门禁质量是否通过。 **取值范围：** - 1：门禁通过。 - 0：门禁不通过。
	IsAccess *string `json:"is_access,omitempty"`

	// **参数解释：** 门禁质量是否通过。 **取值范围：** 不涉及。
	ReviewResult *string `json:"review_result,omitempty"`

	// **参数解释：** 任务触发方式。 **取值范围：** - 0：手动触发。 - 1：定时触发。 - 2：代码提交触发。 - 3：合并请求触发。
	TriggerType *string `json:"trigger_type,omitempty"`

	// **参数解释：** 文件重复率。 **取值范围：** 不涉及。
	FileDuplicationRatio *string `json:"file_duplication_ratio,omitempty"`

	// **参数解释：** 新发现问题数。不包括已忽略和已修复问题。 **取值范围：** 不涉及。
	NewCount *int32 `json:"new_count,omitempty"`

	// **参数解释：** 已解决和已修复问题数。 **取值范围：** 不涉及。
	SolveCount *int32 `json:"solve_count,omitempty"`

	// **参数解释：** 已忽略问题数。 **取值范围：** 不涉及。
	DismissedCount *int32 `json:"dismissed_count,omitempty"`

	// **参数解释：** 重复文件数量。 **取值范围：** 不涉及。
	DuplicatedFiles *int32 `json:"duplicated_files,omitempty"`

	// **参数解释：** 新增致命问题数。不包括已忽略和已修复问题。 **取值范围：** 不涉及。
	NewCriticalCount *string `json:"new_critical_count,omitempty"`

	// **参数解释：** 新增严重问题数。不包括已忽略和已修复问题。 **取值范围：** 不涉及。
	NewMajorCount *string `json:"new_major_count,omitempty"`

	// **参数解释：** 新增一般问题数。不包括已忽略和已修复问题。 **取值范围：** 不涉及。
	NewMinorCount *string `json:"new_minor_count,omitempty"`

	// **参数解释：** 新增提示问题数。不包括已忽略和已修复问题。 **取值范围：** 不涉及。
	NewSuggestionCount *string `json:"new_suggestion_count,omitempty"`
	HttpStatusCode     int     `json:"-"`
}

func (o ShowTaskDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTaskDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowTaskDetailResponse", string(data)}, " ")
}
