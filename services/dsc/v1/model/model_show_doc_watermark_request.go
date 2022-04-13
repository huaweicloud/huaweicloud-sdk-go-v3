package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowDocWatermarkRequest struct {
	Body *ShowDocWatermarkRequestBody `json:"body,omitempty" type:"multipart"`
}

func (o ShowDocWatermarkRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDocWatermarkRequest struct{}"
	}

	return strings.Join([]string{"ShowDocWatermarkRequest", string(data)}, " ")
}
