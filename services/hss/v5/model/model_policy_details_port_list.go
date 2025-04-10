package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PolicyDetailsPortList 端口
type PolicyDetailsPortList struct {

	// 端口
	Port *int32 `json:"port,omitempty"`

	// 协议类型 - \"tcp\" - \"tcp6\"
	Protocol *string `json:"protocol,omitempty"`
}

func (o PolicyDetailsPortList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyDetailsPortList struct{}"
	}

	return strings.Join([]string{"PolicyDetailsPortList", string(data)}, " ")
}
