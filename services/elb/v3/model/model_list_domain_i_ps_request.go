package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDomainIPsRequest Request Object
type ListDomainIPsRequest struct {

	// 负载均衡器ID。
	LoadbalancerId string `json:"loadbalancer_id"`
}

func (o ListDomainIPsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDomainIPsRequest struct{}"
	}

	return strings.Join([]string{"ListDomainIPsRequest", string(data)}, " ")
}
