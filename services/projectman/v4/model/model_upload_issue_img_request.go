package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UploadIssueImgRequest struct {
	// devcloud的项目ID

	ProjectId string `json:"project_id"`

	Body *UploadIssueImgRequestBody `json:"body,omitempty" type:"multipart"`
}

func (o UploadIssueImgRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadIssueImgRequest struct{}"
	}

	return strings.Join([]string{"UploadIssueImgRequest", string(data)}, " ")
}
