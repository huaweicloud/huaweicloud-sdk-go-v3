package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateFaceSetResponse struct {
	FaceSetInfo    *FaceSetInfo `json:"face_set_info,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o CreateFaceSetResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateFaceSetResponse struct{}"
	}

	return strings.Join([]string{"CreateFaceSetResponse", string(data)}, " ")
}
