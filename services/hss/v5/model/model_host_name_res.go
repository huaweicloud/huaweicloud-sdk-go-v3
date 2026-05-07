package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HostNameRes **参数解释**: 服务器名称 **取值范围**: 字符长度1-256位
type HostNameRes struct {
}

func (o HostNameRes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HostNameRes struct{}"
	}

	return strings.Join([]string{"HostNameRes", string(data)}, " ")
}
