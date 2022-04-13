package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListQuotaCouponsRequest struct {
	Body *QueryCouponQuotasReqExt `json:"body,omitempty"`
}

func (o ListQuotaCouponsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListQuotaCouponsRequest struct{}"
	}

	return strings.Join([]string{"ListQuotaCouponsRequest", string(data)}, " ")
}
