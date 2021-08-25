package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowSubscribesRequest struct {
}

func (o ShowSubscribesRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowSubscribesRequest struct{}"
	}

	return strings.Join([]string{"ShowSubscribesRequest", string(data)}, " ")
}
