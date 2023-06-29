package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListClusterConfigurationsParameterRequest Request Object
type ListClusterConfigurationsParameterRequest struct {

	// 集群的ID。
	ClusterId string `json:"cluster_id"`

	// 参数组ID。
	ConfigurationId string `json:"configuration_id"`
}

func (o ListClusterConfigurationsParameterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListClusterConfigurationsParameterRequest struct{}"
	}

	return strings.Join([]string{"ListClusterConfigurationsParameterRequest", string(data)}, " ")
}
