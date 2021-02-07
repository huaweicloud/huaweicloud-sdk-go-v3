package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateManualImageSyncRepoResponse struct {
	Body           *[]string `json:"body,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o CreateManualImageSyncRepoResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateManualImageSyncRepoResponse struct{}"
	}

	return strings.Join([]string{"CreateManualImageSyncRepoResponse", string(data)}, " ")
}
