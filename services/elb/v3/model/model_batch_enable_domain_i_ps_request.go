package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchEnableDomainIPsRequest Request Object
type BatchEnableDomainIPsRequest struct {

	// 负载均衡器ID。
	LoadbalancerId string `json:"loadbalancer_id"`

	Body *BatchEnableDomainIPsRequestBody `json:"body,omitempty"`
}

func (o BatchEnableDomainIPsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchEnableDomainIPsRequest struct{}"
	}

	return strings.Join([]string{"BatchEnableDomainIPsRequest", string(data)}, " ")
}
