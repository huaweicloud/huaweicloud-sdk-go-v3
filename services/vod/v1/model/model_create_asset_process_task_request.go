package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateAssetProcessTaskRequest struct {
	Body *AssetProcessReq `json:"body,omitempty"`
}

func (o CreateAssetProcessTaskRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateAssetProcessTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateAssetProcessTaskRequest", string(data)}, " ")
}
