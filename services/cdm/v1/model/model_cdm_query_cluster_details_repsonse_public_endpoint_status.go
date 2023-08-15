package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CdmQueryClusterDetailsRepsonsePublicEndpointStatus EIP状态信息
type CdmQueryClusterDetailsRepsonsePublicEndpointStatus struct {

	// 状态
	Status *string `json:"status,omitempty"`

	// 错误信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
}

func (o CdmQueryClusterDetailsRepsonsePublicEndpointStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CdmQueryClusterDetailsRepsonsePublicEndpointStatus struct{}"
	}

	return strings.Join([]string{"CdmQueryClusterDetailsRepsonsePublicEndpointStatus", string(data)}, " ")
}
