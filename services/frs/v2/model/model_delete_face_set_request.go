package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteFaceSetRequest struct {
	// 人脸库名称。

	FaceSetName string `json:"face_set_name"`
}

func (o DeleteFaceSetRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteFaceSetRequest struct{}"
	}

	return strings.Join([]string{"DeleteFaceSetRequest", string(data)}, " ")
}
