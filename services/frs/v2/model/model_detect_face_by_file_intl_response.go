package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DetectFaceByFileIntlResponse Response Object
type DetectFaceByFileIntlResponse struct {

	// 检测到的人脸。 调用失败时无此字段。
	Faces          *[]DetectFace `json:"faces,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o DetectFaceByFileIntlResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetectFaceByFileIntlResponse struct{}"
	}

	return strings.Join([]string{"DetectFaceByFileIntlResponse", string(data)}, " ")
}
