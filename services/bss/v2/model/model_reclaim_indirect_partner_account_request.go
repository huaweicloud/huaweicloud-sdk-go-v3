package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ReclaimIndirectPartnerAccountRequest Request Object
type ReclaimIndirectPartnerAccountRequest struct {
	Body *ReclaimIndirectPartnerAccountReq `json:"body,omitempty"`
}

func (o ReclaimIndirectPartnerAccountRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReclaimIndirectPartnerAccountRequest struct{}"
	}

	return strings.Join([]string{"ReclaimIndirectPartnerAccountRequest", string(data)}, " ")
}
