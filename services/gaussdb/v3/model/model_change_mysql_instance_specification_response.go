package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ChangeMysqlInstanceSpecificationResponse struct {
	// 规格变更的任务id，仅变更按需实例时会返回该参数

	JobId *string `json:"job_id,omitempty"`
	// 订单id，仅变更包周期实例时会返回该参数

	OrderId        *string `json:"order_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ChangeMysqlInstanceSpecificationResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ChangeMysqlInstanceSpecificationResponse struct{}"
	}

	return strings.Join([]string{"ChangeMysqlInstanceSpecificationResponse", string(data)}, " ")
}
