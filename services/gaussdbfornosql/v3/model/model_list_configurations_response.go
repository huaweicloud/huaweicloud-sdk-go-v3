package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListConfigurationsResponse struct {
	// 总记录数。

	Count *int32 `json:"count,omitempty"`

	Configurations *[]ListConfigurationsResult `json:"configurations,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o ListConfigurationsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListConfigurationsResponse struct{}"
	}

	return strings.Join([]string{"ListConfigurationsResponse", string(data)}, " ")
}
