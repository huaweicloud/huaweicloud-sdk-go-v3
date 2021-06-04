package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListApiVersionRequest struct {
}

func (o ListApiVersionRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListApiVersionRequest struct{}"
	}

	return strings.Join([]string{"ListApiVersionRequest", string(data)}, " ")
}
