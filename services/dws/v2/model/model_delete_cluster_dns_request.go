package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteClusterDnsRequest Request Object
type DeleteClusterDnsRequest struct {

	// 集群的ID
	ClusterId string `json:"cluster_id"`

	// 域名类型，支持删除公网域名。
	Type string `json:"type"`
}

func (o DeleteClusterDnsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteClusterDnsRequest struct{}"
	}

	return strings.Join([]string{"DeleteClusterDnsRequest", string(data)}, " ")
}
