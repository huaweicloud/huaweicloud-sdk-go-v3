package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HostIp **参数解释**: 服务器IP **取值范围**: 字符长度1-128位
type HostIp struct {
}

func (o HostIp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HostIp struct{}"
	}

	return strings.Join([]string{"HostIp", string(data)}, " ")
}
