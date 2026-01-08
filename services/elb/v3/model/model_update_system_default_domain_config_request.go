package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSystemDefaultDomainConfigRequest Request Object
type UpdateSystemDefaultDomainConfigRequest struct {

	// 负载均衡器ID。
	LoadbalancerId string `json:"loadbalancer_id"`

	Body *UpdateSystemDefaultDomainConfigRequestBody `json:"body,omitempty"`
}

func (o UpdateSystemDefaultDomainConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSystemDefaultDomainConfigRequest struct{}"
	}

	return strings.Join([]string{"UpdateSystemDefaultDomainConfigRequest", string(data)}, " ")
}
