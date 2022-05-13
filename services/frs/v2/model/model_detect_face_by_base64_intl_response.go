package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DetectFaceByBase64IntlResponse struct {

	// 检测到的人脸。 调用失败时无此字段。
	Faces          *[]DetectFace `json:"faces,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o DetectFaceByBase64IntlResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetectFaceByBase64IntlResponse struct{}"
	}

	return strings.Join([]string{"DetectFaceByBase64IntlResponse", string(data)}, " ")
}
