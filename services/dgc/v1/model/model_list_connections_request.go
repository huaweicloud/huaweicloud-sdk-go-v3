package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListConnectionsRequest struct {
}

func (o ListConnectionsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListConnectionsRequest struct{}"
	}

	return strings.Join([]string{"ListConnectionsRequest", string(data)}, " ")
}
