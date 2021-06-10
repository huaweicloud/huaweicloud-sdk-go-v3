package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowAccessoryLimitsResponse struct {
	AccessoryLimit *AccessoryLimitVo `json:"accessory_limit,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ShowAccessoryLimitsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowAccessoryLimitsResponse struct{}"
	}

	return strings.Join([]string{"ShowAccessoryLimitsResponse", string(data)}, " ")
}
