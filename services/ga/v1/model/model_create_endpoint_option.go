package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateEndpointOption 创建终端节点详细信息。
type CreateEndpointOption struct {

	// 对应后端资源的ID，比如EIP的ID。
	ResourceId string `json:"resource_id"`

	ResourceType *EndpointType `json:"resource_type"`

	// 终端节点权重。
	Weight *int32 `json:"weight,omitempty"`

	// IP地址。
	IpAddress string `json:"ip_address"`
}

func (o CreateEndpointOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEndpointOption struct{}"
	}

	return strings.Join([]string{"CreateEndpointOption", string(data)}, " ")
}
