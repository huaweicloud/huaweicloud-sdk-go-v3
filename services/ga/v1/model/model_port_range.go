package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PortRange port range
type PortRange struct {

	// 起始端口。
	FromPort int32 `json:"from_port"`

	// 结束端口。
	ToPort int32 `json:"to_port"`
}

func (o PortRange) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PortRange struct{}"
	}

	return strings.Join([]string{"PortRange", string(data)}, " ")
}
