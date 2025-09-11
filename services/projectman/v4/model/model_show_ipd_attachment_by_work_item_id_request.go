package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowIpdAttachmentByWorkItemIdRequest Request Object
type ShowIpdAttachmentByWorkItemIdRequest struct {

	// 工作项所属项目Id
	ProjectId string `json:"project_id"`

	// 工作项Id
	IssueId string `json:"issue_id"`

	// 原始需求跨项目时，提出项目Id
	SourceProjectId *string `json:"source_project_id,omitempty"`
}

func (o ShowIpdAttachmentByWorkItemIdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowIpdAttachmentByWorkItemIdRequest struct{}"
	}

	return strings.Join([]string{"ShowIpdAttachmentByWorkItemIdRequest", string(data)}, " ")
}
