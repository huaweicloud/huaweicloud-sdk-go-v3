package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DownloadIpdIssueAttachmentResponse Response Object
type DownloadIpdIssueAttachmentResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DownloadIpdIssueAttachmentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadIpdIssueAttachmentResponse struct{}"
	}

	return strings.Join([]string{"DownloadIpdIssueAttachmentResponse", string(data)}, " ")
}
