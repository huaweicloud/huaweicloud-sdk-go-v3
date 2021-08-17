package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowOverviewRequest struct {
}

func (o ShowOverviewRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowOverviewRequest struct{}"
	}

	return strings.Join([]string{"ShowOverviewRequest", string(data)}, " ")
}
