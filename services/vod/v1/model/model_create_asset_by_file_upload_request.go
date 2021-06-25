package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateAssetByFileUploadRequest struct {
	Body *CreateAssetByFileReq `json:"body,omitempty"`
}

func (o CreateAssetByFileUploadRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateAssetByFileUploadRequest struct{}"
	}

	return strings.Join([]string{"CreateAssetByFileUploadRequest", string(data)}, " ")
}
