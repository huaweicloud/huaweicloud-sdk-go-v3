package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateP2cVgwRequest Request Object
type UpdateP2cVgwRequest struct {

	// P2C VPN网关实例ID
	P2cVgwId string `json:"p2c_vgw_id"`

	Body *UpdateP2cVgwRequestBody `json:"body,omitempty"`
}

func (o UpdateP2cVgwRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateP2cVgwRequest struct{}"
	}

	return strings.Join([]string{"UpdateP2cVgwRequest", string(data)}, " ")
}
