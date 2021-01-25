package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListTagResponse struct {
	Resources      *[]ResourceItem `json:"resources,omitempty"`
	TotalCount     *int32          `json:"total_count,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListTagResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListTagResponse struct{}"
	}

	return strings.Join([]string{"ListTagResponse", string(data)}, " ")
}
