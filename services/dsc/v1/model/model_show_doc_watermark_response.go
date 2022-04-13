package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowDocWatermarkResponse struct {
	// 暗水印内容，长度不超过32个字节

	Watermark      *string `json:"watermark,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowDocWatermarkResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDocWatermarkResponse struct{}"
	}

	return strings.Join([]string{"ShowDocWatermarkResponse", string(data)}, " ")
}
