package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowFaceSetResponse struct {
	FaceSetInfo    *FaceSetInfo `json:"face_set_info,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowFaceSetResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowFaceSetResponse struct{}"
	}

	return strings.Join([]string{"ShowFaceSetResponse", string(data)}, " ")
}
