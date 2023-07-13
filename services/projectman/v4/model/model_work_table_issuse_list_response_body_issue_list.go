package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type WorkTableIssuseListResponseBodyIssueList struct {

	// 工作项id
	Id *int32 `json:"id,omitempty"`

	// 工作项标题
	Subject *string `json:"subject,omitempty"`

	// 父工作项id
	ParentIssueId *int32 `json:"parent_issue_id,omitempty"`

	ParentIssue *WorkTableIssuseListResponseBodyParentIssue `json:"parent_issue,omitempty"`

	Project *WorkTableIssuseListResponseBodyProject `json:"project,omitempty"`

	// 发布版本
	ReleaseDev *string `json:"release_dev,omitempty"`

	// 发现发布版本
	FindReleaseDev *string `json:"find_release_dev,omitempty"`

	// 工作项完成度
	DoneRatio *int32 `json:"done_ratio,omitempty"`

	// 预计工时
	ExpectedWorkHours *float64 `json:"expected_work_hours,omitempty"`

	// 实际工时
	ActualWorkHours *float64 `json:"actual_work_hours,omitempty"`

	Tracker *WorkTableIssuseListResponseBodyTracker `json:"tracker,omitempty"`

	Order *WorkTableIssuseListResponseBodyOrder `json:"order,omitempty"`

	Severity *WorkTableIssuseListResponseBodySeverity `json:"severity,omitempty"`

	Priority *WorkTableIssuseListResponseBodyPriority `json:"priority,omitempty"`

	Domain *WorkTableIssuseListResponseBodyDomain `json:"domain,omitempty"`

	// 排序数值
	Position *float64 `json:"position,omitempty"`

	Module *WorkTableIssuseListResponseBodyModule `json:"module,omitempty"`

	AssignedTo *SimpleUserIn `json:"assigned_to,omitempty"`

	Author *SimpleUserIn `json:"author,omitempty"`

	Developer *SimpleUserIn `json:"developer,omitempty"`

	Closeder *SimpleUserIn `json:"closeder,omitempty"`

	Status *WorkTableIssuseListResponseBodyStatus `json:"status,omitempty"`

	// 是否删除工作项
	Deleted *bool `json:"deleted,omitempty"`

	// 是否关注工作项
	IsWatcher *bool `json:"is_watcher,omitempty"`

	// 关闭标志
	ClosedFlag *int32 `json:"closed_flag,omitempty"`

	// 工作项新建时间戳
	CreatedOn *string `json:"created_on,omitempty"`

	// 工作项更新时间戳
	UpdatedOn *string `json:"updated_on,omitempty"`

	// 工作项预计结束时间戳
	DueDate *string `json:"due_date,omitempty"`
}

func (o WorkTableIssuseListResponseBodyIssueList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkTableIssuseListResponseBodyIssueList struct{}"
	}

	return strings.Join([]string{"WorkTableIssuseListResponseBodyIssueList", string(data)}, " ")
}
