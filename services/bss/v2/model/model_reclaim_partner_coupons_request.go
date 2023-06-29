package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ReclaimPartnerCouponsRequest Request Object
type ReclaimPartnerCouponsRequest struct {
	Body *ReclaimPartnerCouponsReq `json:"body,omitempty"`
}

func (o ReclaimPartnerCouponsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReclaimPartnerCouponsRequest struct{}"
	}

	return strings.Join([]string{"ReclaimPartnerCouponsRequest", string(data)}, " ")
}
