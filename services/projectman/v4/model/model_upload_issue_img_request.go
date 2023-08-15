package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UploadIssueImgRequest Request Object
type UploadIssueImgRequest struct {

	// devcloud项目的32位id
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
