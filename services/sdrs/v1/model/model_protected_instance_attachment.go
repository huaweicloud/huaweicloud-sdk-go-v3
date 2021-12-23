package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 保护实例挂载信息结构
type ProtectedInstanceAttachment struct {
	// 复制对ID。

	Replication string `json:"replication"`
	// 挂载点。

	Device string `json:"device"`
}

func (o ProtectedInstanceAttachment) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProtectedInstanceAttachment struct{}"
	}

	return strings.Join([]string{"ProtectedInstanceAttachment", string(data)}, " ")
}
