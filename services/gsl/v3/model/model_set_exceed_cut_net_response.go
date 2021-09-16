package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type SetExceedCutNetResponse struct {
	// 业务受理单号

	WorkOrderId    *int64 `json:"work_order_id,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o SetExceedCutNetResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "SetExceedCutNetResponse struct{}"
	}

	return strings.Join([]string{"SetExceedCutNetResponse", string(data)}, " ")
}
