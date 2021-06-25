package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateAssetMetaRequest struct {
	Body *UpdateAssetMetaReq `json:"body,omitempty"`
}

func (o UpdateAssetMetaRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateAssetMetaRequest struct{}"
	}

	return strings.Join([]string{"UpdateAssetMetaRequest", string(data)}, " ")
}
