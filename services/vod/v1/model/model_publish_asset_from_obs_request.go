package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type PublishAssetFromObsRequest struct {
	Body *PublishAssetFromObsReq `json:"body,omitempty"`
}

func (o PublishAssetFromObsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "PublishAssetFromObsRequest struct{}"
	}

	return strings.Join([]string{"PublishAssetFromObsRequest", string(data)}, " ")
}
