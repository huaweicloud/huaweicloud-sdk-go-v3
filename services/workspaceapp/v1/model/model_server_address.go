package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ServerAddress struct {

	// IP地址信息
	Addr *string `json:"addr,omitempty"`

	// IP地址类型， `4` - IPV4 `6` - IPV6
	Version *int32 `json:"version,omitempty"`

	// MAC地址
	OSEXTIPSMACmacAddr *string `json:"OS-EXT-IPS-MAC:mac_addr,omitempty"`

	// IP地址分配方式，字符串是大小写不敏感格式 * `fixed` - 代表私有IP地址 * `floating` - 代表浮动IP地址
	OSEXTIPStype *string `json:"OS-EXT-IPS:type,omitempty"`

	// IP地址对应的端口ID。
	OSEXTIPSportId *string `json:"OS-EXT-IPS:port_id,omitempty"`
}

func (o ServerAddress) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServerAddress struct{}"
	}

	return strings.Join([]string{"ServerAddress", string(data)}, " ")
}
