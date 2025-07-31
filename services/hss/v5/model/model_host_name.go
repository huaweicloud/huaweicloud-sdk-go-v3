package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HostName **参数解释**: 服务器名称 **取值范围**: 字符长度1-256位
type HostName struct {
}

func (o HostName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HostName struct{}"
	}

	return strings.Join([]string{"HostName", string(data)}, " ")
}
