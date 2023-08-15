package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowImageWatermarkRequest Request Object
type ShowImageWatermarkRequest struct {
	Body *ShowImageWatermarkRequestBody `json:"body,omitempty" type:"multipart"`
}

func (o ShowImageWatermarkRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowImageWatermarkRequest struct{}"
	}

	return strings.Join([]string{"ShowImageWatermarkRequest", string(data)}, " ")
}
