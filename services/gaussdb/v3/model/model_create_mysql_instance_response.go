package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateMysqlInstanceResponse struct {
	Instance *MysqlInstanceResponse `json:"instance,omitempty"`
	// 实例创建的任务id。  仅创建按需实例时会返回该参数。

	JobId *string `json:"job_id,omitempty"`
	// 订单号，创建包年包月时返回该参数。

	OrderId        *string `json:"order_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateMysqlInstanceResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateMysqlInstanceResponse struct{}"
	}

	return strings.Join([]string{"CreateMysqlInstanceResponse", string(data)}, " ")
}
