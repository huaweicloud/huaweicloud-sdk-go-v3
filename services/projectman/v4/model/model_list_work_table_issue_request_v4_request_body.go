package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListWorkTableIssueRequestV4RequestBody struct {

	// 偏移量,offset是limit的整数倍，limit=10,offset=0,10,20...
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示数量
	Limit *int32 `json:"limit,omitempty"`

	// 搜索关键词
	Subject *string `json:"subject,omitempty"`

	// 工作项创建时间区间
	CreatedOn *string `json:"created_on,omitempty"`

	// 工作项更新时间区间
	UpdatedOn *string `json:"updated_on,omitempty"`

	// 工作项结束时间区间
	ClosedOn *string `json:"closed_on,omitempty"`

	// 工作项预计开始日期区间
	StartDate *string `json:"start_date,omitempty"`

	// 工作项预计结束日期区间
	DueDate *string `json:"due_date,omitempty"`

	// 工作项类型
	TrackerId *string `json:"tracker_id,omitempty"`

	// 工作项状态
	StatusId *string `json:"status_id,omitempty"`

	// 工作项创建人id
	AuthorId *string `json:"author_id,omitempty"`

	// 工作项开发人员id
	DeveloperId *string `json:"developer_id,omitempty"`

	// 工作项优先级id
	PriorityId *string `json:"priority_id,omitempty"`
}

func (o ListWorkTableIssueRequestV4RequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWorkTableIssueRequestV4RequestBody struct{}"
	}

	return strings.Join([]string{"ListWorkTableIssueRequestV4RequestBody", string(data)}, " ")
}
