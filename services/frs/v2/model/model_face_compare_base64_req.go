package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FaceCompareBase64Req struct {

	// 图像数据，Base64编码，要求： • Base64编码后大小不超过8MB，建议小于1MB。 • 图片为JPG/JPEG/BMP/PNG格式。
	Image2Base64 string `json:"image2_base64" xml:"image2_base64"`

	// 图像数据，Base64编码，要求： • Base64编码后大小不超过8MB，建议小于1MB。 • 图片为JPG/JPEG/BMP/PNG格式。
	Image1Base64 string `json:"image1_base64" xml:"image1_base64"`
}

func (o FaceCompareBase64Req) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FaceCompareBase64Req struct{}"
	}

	return strings.Join([]string{"FaceCompareBase64Req", string(data)}, " ")
}
