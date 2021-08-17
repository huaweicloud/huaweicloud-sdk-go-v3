package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListNewConfigsRequest struct {
}

func (o ListNewConfigsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListNewConfigsRequest struct{}"
	}

	return strings.Join([]string{"ListNewConfigsRequest", string(data)}, " ")
}
