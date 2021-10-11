package model

import (
	"encoding/json"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/def"
	"strings"
)

type DetectFaceByFileRequestBody struct {
	// 本地图片文件，图片不能超过8MB。上传文件时，请求格式为multipart。
	ImageFile *def.FilePart `json:"image_file"`

	// 是否返回人脸属性，希望获取的属性列表，多个属性间使用逗号（,）隔开。目前支持的属性有： • 2：年龄 • 4：装束（帽子、眼镜） • 6：口罩 • 7：发型 • 8：胡须 • 11：图片类型 • 12：质量 • 13：表情 • 21：人脸图片旋转角（顺时针偏转角度），支持0°、90°、180°和270°图片旋转
	Attributes *def.MultiPart `json:"attributes,omitempty"`
}

func (o DetectFaceByFileRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DetectFaceByFileRequestBody struct{}"
	}

	return strings.Join([]string{"DetectFaceByFileRequestBody", string(data)}, " ")
}
