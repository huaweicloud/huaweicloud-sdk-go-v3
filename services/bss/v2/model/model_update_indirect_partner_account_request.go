package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateIndirectPartnerAccountRequest Request Object
type UpdateIndirectPartnerAccountRequest struct {
	Body *AdjustToIndirectPartnerReq `json:"body,omitempty"`
}

func (o UpdateIndirectPartnerAccountRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateIndirectPartnerAccountRequest struct{}"
	}

	return strings.Join([]string{"UpdateIndirectPartnerAccountRequest", string(data)}, " ")
}
