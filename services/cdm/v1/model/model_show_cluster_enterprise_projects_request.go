package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowClusterEnterpriseProjectsRequest Request Object
type ShowClusterEnterpriseProjectsRequest struct {

	// 集群ID
	ClusterId string `json:"cluster_id"`
}

func (o ShowClusterEnterpriseProjectsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowClusterEnterpriseProjectsRequest struct{}"
	}

	return strings.Join([]string{"ShowClusterEnterpriseProjectsRequest", string(data)}, " ")
}
