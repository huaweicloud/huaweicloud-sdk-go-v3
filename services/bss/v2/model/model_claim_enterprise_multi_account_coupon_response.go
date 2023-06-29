package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ClaimEnterpriseMultiAccountCouponResponse Response Object
type ClaimEnterpriseMultiAccountCouponResponse struct {

	// |参数名称：实际回收金额。| |参数的约束及描述：成功时返回|
	SubCouponId    *string `json:"sub_coupon_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ClaimEnterpriseMultiAccountCouponResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClaimEnterpriseMultiAccountCouponResponse struct{}"
	}

	return strings.Join([]string{"ClaimEnterpriseMultiAccountCouponResponse", string(data)}, " ")
}
