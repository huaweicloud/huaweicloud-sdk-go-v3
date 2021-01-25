package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeletePredefineTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeletePredefineTagsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeletePredefineTagsResponse struct{}"
	}

	return strings.Join([]string{"DeletePredefineTagsResponse", string(data)}, " ")
}
