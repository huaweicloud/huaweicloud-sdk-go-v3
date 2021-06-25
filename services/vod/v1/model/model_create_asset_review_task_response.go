package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateAssetReviewTaskResponse struct {
	// 媒资ID

	AssetId *string `json:"asset_id,omitempty"`

	Review         *Review `json:"review,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateAssetReviewTaskResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateAssetReviewTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateAssetReviewTaskResponse", string(data)}, " ")
}
