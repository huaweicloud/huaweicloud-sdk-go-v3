package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 弹性公网EIP信息
type PublicIpInfo struct {
	// 弹性公网ip配置id

	PublicipId string `json:"publicip_id"`
	// IP地址

	PublicipAddress string `json:"publicip_address"`
	// IP版本信息。 取值： - 4：IPv4 - 6：IPv6  [不支持IPv6，请勿设置为6。](tag:dt,dt_test)

	IpVersion int32 `json:"ip_version"`
}

func (o PublicIpInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublicIpInfo struct{}"
	}

	return strings.Join([]string{"PublicIpInfo", string(data)}, " ")
}
