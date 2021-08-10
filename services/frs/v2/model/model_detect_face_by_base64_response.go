package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DetectFaceByBase64Response struct {
	// 检测到的人脸。 调用失败时无此字段。

	Faces          *[]DetectFace `json:"faces,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o DetectFaceByBase64Response) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DetectFaceByBase64Response struct{}"
	}

	return strings.Join([]string{"DetectFaceByBase64Response", string(data)}, " ")
}
