package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EIP状态信息
type CdmQueryClusterDetailsRepsonsePublicEndpointStatus struct {

	// 状态
	Status *string `json:"status,omitempty" xml:"status"`

	// 错误信息
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage"`
}

func (o CdmQueryClusterDetailsRepsonsePublicEndpointStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CdmQueryClusterDetailsRepsonsePublicEndpointStatus struct{}"
	}

	return strings.Join([]string{"CdmQueryClusterDetailsRepsonsePublicEndpointStatus", string(data)}, " ")
}
