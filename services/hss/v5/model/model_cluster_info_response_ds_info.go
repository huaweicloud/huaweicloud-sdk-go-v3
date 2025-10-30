package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ClusterInfoResponseDsInfo daemonset状态
type ClusterInfoResponseDsInfo struct {

	// **参数解释** 目标数量 **取值范围** 取值0-65535
	DesiredNum *int32 `json:"desired_num,omitempty"`

	// **参数解释** 当前数量 **取值范围** 取值0-65535
	CurrentNum *int32 `json:"current_num,omitempty"`

	// **参数解释** 就绪数量 **取值范围** 取值0-65535
	ReadyNum *int32 `json:"ready_num,omitempty"`
}

func (o ClusterInfoResponseDsInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterInfoResponseDsInfo struct{}"
	}

	return strings.Join([]string{"ClusterInfoResponseDsInfo", string(data)}, " ")
}
