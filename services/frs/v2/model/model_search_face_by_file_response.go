package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchFaceByFileResponse Response Object
type SearchFaceByFileResponse struct {

	// [查找的人脸集合，详见[SearchFace](https://support.huaweicloud.com/api-face/face_02_0019.html)。调用失败时无此字段。](tag:hc) [查找的人脸集合，详见[SearchFace](https://support.huaweicloud.com/intl/zh-cn/api-face/face_02_0019.html)。调用失败时无此字段。](tag:hk)
	Faces *[]SearchFace `json:"faces,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SearchFaceByFileResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchFaceByFileResponse struct{}"
	}

	return strings.Join([]string{"SearchFaceByFileResponse", string(data)}, " ")
}
