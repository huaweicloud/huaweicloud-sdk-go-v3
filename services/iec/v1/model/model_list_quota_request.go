package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListQuotaRequest struct {
}

func (o ListQuotaRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListQuotaRequest struct{}"
	}

	return strings.Join([]string{"ListQuotaRequest", string(data)}, " ")
}
