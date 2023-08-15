package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RespAddr 云服务器的vpc信息。
type RespAddr struct {

	// 云服务器的vpc ip。
	Addr string `json:"addr"`

	// 云服务器的vpc版本。
	Version int32 `json:"version"`

	// 扩展属性，分配IP地址方式。
	OSEXTIPStype *string `json:"OS-EXT-IPS:type,omitempty"`

	// 扩展属性，MAC地址。
	OSEXTIPSMACmacAddr *string `json:"OS-EXT-IPS-MAC:mac_addr,omitempty"`
}

func (o RespAddr) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RespAddr struct{}"
	}

	return strings.Join([]string{"RespAddr", string(data)}, " ")
}
