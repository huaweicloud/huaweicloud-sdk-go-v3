package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAttachmentRequest Request Object
type DeleteAttachmentRequest struct {

	// devcloud项目的32位id
	ProjectId string `json:"project_id"`

	// 工作项id
	IssueId string `json:"issue_id"`

	// 附件id
	AttachmentId string `json:"attachment_id"`
}

func (o DeleteAttachmentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAttachmentRequest struct{}"
	}

	return strings.Join([]string{"DeleteAttachmentRequest", string(data)}, " ")
}
