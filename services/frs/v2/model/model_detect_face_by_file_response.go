package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DetectFaceByFileResponse Response Object
type DetectFaceByFileResponse struct {

	// 检测到的人脸。 调用失败时无此字段。
	Faces          *[]DetectFace `json:"faces,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o DetectFaceByFileResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetectFaceByFileResponse struct{}"
	}

	return strings.Join([]string{"DetectFaceByFileResponse", string(data)}, " ")
}
