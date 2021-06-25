package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UnpublishAssetsRequest struct {
	Body *PublishAssetReq `json:"body,omitempty"`
}

func (o UnpublishAssetsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UnpublishAssetsRequest struct{}"
	}

	return strings.Join([]string{"UnpublishAssetsRequest", string(data)}, " ")
}
