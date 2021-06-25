package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type EnableLogCollectionRequest struct {
}

func (o EnableLogCollectionRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "EnableLogCollectionRequest struct{}"
	}

	return strings.Join([]string{"EnableLogCollectionRequest", string(data)}, " ")
}
