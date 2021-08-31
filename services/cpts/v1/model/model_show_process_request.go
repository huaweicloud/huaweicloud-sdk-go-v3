package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowProcessRequest struct {
}

func (o ShowProcessRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowProcessRequest struct{}"
	}

	return strings.Join([]string{"ShowProcessRequest", string(data)}, " ")
}
