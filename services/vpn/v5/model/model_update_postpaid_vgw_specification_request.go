package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePostpaidVgwSpecificationRequest Request Object
type UpdatePostpaidVgwSpecificationRequest struct {

	// VPN网关实例ID
	VgwId string `json:"vgw_id"`

	Body *UpdateVgwSpecificationRequestBody `json:"body,omitempty"`
}

func (o UpdatePostpaidVgwSpecificationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePostpaidVgwSpecificationRequest struct{}"
	}

	return strings.Join([]string{"UpdatePostpaidVgwSpecificationRequest", string(data)}, " ")
}
