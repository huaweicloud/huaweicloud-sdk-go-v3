package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowComponentDetailRequest struct {
	ApplicationId string `json:"application_id"`
	ComponentId   string `json:"component_id"`
}

func (o ShowComponentDetailRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowComponentDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowComponentDetailRequest", string(data)}, " ")
}
