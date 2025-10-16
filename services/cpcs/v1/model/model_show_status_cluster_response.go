package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowStatusClusterResponse Response Object
type ShowStatusClusterResponse struct {

	// 资源名称
	MetricName *string `json:"metric_name,omitempty"`

	// 集群列表
	ClusterList    *[]ShowStatusClusterItem `json:"cluster_list,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ShowStatusClusterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowStatusClusterResponse struct{}"
	}

	return strings.Join([]string{"ShowStatusClusterResponse", string(data)}, " ")
}
