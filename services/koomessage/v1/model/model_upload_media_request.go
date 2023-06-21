package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UploadMediaRequest struct {
	Body *UploadMediaRequestBody `json:"body,omitempty" type:"multipart"`
}

func (o UploadMediaRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadMediaRequest struct{}"
	}

	return strings.Join([]string{"UploadMediaRequest", string(data)}, " ")
}
