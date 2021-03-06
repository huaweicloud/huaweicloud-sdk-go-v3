package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListConversionsResponse struct {
	// 度量单位的换算信息，具体参见表2。

	Conversions    *[]Conversion `json:"conversions,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ListConversionsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListConversionsResponse struct{}"
	}

	return strings.Join([]string{"ListConversionsResponse", string(data)}, " ")
}
