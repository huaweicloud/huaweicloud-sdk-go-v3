package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PublicIpRes **参数解释**: 服务器弹性IP地址 **取值范围**: IPv4格式（长度7-15位）、IPv6格式（长度15-39位）
type PublicIpRes struct {
}

func (o PublicIpRes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublicIpRes struct{}"
	}

	return strings.Join([]string{"PublicIpRes", string(data)}, " ")
}
