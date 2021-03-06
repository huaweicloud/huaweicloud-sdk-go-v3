package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListAuthorizationsRequest struct {
}

func (o ListAuthorizationsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListAuthorizationsRequest struct{}"
	}

	return strings.Join([]string{"ListAuthorizationsRequest", string(data)}, " ")
}
