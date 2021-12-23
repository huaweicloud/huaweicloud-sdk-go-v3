package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListIssueCommentsV4Request struct {
	// devcloud的项目id

	ProjectId string `json:"project_id"`
	// 工作项id

	IssueId int32 `json:"issue_id"`
	// 分页索引，偏移量

	Offset *int32 `json:"offset,omitempty"`
	// 每页显示的条数,最大显示100条

	Limit *int32 `json:"limit,omitempty"`
}

func (o ListIssueCommentsV4Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIssueCommentsV4Request struct{}"
	}

	return strings.Join([]string{"ListIssueCommentsV4Request", string(data)}, " ")
}
