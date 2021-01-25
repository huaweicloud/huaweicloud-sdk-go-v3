package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowPriceRequest struct {
}

func (o ShowPriceRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowPriceRequest struct{}"
	}

	return strings.Join([]string{"ShowPriceRequest", string(data)}, " ")
}
