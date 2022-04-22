package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LiveDetectFaceBase64Req struct {

	// 图像数据，Base64编码，要求： • Base64编码后大小不超过8MB。 • 图片编码格式： JPG、PNG、JPEG、BMP格式的图片。
	ImageBase64 string `json:"image_base64"`
}

func (o LiveDetectFaceBase64Req) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LiveDetectFaceBase64Req struct{}"
	}

	return strings.Join([]string{"LiveDetectFaceBase64Req", string(data)}, " ")
}
