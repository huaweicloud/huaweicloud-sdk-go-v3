package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowFaceSetRequest struct {
	// 人脸库名称。

	FaceSetName string `json:"face_set_name"`
}

func (o ShowFaceSetRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowFaceSetRequest struct{}"
	}

	return strings.Join([]string{"ShowFaceSetRequest", string(data)}, " ")
}
