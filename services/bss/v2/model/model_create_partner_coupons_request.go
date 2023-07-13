package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePartnerCouponsRequest Request Object
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
