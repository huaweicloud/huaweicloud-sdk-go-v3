package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateCouponQuotasResponse struct {
	// 发放失败的精英服务商额度信息，具体参见表1，只有HTTP STATUS 200的时候才有这个结构体。

	ErrorDetails *[]ErrorDetail `json:"error_details,omitempty"`
	// 发放成功的精英服务商额度信息，具体参见表2，只有HTTP STATUS 200的时候才有这个结构体。

	SimpleQuotaInfos *[]QuotaSimpleInfo `json:"simple_quota_infos,omitempty"`
	HttpStatusCode   int                `json:"-"`
}

func (o UpdateCouponQuotasResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateCouponQuotasResponse struct{}"
	}

	return strings.Join([]string{"UpdateCouponQuotasResponse", string(data)}, " ")
}
