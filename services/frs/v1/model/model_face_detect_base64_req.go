package model

import (
	"encoding/json"

	"strings"
)

type FaceDetectBase64Req struct {
	// 图像数据，Base64编码，要求： • Base64编码后大小不超过8MB，建议小于1MB。 • 图片为JPG/JPEG/BMP/PNG格式。

	ImageBase64 string `json:"image_base64"`
	// 是否返回人脸属性，希望获取的属性列表，多个属性间使用逗号（,）隔开。目前支持的属性有： • 0：人脸姿态 • 1：性别 • 2：年龄 • 3：人脸关键点 • 4：装束（帽子、眼镜） • 5：笑脸

	Attributes *string `json:"attributes,omitempty"`
}

func (o FaceDetectBase64Req) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "FaceDetectBase64Req struct{}"
	}

	return strings.Join([]string{"FaceDetectBase64Req", string(data)}, " ")
}
