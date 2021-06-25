package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DisableLogCollectionRequest struct {
}

func (o DisableLogCollectionRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DisableLogCollectionRequest struct{}"
	}

	return strings.Join([]string{"DisableLogCollectionRequest", string(data)}, " ")
}
