package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePartnerCouponsResponse Response Object
type CreatePartnerCouponsResponse struct {

	// 错误的客户列表和错误信息，只有HTTP 200的时候才会返回这个结构体，具体参见表1。
	ErrorDetails *[]ErrorDetail `json:"error_details,omitempty"`

	// 成功的客户ID和对应的券ID列表，只有HTTP 200的时候才会返回这个结构体，具体参见表2。
	CouponInfos    *[]CouponSimpleInfo `json:"coupon_infos,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o CreatePartnerCouponsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePartnerCouponsResponse struct{}"
	}

	return strings.Join([]string{"CreatePartnerCouponsResponse", string(data)}, " ")
}
