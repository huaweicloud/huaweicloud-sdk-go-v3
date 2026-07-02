package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListConnectionIpsecSaRequest Request Object
type ListConnectionIpsecSaRequest struct {

	// vpn连接ID
	VpnConnectionId string `json:"vpn_connection_id"`
}

func (o ListConnectionIpsecSaRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConnectionIpsecSaRequest struct{}"
	}

	return strings.Join([]string{"ListConnectionIpsecSaRequest", string(data)}, " ")
}
