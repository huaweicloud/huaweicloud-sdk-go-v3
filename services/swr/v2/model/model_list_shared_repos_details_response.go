package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListSharedReposDetailsResponse struct {
	Body *[]ShowReposResp `json:"body,omitempty"`

	ContentRange   *string `json:"Content-Range,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListSharedReposDetailsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListSharedReposDetailsResponse struct{}"
	}

	return strings.Join([]string{"ListSharedReposDetailsResponse", string(data)}, " ")
}
