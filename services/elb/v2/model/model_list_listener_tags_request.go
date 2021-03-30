package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListListenerTagsRequest struct {
}

func (o ListListenerTagsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListListenerTagsRequest struct{}"
	}

	return strings.Join([]string{"ListListenerTagsRequest", string(data)}, " ")
}
