package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowDefaultConfigRequest struct {
}

func (o ShowDefaultConfigRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowDefaultConfigRequest struct{}"
	}

	return strings.Join([]string{"ShowDefaultConfigRequest", string(data)}, " ")
}
