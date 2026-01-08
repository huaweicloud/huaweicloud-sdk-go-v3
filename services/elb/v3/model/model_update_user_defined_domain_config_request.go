package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateUserDefinedDomainConfigRequest Request Object
type UpdateUserDefinedDomainConfigRequest struct {

	// 负载均衡器ID。
	LoadbalancerId string `json:"loadbalancer_id"`

	Body *UpdateUserDefinedDomainConfigRequestBody `json:"body,omitempty"`
}

func (o UpdateUserDefinedDomainConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateUserDefinedDomainConfigRequest struct{}"
	}

	return strings.Join([]string{"UpdateUserDefinedDomainConfigRequest", string(data)}, " ")
}
