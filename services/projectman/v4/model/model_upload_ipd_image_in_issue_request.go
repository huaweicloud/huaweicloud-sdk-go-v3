package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UploadIpdImageInIssueRequest Request Object
type UploadIpdImageInIssueRequest struct {
	ProjectId string `json:"project_id"`

	// 工作项Id
	IssueId string `json:"issue_id"`

	Body *UploadIpdImageInIssueRequestBody `json:"body,omitempty" type:"multipart"`
}

func (o UploadIpdImageInIssueRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadIpdImageInIssueRequest struct{}"
	}

	return strings.Join([]string{"UploadIpdImageInIssueRequest", string(data)}, " ")
}
