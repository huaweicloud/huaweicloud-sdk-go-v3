package model

import (
	"encoding/json"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/def"
	"strings"
)

type DetectFaceByFileRequestBody struct {
	// 本地图片文件，图片不能超过8MB，建议小于1MB。上传文件时，请求格式为multipart。
	ImageFile *def.FilePart `json:"image_file"`

	// 是否返回人脸属性，希望获取的属性列表，多个属性间使用逗号（,）隔开。目前支持的属性有： • 0：人脸姿态 • 1：性别 • 2：年龄 • 3：人脸关键点 • 4：装束（帽子、眼镜） • 5：笑脸
	Attributes *def.MultiPart `json:"attributes,omitempty"`
}

func (o DetectFaceByFileRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DetectFaceByFileRequestBody struct{}"
	}

	return strings.Join([]string{"DetectFaceByFileRequestBody", string(data)}, " ")
}
