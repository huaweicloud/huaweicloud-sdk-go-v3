package model

import (
	"encoding/json"

	"strings"
)

type FaceDetectUrlReq struct {
	// 图片的URL路径，目前仅支持华为云上OBS的URL，且人脸识别服务有权限读取该OBS桶的数据。 开通读取权限的操作请参见服务授权。

	ImageUrl string `json:"image_url"`
	// 是否返回人脸属性，希望获取的属性列表，多个属性间使用逗号（,）隔开。目前支持的属性有： • 0：人脸姿态 • 1：性别 • 2：年龄 • 3：人脸关键点 • 4：装束（帽子、眼镜） • 5：笑脸

	Attributes *string `json:"attributes,omitempty"`
}

func (o FaceDetectUrlReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "FaceDetectUrlReq struct{}"
	}

	return strings.Join([]string{"FaceDetectUrlReq", string(data)}, " ")
}
