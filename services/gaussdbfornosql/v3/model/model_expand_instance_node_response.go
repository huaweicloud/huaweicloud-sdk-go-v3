package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ExpandInstanceNodeResponse struct {
	// 任务ID，仅按需实例时会返回该参数。

	JobId *string `json:"job_id,omitempty"`
	// 订单ID，仅创建包年包月实例时返回该参数。

	OrderId        *string `json:"order_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExpandInstanceNodeResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ExpandInstanceNodeResponse struct{}"
	}

	return strings.Join([]string{"ExpandInstanceNodeResponse", string(data)}, " ")
}
