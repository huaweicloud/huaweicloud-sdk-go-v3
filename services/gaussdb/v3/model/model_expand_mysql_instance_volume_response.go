package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ExpandMysqlInstanceVolumeResponse struct {
	// 扩容后容量。

	Size *int32 `json:"size,omitempty"`
	// 订单号，创建包年包月时返回该参数。

	OrderId        *string `json:"order_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExpandMysqlInstanceVolumeResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ExpandMysqlInstanceVolumeResponse struct{}"
	}

	return strings.Join([]string{"ExpandMysqlInstanceVolumeResponse", string(data)}, " ")
}
