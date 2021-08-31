package model

import (
	"encoding/json"

	"strings"
)

// 静默活体检测结果，LivelessDetectResult结构见表[结构格式说明表](zh-cn_topic_0238576582.xml)。 调用失败时无此字段。
type LiveDetectFaceRespResult struct {
	// 是否是活体。

	Alive *bool `json:"alive,omitempty"`
	// 置信度，取值范围0～1。

	Confidence *float64 `json:"confidence,omitempty"`
	// 检测出最大人脸的图片base64字符串。

	Picture *string `json:"picture,omitempty"`
}

func (o LiveDetectFaceRespResult) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "LiveDetectFaceRespResult struct{}"
	}

	return strings.Join([]string{"LiveDetectFaceRespResult", string(data)}, " ")
}
