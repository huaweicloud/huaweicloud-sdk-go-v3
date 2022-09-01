package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FaceCompareUrlReq struct {

	// [图片的URL路径，目前仅支持华为云上OBS的URL，且人脸识别服务有权限读取该OBS桶的数据。开通读取权限的操作请参见[服务授权](https://support.huaweicloud.com/api-face/face_02_0006.html)。与image1_file、image1_base64三选一](tag:hc) [图片的URL路径，目前仅支持华为云上OBS的URL，且人脸识别服务有权限读取该OBS桶的数据。开通读取权限的操作请参见[服务授权](https://support.huaweicloud.com/intl/zh-cn/api-face/face_02_0006.html)。与image1_file、image1_base64三选一](tag:hk)
	Image1Url string `json:"image1_url" xml:"image1_url"`

	// [图片的URL路径，目前仅支持华为云上OBS的URL，且人脸识别服务有权限读取该OBS桶的数据。开通读取权限的操作请参见[服务授权](https://support.huaweicloud.com/api-face/face_02_0006.html)。与image2_file、image2_base64三选一](tag:hc) [图片的URL路径，目前仅支持华为云上OBS的URL，且人脸识别服务有权限读取该OBS桶的数据。开通读取权限的操作请参见[服务授权](https://support.huaweicloud.com/intl/zh-cn/api-face/face_02_0006.html)。与image2_file、image2_base64三选一](tag:hk)
	Image2Url string `json:"image2_url" xml:"image2_url"`
}

func (o FaceCompareUrlReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FaceCompareUrlReq struct{}"
	}

	return strings.Join([]string{"FaceCompareUrlReq", string(data)}, " ")
}
