package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteVpnConnectionResponse Response Object
type DeleteVpnConnectionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteVpnConnectionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteVpnConnectionResponse struct{}"
	}

	return strings.Join([]string{"DeleteVpnConnectionResponse", string(data)}, " ")
}
