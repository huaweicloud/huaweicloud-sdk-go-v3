package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteP2cVgwRequest Request Object
type DeleteP2cVgwRequest struct {

	// P2C VPN网关实例ID
	P2cVgwId string `json:"p2c_vgw_id"`
}

func (o DeleteP2cVgwRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteP2cVgwRequest struct{}"
	}

	return strings.Join([]string{"DeleteP2cVgwRequest", string(data)}, " ")
}
