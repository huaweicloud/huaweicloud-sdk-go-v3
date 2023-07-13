package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ReclaimCouponQuotasRequest Request Object
type ReclaimCouponQuotasRequest struct {
	Body *ReclaimCouponQuotasReq `json:"body,omitempty"`
}

func (o ReclaimCouponQuotasRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReclaimCouponQuotasRequest struct{}"
	}

	return strings.Join([]string{"ReclaimCouponQuotasRequest", string(data)}, " ")
}
