package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateFaceSetRequest struct {
	Body *CreateFaceSetReq `json:"body,omitempty"`
}

func (o CreateFaceSetRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateFaceSetRequest struct{}"
	}

	return strings.Join([]string{"CreateFaceSetRequest", string(data)}, " ")
}
