package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type RenewalResourcesResponse struct {
	// 续订资源生成的订单ID的列表。
	OrderIds       *[]string `json:"order_ids,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o RenewalResourcesResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RenewalResourcesResponse struct{}"
	}

	return strings.Join([]string{"RenewalResourcesResponse", string(data)}, " ")
}
