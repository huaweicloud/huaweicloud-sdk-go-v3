package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 在线会场信息。
type RealTimeConfInfo struct {
	// 主持人与会者标识。

	ChairID *string `json:"chairID,omitempty"`
}

func (o RealTimeConfInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RealTimeConfInfo struct{}"
	}

	return strings.Join([]string{"RealTimeConfInfo", string(data)}, " ")
}
