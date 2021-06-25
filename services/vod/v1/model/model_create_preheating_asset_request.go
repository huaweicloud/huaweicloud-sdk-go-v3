package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreatePreheatingAssetRequest struct {
	Body *CreatePreheatingAssetReq `json:"body,omitempty"`
}

func (o CreatePreheatingAssetRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreatePreheatingAssetRequest struct{}"
	}

	return strings.Join([]string{"CreatePreheatingAssetRequest", string(data)}, " ")
}
