package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopRedistributionRequest Request Object
type StopRedistributionRequest struct {

	// 指定暂停重分布集群的ID
	ClusterId string `json:"cluster_id"`
}

func (o StopRedistributionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopRedistributionRequest struct{}"
	}

	return strings.Join([]string{"StopRedistributionRequest", string(data)}, " ")
}
