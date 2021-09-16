package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListSimPricePlansResponse struct {
	Body           *[]SimPricePlanVo `json:"body,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ListSimPricePlansResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListSimPricePlansResponse struct{}"
	}

	return strings.Join([]string{"ListSimPricePlansResponse", string(data)}, " ")
}
