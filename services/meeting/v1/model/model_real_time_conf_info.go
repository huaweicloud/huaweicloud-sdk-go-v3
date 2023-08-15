package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RealTimeConfInfo 在线会议信息。
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
