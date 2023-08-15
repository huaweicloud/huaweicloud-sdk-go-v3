package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAvailableDisasterClustersResponse Response Object
type ListAvailableDisasterClustersResponse struct {

	// 容灾可用集群列表
	DisasterRecoveryClusters *[]DisasterRecoveryClusterVo `json:"disaster_recovery_clusters,omitempty"`
	HttpStatusCode           int                          `json:"-"`
}

func (o ListAvailableDisasterClustersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAvailableDisasterClustersResponse struct{}"
	}

	return strings.Join([]string{"ListAvailableDisasterClustersResponse", string(data)}, " ")
}
