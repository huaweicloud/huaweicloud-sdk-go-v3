package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DownloadAttachmentRequest Request Object
type DownloadAttachmentRequest struct {

	// devcloud项目的32位id
	ProjectId string `json:"project_id"`

	// 工作项id
	IssueId string `json:"issue_id"`

	// 附件id
	AttachmentId string `json:"attachment_id"`
}

func (o DownloadAttachmentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadAttachmentRequest struct{}"
	}

	return strings.Join([]string{"DownloadAttachmentRequest", string(data)}, " ")
}
