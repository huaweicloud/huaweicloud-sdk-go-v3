package model

import (
	"encoding/json"

	"strings"
)

type LiveDetectFaceUrlReq struct {
	// 图片的URL路径，目前仅支持华为云上OBS的URL，且人脸识别服务有权限读取该OBS桶的数据。开通读取权限的操作请参见[申请服务](zh-cn_topic_0107696818.xml)。

	ImageUrl string `json:"image_url"`
}

func (o LiveDetectFaceUrlReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "LiveDetectFaceUrlReq struct{}"
	}

	return strings.Join([]string{"LiveDetectFaceUrlReq", string(data)}, " ")
}
