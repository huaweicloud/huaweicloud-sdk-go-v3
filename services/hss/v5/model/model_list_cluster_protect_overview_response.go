package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListClusterProtectOverviewResponse Response Object
type ListClusterProtectOverviewResponse struct {

	// 集群数量
	ClusterNum *int32 `json:"cluster_num,omitempty"`

	// 防护中数量
	UnderProtectNum *int32 `json:"under_protect_num,omitempty"`

	// 策略数量
	PolicyNum *int32 `json:"policy_num,omitempty"`

	// 事件数量
	EventNum *int32 `json:"event_num,omitempty"`

	// 已阻断事件数量
	BlockNum       *int32 `json:"block_num,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListClusterProtectOverviewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListClusterProtectOverviewResponse struct{}"
	}

	return strings.Join([]string{"ListClusterProtectOverviewResponse", string(data)}, " ")
}
