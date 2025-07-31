package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PortStatisticResponseInfo **参数解释**: 开放端口统计信息
type PortStatisticResponseInfo struct {

	// **参数解释**: 端口号 **取值范围**: 最小值0，最大值65535
	Port *int32 `json:"port,omitempty"`

	// **参数解释**: 端口类型 **取值范围**: - UDP - UDP6 - TCP - TCP6
	Type *string `json:"type,omitempty"`

	// **参数解释**: 端口数量 **取值范围**: 最小值0，最大值65535
	Num *int32 `json:"num,omitempty"`

	// **参数解释**: 危险类型 **取值范围**: - danger: 危险端口 - normal: 正常端口 - unknow: 无已知危险的端口
	Status *string `json:"status,omitempty"`
}

func (o PortStatisticResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PortStatisticResponseInfo struct{}"
	}

	return strings.Join([]string{"PortStatisticResponseInfo", string(data)}, " ")
}
