package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteFaceByFaceIdResponse struct {
	// 删除的人脸数量。 调用失败时无此字段。

	FaceNumber *int32 `json:"face_number,omitempty"`
	// 人脸库ID。 调用失败时无此字段。

	FaceSetId *string `json:"face_set_id,omitempty"`
	// 人脸库名称。 调用失败时无此字段。

	FaceSetName    *string `json:"face_set_name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteFaceByFaceIdResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteFaceByFaceIdResponse struct{}"
	}

	return strings.Join([]string{"DeleteFaceByFaceIdResponse", string(data)}, " ")
}
