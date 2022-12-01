package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateClusterDnsRequest struct {

	// 集群的ID
	ClusterId string `json:"cluster_id"`

	Body *ModifyClusterDns `json:"body,omitempty"`
}

func (o UpdateClusterDnsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateClusterDnsRequest struct{}"
	}

	return strings.Join([]string{"UpdateClusterDnsRequest", string(data)}, " ")
}
