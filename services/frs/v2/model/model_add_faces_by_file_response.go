package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type AddFacesByFileResponse struct {
	// 人脸库ID。 调用失败时无此字段。

	FaceSetId *string `json:"face_set_id,omitempty"`
	// 人脸库名称。 调用失败时无此字段。

	FaceSetName *string `json:"face_set_name,omitempty"`
	// 人脸库当中的人脸结构，详见[FaceSetFace](https://support.huaweicloud.com/api-face/face_02_0018.html)。 调用失败时无此字段。

	Faces          *[]FaceSetFace `json:"faces,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o AddFacesByFileResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddFacesByFileResponse struct{}"
	}

	return strings.Join([]string{"AddFacesByFileResponse", string(data)}, " ")
}
