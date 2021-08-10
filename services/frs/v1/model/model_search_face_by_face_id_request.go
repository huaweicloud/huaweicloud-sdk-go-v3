package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type SearchFaceByFaceIdRequest struct {
	// 人脸库名称。

	FaceSetName string `json:"face_set_name"`

	Body *FaceSearchFaceIdReq `json:"body,omitempty"`
}

func (o SearchFaceByFaceIdRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "SearchFaceByFaceIdRequest struct{}"
	}

	return strings.Join([]string{"SearchFaceByFaceIdRequest", string(data)}, " ")
}
