package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateImageWatermarkRequest Request Object
type CreateImageWatermarkRequest struct {
	Body *CreateImageWatermarkRequestBody `json:"body,omitempty" type:"multipart"`
}

func (o CreateImageWatermarkRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateImageWatermarkRequest struct{}"
	}

	return strings.Join([]string{"CreateImageWatermarkRequest", string(data)}, " ")
}
