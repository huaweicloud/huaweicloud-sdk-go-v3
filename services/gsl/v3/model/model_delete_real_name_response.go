package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteRealNameResponse struct {
	// 业务受理单号

	WorkOrderId    *int64 `json:"work_order_id,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o DeleteRealNameResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteRealNameResponse struct{}"
	}

	return strings.Join([]string{"DeleteRealNameResponse", string(data)}, " ")
}
