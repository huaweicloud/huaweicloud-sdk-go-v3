package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 端口监听范围（闭区间)，最多指定20个端口组，每个组范围不可有重叠部分 >仅当protocol_port为0时可以传入。
type PortRange struct {

	// 起始端口
	StartPort *int32 `json:"start_port,omitempty"`

	// 结束端口，需大于等于起始端口
	EndPort *int32 `json:"end_port,omitempty"`
}

func (o PortRange) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PortRange struct{}"
	}

	return strings.Join([]string{"PortRange", string(data)}, " ")
}
