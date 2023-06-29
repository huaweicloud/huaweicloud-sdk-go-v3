package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ReclaimToPartnerAccountRequest Request Object
type ReclaimToPartnerAccountRequest struct {
	Body *ReclaimToPartnerAccountBalancesReq `json:"body,omitempty"`
}

func (o ReclaimToPartnerAccountRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReclaimToPartnerAccountRequest struct{}"
	}

	return strings.Join([]string{"ReclaimToPartnerAccountRequest", string(data)}, " ")
}
