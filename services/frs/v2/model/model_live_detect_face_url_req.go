package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LiveDetectFaceUrlReq struct {

	// [图片的URL路径，目前仅支持华为云上OBS的URL，且人脸识别服务有权限读取该OBS桶的数据。开通读取权限的操作请参见[服务授权](https://support.huaweicloud.com/api-face/face_02_0006.html)。(tag:hc) [图片的URL路径，目前仅支持华为云上OBS的URL，且人脸识别服务有权限读取该OBS桶的数据。开通读取权限的操作请参见[服务授权](https://support.huaweicloud.com/intl/zh-cn/api-face/face_02_0006.html)。(tag:hk)
	ImageUrl string `json:"image_url"`
}

func (o LiveDetectFaceUrlReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LiveDetectFaceUrlReq struct{}"
	}

	return strings.Join([]string{"LiveDetectFaceUrlReq", string(data)}, " ")
}
