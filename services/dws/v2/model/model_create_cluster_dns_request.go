package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateClusterDnsRequest Request Object
type CreateClusterDnsRequest struct {

	// 集群的ID
	ClusterId string `json:"cluster_id"`

	Body *CreateClusterDns `json:"body,omitempty"`
}

func (o CreateClusterDnsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateClusterDnsRequest struct{}"
	}

	return strings.Join([]string{"CreateClusterDnsRequest", string(data)}, " ")
}
