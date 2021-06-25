package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateAssetMetaResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateAssetMetaResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateAssetMetaResponse struct{}"
	}

	return strings.Join([]string{"UpdateAssetMetaResponse", string(data)}, " ")
}
