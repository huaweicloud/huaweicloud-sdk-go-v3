package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListConfigurationsRequest struct {
}

func (o ListConfigurationsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListConfigurationsRequest struct{}"
	}

	return strings.Join([]string{"ListConfigurationsRequest", string(data)}, " ")
}
