package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PublicIpInfo struct {

	// 弹性公网IP配置ID。
	PublicipId *string `json:"publicip_id,omitempty"`

	// 弹性公网IP地址。
	PublicipAddress *string `json:"publicip_address,omitempty"`

	// IP版本信息。
	IpVersion *int32 `json:"ip_version,omitempty"`
}

func (o PublicIpInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublicIpInfo struct{}"
	}

	return strings.Join([]string{"PublicIpInfo", string(data)}, " ")
}
