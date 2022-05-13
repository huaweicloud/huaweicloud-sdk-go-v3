package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// [活体检测结果，VideoDetectResult结构见[VideoDetectResult](https://support.huaweicloud.com/api-face/face_02_0010.html)。 调用失败时无此字段。](tag:hc) [活体检测结果，VideoDetectResult结构见[VideoDetectResult](https://support.huaweicloud.com/intl/zh-cn/api-face/face_02_0010.html)。 调用失败时无此字段。](tag:hk)
type LiveDetectRespVideoresult struct {

	// 是否是活体。
	Alive *bool `json:"alive,omitempty"`

	// 动作列表。
	Actions *[]ActionsList `json:"actions,omitempty"`

	// 检测出最大人脸的图片base64。
	Picture *string `json:"picture,omitempty"`
}

func (o LiveDetectRespVideoresult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LiveDetectRespVideoresult struct{}"
	}

	return strings.Join([]string{"LiveDetectRespVideoresult", string(data)}, " ")
}
