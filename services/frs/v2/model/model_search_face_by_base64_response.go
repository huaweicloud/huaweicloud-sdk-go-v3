package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type SearchFaceByBase64Response struct {

	// [查找的人脸集合，详见[SearchFace](https://support.huaweicloud.com/api-face/face_02_0019.html)。调用失败时无此字段。](tag:hc) [查找的人脸集合，详见[SearchFace](https://support.huaweicloud.com/intl/zh-cn/api-face/face_02_0019.html)。调用失败时无此字段。](tag:hk)
	Faces          *[]SearchFace `json:"faces,omitempty" xml:"faces"`
	HttpStatusCode int           `json:"-"`
}

func (o SearchFaceByBase64Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchFaceByBase64Response struct{}"
	}

	return strings.Join([]string{"SearchFaceByBase64Response", string(data)}, " ")
}
