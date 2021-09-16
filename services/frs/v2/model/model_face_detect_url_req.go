package model

import (
	"encoding/json"

	"strings"
)

type FaceDetectUrlReq struct {
	// 图片的URL路径，目前仅支持华为云上OBS的URL，且人脸识别服务有权限读取该OBS桶的数据。 开通读取权限的操作请参见[服务授权](https://support.huaweicloud.com/api-face/face_02_0006.html)。

	ImageUrl string `json:"image_url"`
	// 是否返回人脸属性，希望获取的属性列表，多个属性间使用逗号（,）隔开。目前支持的属性有：   • 1：性别   • 2：年龄   • 4：装束（帽子、眼镜）   • 6：口罩   • 7：发型   • 8：胡须   • 9：肤色   • 10：民族   • 11：图片类型   • 12：质量   • 13：表情   • 21：人脸图片旋转角（顺时针偏转角度），支持0°、90°、180°和270°图片旋转

	Attributes *string `json:"attributes,omitempty"`
}

func (o FaceDetectUrlReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "FaceDetectUrlReq struct{}"
	}

	return strings.Join([]string{"FaceDetectUrlReq", string(data)}, " ")
}
