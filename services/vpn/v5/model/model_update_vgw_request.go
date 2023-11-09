package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateVgwRequest Request Object
type UpdateVgwRequest struct {

	// VPN网关实例ID
	VgwId string `json:"vgw_id"`

	Body *UpdateVgwRequestBody `json:"body,omitempty"`
}

func (o UpdateVgwRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVgwRequest struct{}"
	}

	return strings.Join([]string{"UpdateVgwRequest", string(data)}, " ")
}
