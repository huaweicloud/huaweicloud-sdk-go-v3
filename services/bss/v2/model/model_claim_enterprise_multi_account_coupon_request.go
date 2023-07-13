package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ClaimEnterpriseMultiAccountCouponRequest Request Object
type ClaimEnterpriseMultiAccountCouponRequest struct {
	Body *TransferEnterpriseMultiAccountCouponReq `json:"body,omitempty"`
}

func (o ClaimEnterpriseMultiAccountCouponRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClaimEnterpriseMultiAccountCouponRequest struct{}"
	}

	return strings.Join([]string{"ClaimEnterpriseMultiAccountCouponRequest", string(data)}, " ")
}
