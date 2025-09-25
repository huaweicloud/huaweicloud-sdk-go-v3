package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAgentStatisticsResponse Response Object
type ShowAgentStatisticsResponse struct {

	// 待升级数量
	WaitUpgradeNum *int32 `json:"wait_upgrade_num,omitempty"`

	// 在线数量
	OnlineNum *int32 `json:"online_num,omitempty"`

	// 不在线数量
	NotOnlineNum *int32 `json:"not_online_num,omitempty"`

	// 离线数量
	OfflineNum *int32 `json:"offline_num,omitempty"`

	// 集群内节点数
	InclusterNum *int32 `json:"incluster_num,omitempty"`

	// 非集群内节点数
	NotInclusterNum *int32 `json:"not_incluster_num,omitempty"`
	HttpStatusCode  int    `json:"-"`
}

func (o ShowAgentStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAgentStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ShowAgentStatisticsResponse", string(data)}, " ")
}
