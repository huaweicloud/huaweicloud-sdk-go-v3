package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 主机端口范围，在范围内为应用实例自动分配主机端口；与主机端口参数二选一；仅铂金版可用
type HostPortRange struct {
	// 主机端口下限制,1到65535之间的整数

	MinPort int32 `json:"min_port"`
	// 主机端口上限值,1到65535之间的整数；max_port需大于min_port

	MaxPort int32 `json:"max_port"`
}

func (o HostPortRange) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HostPortRange struct{}"
	}

	return strings.Join([]string{"HostPortRange", string(data)}, " ")
}
