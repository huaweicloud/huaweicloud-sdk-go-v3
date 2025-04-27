package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEndpointVpcsResponse Response Object
type ListEndpointVpcsResponse struct {

	// 终端节点VPC列表。
	Vpcs *[]VpcsData `json:"vpcs,omitempty"`

	Metadata       *Metadata `json:"metadata,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListEndpointVpcsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEndpointVpcsResponse struct{}"
	}

	return strings.Join([]string{"ListEndpointVpcsResponse", string(data)}, " ")
}
