package model

import (
	"encoding/json"

	"strings"
)

// 活体检测结果，VideoDetectResult结构见[VideoDetectResult](https://support.huaweicloud.com/api-face/face_02_0010.html)。 调用失败时无此字段。
type LiveDetectRespVideoresult struct {
	// 是否是活体。

	Alive *bool `json:"alive,omitempty"`
	// 动作列表。

	Actions *[]ActionsList `json:"actions,omitempty"`
	// 检测出最大人脸的图片base64。

	Picture *string `json:"picture,omitempty"`
}

func (o LiveDetectRespVideoresult) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "LiveDetectRespVideoresult struct{}"
	}

	return strings.Join([]string{"LiveDetectRespVideoresult", string(data)}, " ")
}
