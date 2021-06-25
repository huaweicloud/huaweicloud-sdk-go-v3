package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateAssetReviewTaskRequest struct {
	Body *AssetReviewReq `json:"body,omitempty"`
}

func (o CreateAssetReviewTaskRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateAssetReviewTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateAssetReviewTaskRequest", string(data)}, " ")
}
