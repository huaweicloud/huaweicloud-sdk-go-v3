package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListKeypairsResponse struct {
	Body           *[]SimpleKeypair `json:"body,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListKeypairsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListKeypairsResponse struct{}"
	}

	return strings.Join([]string{"ListKeypairsResponse", string(data)}, " ")
}
