package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateVpcEndpointServiceResponse Response Object
type UpdateVpcEndpointServiceResponse struct {

	// 创建成功
	Success        *bool `json:"success,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o UpdateVpcEndpointServiceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVpcEndpointServiceResponse struct{}"
	}

	return strings.Join([]string{"UpdateVpcEndpointServiceResponse", string(data)}, " ")
}
