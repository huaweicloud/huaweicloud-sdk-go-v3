package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowKakfaClusterResponseCluster **参数解释**： 集群信息。
type ShowKakfaClusterResponseCluster struct {

	// **参数解释**： broker详情。
	Brokers *[]ShowKakfaClusterResponseClusterBrokers `json:"brokers,omitempty"`
}

func (o ShowKakfaClusterResponseCluster) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowKakfaClusterResponseCluster struct{}"
	}

	return strings.Join([]string{"ShowKakfaClusterResponseCluster", string(data)}, " ")
}
