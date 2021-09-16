package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type RegisterImeiResponse struct {
	// 业务受理单号

	WorkOrderId    *int64 `json:"work_order_id,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o RegisterImeiResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RegisterImeiResponse struct{}"
	}

	return strings.Join([]string{"RegisterImeiResponse", string(data)}, " ")
}
