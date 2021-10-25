package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListActiveActiveDomainsRequest struct {
}

func (o ListActiveActiveDomainsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListActiveActiveDomainsRequest struct{}"
	}

	return strings.Join([]string{"ListActiveActiveDomainsRequest", string(data)}, " ")
}
