package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateCouponQuotasResponse Response Object
type UpdateCouponQuotasResponse struct {

	// 发放失败的云经销商额度信息，具体参见表1，只有HTTP STATUS 200的时候才有这个结构体。
	ErrorDetails *[]ErrorDetail `json:"error_details,omitempty"`

	// 发放成功的云经销商额度信息，具体参见表2，只有HTTP STATUS 200的时候才有这个结构体。
	SimpleQuotaInfos *[]QuotaSimpleInfo `json:"simple_quota_infos,omitempty"`
	HttpStatusCode   int                `json:"-"`
}

func (o UpdateCouponQuotasResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCouponQuotasResponse struct{}"
	}

	return strings.Join([]string{"UpdateCouponQuotasResponse", string(data)}, " ")
}
