package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HostPortRange struct {

	// 主机端口上限值,1到65535之间的整数；max_port需大于min_port
	MaxPort int64 `json:"max_port"`

	// 主机端口下限制,1到65535之间的整数
	MinPort int64 `json:"min_port"`
}

func (o HostPortRange) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HostPortRange struct{}"
	}

	return strings.Join([]string{"HostPortRange", string(data)}, " ")
}
