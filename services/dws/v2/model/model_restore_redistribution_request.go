package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestoreRedistributionRequest Request Object
type RestoreRedistributionRequest struct {

	// 指定恢复重分布集群的ID
	ClusterId string `json:"cluster_id"`
}

func (o RestoreRedistributionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreRedistributionRequest struct{}"
	}

	return strings.Join([]string{"RestoreRedistributionRequest", string(data)}, " ")
}
