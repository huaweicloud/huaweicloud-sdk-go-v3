package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateAssetRequest struct {
	Body *UpdateAssetReq `json:"body,omitempty"`
}

func (o UpdateAssetRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateAssetRequest struct{}"
	}

	return strings.Join([]string{"UpdateAssetRequest", string(data)}, " ")
}
