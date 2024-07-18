package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowP2cVgwRequest Request Object
type ShowP2cVgwRequest struct {

	// P2C VPN网关实例ID
	P2cVgwId string `json:"p2c_vgw_id"`
}

func (o ShowP2cVgwRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowP2cVgwRequest struct{}"
	}

	return strings.Join([]string{"ShowP2cVgwRequest", string(data)}, " ")
}
