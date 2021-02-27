package model

import (
	"encoding/json"

	"strings"
)

type CustomerInfoV2 struct {
	// 客户账号ID。您可以调用查询客户列表接口获取customer_id。
	CustomerId string `json:"customer_id"`
}

func (o CustomerInfoV2) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CustomerInfoV2 struct{}"
	}

	return strings.Join([]string{"CustomerInfoV2", string(data)}, " ")
}
