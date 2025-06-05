package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ClusterInfoResponseDsInfo ds状态
type ClusterInfoResponseDsInfo struct {

	// 目标数量
	DesiredNum *int32 `json:"desired_num,omitempty"`

	// 当前数量
	CurrentNum *int32 `json:"current_num,omitempty"`

	// 就绪数量
	ReadyNum *int32 `json:"ready_num,omitempty"`
}

func (o ClusterInfoResponseDsInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterInfoResponseDsInfo struct{}"
	}

	return strings.Join([]string{"ClusterInfoResponseDsInfo", string(data)}, " ")
}
