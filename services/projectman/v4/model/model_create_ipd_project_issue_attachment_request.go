package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateIpdProjectIssueAttachmentRequest Request Object
type CreateIpdProjectIssueAttachmentRequest struct {

	// devcloud项目的32位id
	ProjectId string `json:"project_id"`

	// 工作项Id
	IssueId string `json:"issue_id"`

	Body *CreateIpdProjectIssueAttachmentRequestBody `json:"body,omitempty" type:"multipart"`
}

func (o CreateIpdProjectIssueAttachmentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateIpdProjectIssueAttachmentRequest struct{}"
	}

	return strings.Join([]string{"CreateIpdProjectIssueAttachmentRequest", string(data)}, " ")
}
