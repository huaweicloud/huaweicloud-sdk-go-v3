package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConnectEndpoint 连接网关信息
type ConnectEndpoint struct {

	// 取值：当前仅支持 VPCEP
	Type *string `json:"type,omitempty"`

	// 资源名称，当type是VPCEP时为VPCEP Service的名称
	Name *string `json:"name,omitempty"`

	// 资源id，当type是VPCEP时为VPCEP Service的id
	Id *string `json:"id,omitempty"`

	// 资源所在的region
	Region *string `json:"region,omitempty"`
}

func (o ConnectEndpoint) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConnectEndpoint struct{}"
	}

	return strings.Join([]string{"ConnectEndpoint", string(data)}, " ")
}
