package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteP2cVgwConnectionRequest Request Object
type DeleteP2cVgwConnectionRequest struct {

	// P2C VPN网关实例ID
	P2cVgwId string `json:"p2c_vgw_id"`

	// 连接ID
	ConnectionId string `json:"connection_id"`
}

func (o DeleteP2cVgwConnectionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteP2cVgwConnectionRequest struct{}"
	}

	return strings.Join([]string{"DeleteP2cVgwConnectionRequest", string(data)}, " ")
}
