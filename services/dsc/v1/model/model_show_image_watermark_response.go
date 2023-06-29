package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowImageWatermarkResponse Response Object
type ShowImageWatermarkResponse struct {

	// 暗水印内容，长度不超过32个字节
	Watermark      *string `json:"watermark,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowImageWatermarkResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowImageWatermarkResponse struct{}"
	}

	return strings.Join([]string{"ShowImageWatermarkResponse", string(data)}, " ")
}
