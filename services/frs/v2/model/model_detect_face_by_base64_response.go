package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DetectFaceByBase64Response struct {

	// 检测到的人脸。 调用失败时无此字段。
	Faces          *[]DetectFace `json:"faces,omitempty" xml:"faces"`
	HttpStatusCode int           `json:"-"`
}

func (o DetectFaceByBase64Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetectFaceByBase64Response struct{}"
	}

	return strings.Join([]string{"DetectFaceByBase64Response", string(data)}, " ")
}
