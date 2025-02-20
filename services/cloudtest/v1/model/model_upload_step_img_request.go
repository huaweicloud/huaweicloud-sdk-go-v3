package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UploadStepImgRequest Request Object
type UploadStepImgRequest struct {

	// 项目id
	ProjectId string `json:"project_id"`

	Body *UploadStepImgRequestBody `json:"body,omitempty" type:"multipart"`
}

func (o UploadStepImgRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadStepImgRequest struct{}"
	}

	return strings.Join([]string{"UploadStepImgRequest", string(data)}, " ")
}
