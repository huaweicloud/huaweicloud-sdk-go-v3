package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteFaceByExternalImageIdResponse struct {
	// 删除的人脸数量。 调用失败时无此字段。

	FaceNumber *int32 `json:"face_number,omitempty"`
	// 人脸库ID。 调用失败时无此字段。

	FaceSetId *string `json:"face_set_id,omitempty"`
	// 人脸库名称。 调用失败时无此字段。

	FaceSetName    *string `json:"face_set_name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteFaceByExternalImageIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteFaceByExternalImageIdResponse struct{}"
	}

	return strings.Join([]string{"DeleteFaceByExternalImageIdResponse", string(data)}, " ")
}
