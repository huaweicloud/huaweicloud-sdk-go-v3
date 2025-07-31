package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDaemonsetInfo 升级Ds请求信息
type UpdateDaemonsetInfo struct {

	// agent版本
	AgentVersion string `json:"agent_version"`

	// 集群Id
	ClusterId string `json:"cluster_id"`
}

func (o UpdateDaemonsetInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDaemonsetInfo struct{}"
	}

	return strings.Join([]string{"UpdateDaemonsetInfo", string(data)}, " ")
}
