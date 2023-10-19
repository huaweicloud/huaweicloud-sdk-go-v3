package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFacesByLimitResponse Response Object
type ShowFacesByLimitResponse struct {

	// 人脸库ID，随机生成的包含八个字符的字符串。 调用失败时无此字段。
	FaceSetId *string `json:"face_set_id,omitempty"`

	// 人脸库名称。 调用失败时无此字段。
	FaceSetName *string `json:"face_set_name,omitempty"`

	// [人脸库当中的人脸结构，详见[FaceSetFace](https://support.huaweicloud.com/api-face/face_02_0018.html)。调用失败时无此字段。](tag:hc) [人脸库当中的人脸结构，详见[FaceSetFace](https://support.huaweicloud.com/intl/zh-cn/api-face/face_02_0018.html)。调用失败时无此字段。](tag:hk)
	Faces *[]FaceSetFace `json:"faces,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowFacesByLimitResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFacesByLimitResponse struct{}"
	}

	return strings.Join([]string{"ShowFacesByLimitResponse", string(data)}, " ")
}
