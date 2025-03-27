package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteVpnConnectionsLogConfigResponse Response Object
type DeleteVpnConnectionsLogConfigResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteVpnConnectionsLogConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteVpnConnectionsLogConfigResponse struct{}"
	}

	return strings.Join([]string{"DeleteVpnConnectionsLogConfigResponse", string(data)}, " ")
}
