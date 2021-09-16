package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type EnableSimCardResponse struct {
	// 业务受理单号

	WorkOrderId    *int64 `json:"work_order_id,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o EnableSimCardResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "EnableSimCardResponse struct{}"
	}

	return strings.Join([]string{"EnableSimCardResponse", string(data)}, " ")
}
