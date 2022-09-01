package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteFaceByFaceIdResponse struct {

	// 删除的人脸数量。 调用失败时无此字段。
	FaceNumber *int32 `json:"face_number,omitempty" xml:"face_number"`

	// 人脸库ID。 调用失败时无此字段。
	FaceSetId *string `json:"face_set_id,omitempty" xml:"face_set_id"`

	// 人脸库名称。 调用失败时无此字段。
	FaceSetName    *string `json:"face_set_name,omitempty" xml:"face_set_name"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteFaceByFaceIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteFaceByFaceIdResponse struct{}"
	}

	return strings.Join([]string{"DeleteFaceByFaceIdResponse", string(data)}, " ")
}
