package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
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
