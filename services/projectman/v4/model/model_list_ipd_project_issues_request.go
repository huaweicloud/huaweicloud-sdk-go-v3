package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListIpdProjectIssuesRequest Request Object
type ListIpdProjectIssuesRequest struct {

	// devcloud项目的32位id
	ProjectId string `json:"project_id"`

	// 是否backlog查询
	IsBacklog *bool `json:"is_backlog,omitempty"`

	// 工作项分类：[Epic,FE,IR,RR,SR,US,AR,Bug,Task]
	IssueType string `json:"issue_type"`

	// 提出项目Id
	SrcDomainId *string `json:"src_domain_id,omitempty"`

	// 视图模式[tree,list]
	View *string `json:"view,omitempty"`

	Body *SearchIpdIssuesRequestBody `json:"body,omitempty"`
}

func (o ListIpdProjectIssuesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIpdProjectIssuesRequest struct{}"
	}

	return strings.Join([]string{"ListIpdProjectIssuesRequest", string(data)}, " ")
}
