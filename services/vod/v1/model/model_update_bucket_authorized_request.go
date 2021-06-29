package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateBucketAuthorizedRequest struct {
	Body *UpdateBucketAuthorizedReq `json:"body,omitempty"`
}

func (o UpdateBucketAuthorizedRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateBucketAuthorizedRequest struct{}"
	}

	return strings.Join([]string{"UpdateBucketAuthorizedRequest", string(data)}, " ")
}
