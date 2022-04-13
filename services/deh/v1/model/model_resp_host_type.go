package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 可用的专属主机类型。
type RespHostType struct {
	// 专属主机类型。

	HostType string `json:"host_type"`
	// 专属主机类型名字。

	HostTypeName string `json:"host_type_name"`
}

func (o RespHostType) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RespHostType struct{}"
	}

	return strings.Join([]string{"RespHostType", string(data)}, " ")
}
