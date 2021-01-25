package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListResourcesRequest struct {
}

func (o ListResourcesRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListResourcesRequest struct{}"
	}

	return strings.Join([]string{"ListResourcesRequest", string(data)}, " ")
}
