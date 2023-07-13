package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ReclaimEnterpriseMultiAccountCouponResponse Response Object
type ReclaimEnterpriseMultiAccountCouponResponse struct {

	// |参数名称：实际回收金额。| |参数的约束及描述：成功时返回|
	RealRetrieveAmout *string `json:"real_retrieve_amout,omitempty"`
	HttpStatusCode    int     `json:"-"`
}

func (o ReclaimEnterpriseMultiAccountCouponResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReclaimEnterpriseMultiAccountCouponResponse struct{}"
	}

	return strings.Join([]string{"ReclaimEnterpriseMultiAccountCouponResponse", string(data)}, " ")
}
