package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateLoadBalancersRequest Request Object
type BatchCreateLoadBalancersRequest struct {
	Body *BatchCreateLoadBalancersRequestBody `json:"body,omitempty"`
}

func (o BatchCreateLoadBalancersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateLoadBalancersRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateLoadBalancersRequest", string(data)}, " ")
}
