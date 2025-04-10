package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HoneypotPortPolicyContentPortsList 端口
type HoneypotPortPolicyContentPortsList struct {

	// 端口
	Port int32 `json:"port"`

	// 协议类型： - \"tcp\" - \"tcp6\"
	Protocol string `json:"protocol"`
}

func (o HoneypotPortPolicyContentPortsList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HoneypotPortPolicyContentPortsList struct{}"
	}

	return strings.Join([]string{"HoneypotPortPolicyContentPortsList", string(data)}, " ")
}
