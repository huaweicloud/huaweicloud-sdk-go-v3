package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListQuotaCouponsRequest Request Object
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
