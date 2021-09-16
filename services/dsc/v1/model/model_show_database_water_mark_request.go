package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowDatabaseWaterMarkRequest struct {
	Body *ExtractedDatabaseWatermark `json:"body,omitempty"`
}

func (o ShowDatabaseWaterMarkRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowDatabaseWaterMarkRequest struct{}"
	}

	return strings.Join([]string{"ShowDatabaseWaterMarkRequest", string(data)}, " ")
}
