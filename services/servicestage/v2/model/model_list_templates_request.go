package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListTemplatesRequest struct {
}

func (o ListTemplatesRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListTemplatesRequest struct{}"
	}

	return strings.Join([]string{"ListTemplatesRequest", string(data)}, " ")
}
