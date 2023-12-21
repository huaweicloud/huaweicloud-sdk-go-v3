package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowOmUrlRequest Request Object
type ShowOmUrlRequest struct {

	// 云堡垒机服务器ID
	ServerId string `json:"server_id"`

	// 被纳管主机IP
	IpAddress string `json:"ip_address"`

	// 被纳管主机的账户
	HostAccountName string `json:"host_account_name"`
}

func (o ShowOmUrlRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowOmUrlRequest struct{}"
	}

	return strings.Join([]string{"ShowOmUrlRequest", string(data)}, " ")
}
