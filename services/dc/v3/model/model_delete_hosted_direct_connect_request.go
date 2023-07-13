package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteHostedDirectConnectRequest Request Object
type DeleteHostedDirectConnectRequest struct {

	// 托管专线连接ID。
	HostedConnectId string `json:"hosted_connect_id"`
}

func (o DeleteHostedDirectConnectRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteHostedDirectConnectRequest struct{}"
	}

	return strings.Join([]string{"DeleteHostedDirectConnectRequest", string(data)}, " ")
}
