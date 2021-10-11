package model

import (
	"encoding/json"

	"strings"
)

type FaceDetectBase64Req struct {
	// 图像数据，Base64编码，要求： • Base64编码后大小不超过8MB，建议小于1MB。 • 图片为JPG/JPEG/BMP/PNG格式。

	ImageBase64 string `json:"image_base64"`
	// 是否返回人脸属性，希望获取的属性列表，多个属性间使用逗号（,）隔开。目前支持的属性有：   • 2：年龄   • 4：装束（帽子、眼镜）   • 6：口罩   • 7：发型   • 8：胡须   • 11：图片类型   • 12：质量   • 13：表情   • 21：人脸图片旋转角（顺时针偏转角度），支持0°、90°、180°和270°图片旋转

	Attributes *string `json:"attributes,omitempty"`
}

func (o FaceDetectBase64Req) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "FaceDetectBase64Req struct{}"
	}

	return strings.Join([]string{"FaceDetectBase64Req", string(data)}, " ")
}
