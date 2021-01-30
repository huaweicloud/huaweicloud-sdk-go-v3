package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListAuditlogLinksRequest struct {
	InstanceId string                           `json:"instance_id"`
	Body       *ProduceAuditlogLinksRequestBody `json:"body,omitempty"`
}

func (o ListAuditlogLinksRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListAuditlogLinksRequest struct{}"
	}

	return strings.Join([]string{"ListAuditlogLinksRequest", string(data)}, " ")
}
