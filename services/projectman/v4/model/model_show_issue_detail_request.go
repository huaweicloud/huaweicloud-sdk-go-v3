package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowIssueDetailRequest Request Object
type ShowIssueDetailRequest struct {

	// devcloud项目的32位id
	ProjectId string `json:"project_id"`

	// 工作项唯一Id
	IssueId string `json:"issue_id"`

	// 工作项分类
	IssueType string `json:"issue_type"`

	// 项目所属domainId
	DomainId *string `json:"domain_id,omitempty"`
}

func (o ShowIssueDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowIssueDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowIssueDetailRequest", string(data)}, " ")
}
