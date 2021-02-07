package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListReposDetailsResponse struct {
	Body           *[]ShowReposResp `json:"body,omitempty"`
	ContentRange   *string          `json:"Content-Range,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListReposDetailsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListReposDetailsResponse struct{}"
	}

	return strings.Join([]string{"ListReposDetailsResponse", string(data)}, " ")
}
