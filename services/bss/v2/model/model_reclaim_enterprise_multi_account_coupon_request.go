package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ReclaimEnterpriseMultiAccountCouponRequest Request Object
type ReclaimEnterpriseMultiAccountCouponRequest struct {
	Body *RetrieveEnterpriseMultiAccountCouponReq `json:"body,omitempty"`
}

func (o ReclaimEnterpriseMultiAccountCouponRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReclaimEnterpriseMultiAccountCouponRequest struct{}"
	}

	return strings.Join([]string{"ReclaimEnterpriseMultiAccountCouponRequest", string(data)}, " ")
}
