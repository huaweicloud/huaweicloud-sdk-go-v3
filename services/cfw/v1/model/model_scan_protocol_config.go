package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ScanProtocolConfig struct {

	// 反病毒动作，0：观察 1：拦截 2：禁用
	Action *int32 `json:"action,omitempty"`

	// 协议类型
	ProtocolType *int32 `json:"protocol_type,omitempty"`
}

func (o ScanProtocolConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScanProtocolConfig struct{}"
	}

	return strings.Join([]string{"ScanProtocolConfig", string(data)}, " ")
}
