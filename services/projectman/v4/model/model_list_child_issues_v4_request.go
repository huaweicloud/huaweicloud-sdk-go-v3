package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListChildIssuesV4Request struct {
	// devcloud的项目id

	ProjectId string `json:"project_id"`
	// 工作项id

	IssueId int32 `json:"issue_id"`
}

func (o ListChildIssuesV4Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListChildIssuesV4Request struct{}"
	}

	return strings.Join([]string{"ListChildIssuesV4Request", string(data)}, " ")
}
