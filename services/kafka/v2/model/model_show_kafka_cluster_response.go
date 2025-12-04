package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowKafkaClusterResponse Response Object
type ShowKafkaClusterResponse struct {
	Cluster        *ShowKakfaClusterResponseCluster `json:"cluster,omitempty"`
	HttpStatusCode int                              `json:"-"`
}

func (o ShowKafkaClusterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowKafkaClusterResponse struct{}"
	}

	return strings.Join([]string{"ShowKafkaClusterResponse", string(data)}, " ")
}
