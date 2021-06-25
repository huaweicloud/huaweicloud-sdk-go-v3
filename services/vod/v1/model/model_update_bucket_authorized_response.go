package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateBucketAuthorizedResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateBucketAuthorizedResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateBucketAuthorizedResponse struct{}"
	}

	return strings.Join([]string{"UpdateBucketAuthorizedResponse", string(data)}, " ")
}
