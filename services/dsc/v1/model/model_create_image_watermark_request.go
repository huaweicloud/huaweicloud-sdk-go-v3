package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateImageWatermarkRequest struct {
	Body *CreateImageWatermarkRequestBody `json:"body,omitempty" xml:"body" type:"multipart"`
}

func (o CreateImageWatermarkRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateImageWatermarkRequest struct{}"
	}

	return strings.Join([]string{"CreateImageWatermarkRequest", string(data)}, " ")
}
