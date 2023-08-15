package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowImageWatermarkByAddressResponse Response Object
type ShowImageWatermarkByAddressResponse struct {

	// 暗水印内容，长度不超过32个字节
	Watermark      *string `json:"watermark,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowImageWatermarkByAddressResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowImageWatermarkByAddressResponse struct{}"
	}

	return strings.Join([]string{"ShowImageWatermarkByAddressResponse", string(data)}, " ")
}
