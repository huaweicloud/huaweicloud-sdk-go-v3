package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DownloadIpdIssueAttachmentRequest Request Object
type DownloadIpdIssueAttachmentRequest struct {

	// 工作项所属项目Id
	ProjectId string `json:"project_id"`

	// 附件Id
	Id string `json:"id"`
}

func (o DownloadIpdIssueAttachmentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadIpdIssueAttachmentRequest struct{}"
	}

	return strings.Join([]string{"DownloadIpdIssueAttachmentRequest", string(data)}, " ")
}
