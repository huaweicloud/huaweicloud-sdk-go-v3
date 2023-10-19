package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DetectFaceByUrlIntlResponse Response Object
type DetectFaceByUrlIntlResponse struct {

	// 检测到的人脸。 调用失败时无此字段。
	Faces *[]DetectFace `json:"faces,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DetectFaceByUrlIntlResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetectFaceByUrlIntlResponse struct{}"
	}

	return strings.Join([]string{"DetectFaceByUrlIntlResponse", string(data)}, " ")
}
