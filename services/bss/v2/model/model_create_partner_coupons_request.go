package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreatePartnerCouponsRequest struct {
	Body *CreatePartnerCouponsReq `json:"body,omitempty"`
}

func (o CreatePartnerCouponsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePartnerCouponsRequest struct{}"
	}

	return strings.Join([]string{"CreatePartnerCouponsRequest", string(data)}, " ")
}
