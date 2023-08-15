package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UploadAttachmentsRequest Request Object
type UploadAttachmentsRequest struct {

	// 项目id
	ProjectId string `json:"project_id"`

	// 工作项id
	IssueId string `json:"issue_id"`

	Body *UploadAttachmentsRequestBody `json:"body,omitempty" type:"multipart"`
}

func (o UploadAttachmentsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadAttachmentsRequest struct{}"
	}

	return strings.Join([]string{"UploadAttachmentsRequest", string(data)}, " ")
}
