package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowQuotasRequest Request Object
type ShowQuotasRequest struct {

	// 支持过滤的配额类型： - [physicalConnect:  物理连接direct_connect实例的配额和使用量] - [virtualInterface: 虚拟接口virtual-interface的配额和使用量] - [connectGateway: 连接网关（用于关联GEIP）的配额和使用量](tag:hws) - [geip: 每租户可以关联GEIP的配额和使用量](tag:hws) - [globalDcGateway 专线全球接入网关的配额和使用量] - [peerLinkPerGdgw: 接入网关的关联连接的配额和使用量]
	Type *[]string `json:"type,omitempty"`
}

func (o ShowQuotasRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowQuotasRequest struct{}"
	}

	return strings.Join([]string{"ShowQuotasRequest", string(data)}, " ")
}
