package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListClusterConfigurationsResponse struct {

	// 集群所关联的参数组信息。
	Configurations *[]ClusterConfiguration `json:"configurations,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ListClusterConfigurationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListClusterConfigurationsResponse struct{}"
	}

	return strings.Join([]string{"ListClusterConfigurationsResponse", string(data)}, " ")
}
