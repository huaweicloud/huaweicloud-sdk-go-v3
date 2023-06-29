package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ReclaimCouponQuotasResponse Response Object
type ReclaimCouponQuotasResponse struct {

	// 回收失败的云经销商额度信息，具体参见表1，只有HTTP STATUS 200的时候才有这个结构体。
	ErrorDetails *[]ErrorDetail `json:"error_details,omitempty"`

	// 回收成功的云经销商额度信息，具体参见表2，只有HTTP STATUS 200的时候才有这个结构体。
	SimpleQuotaInfos *[]QuotaReclaim `json:"simple_quota_infos,omitempty"`
	HttpStatusCode   int             `json:"-"`
}

func (o ReclaimCouponQuotasResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReclaimCouponQuotasResponse struct{}"
	}

	return strings.Join([]string{"ReclaimCouponQuotasResponse", string(data)}, " ")
}
