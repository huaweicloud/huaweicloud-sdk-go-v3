package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateCouponQuotasRequest Request Object
type UpdateCouponQuotasRequest struct {
	Body *AdjustCouponQuotasReq `json:"body,omitempty"`
}

func (o UpdateCouponQuotasRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCouponQuotasRequest struct{}"
	}

	return strings.Join([]string{"UpdateCouponQuotasRequest", string(data)}, " ")
}
