package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VpcsData struct {

	// VPC的ID，UUID形式的一个资源标识。
	Id *string `json:"id,omitempty"`

	// VPC下入站终端节点的数量。
	InboundEndpointCount *int32 `json:"inbound_endpoint_count,omitempty"`

	// VPC下出站终端节点的数量。
	OutboundEndpointCount *int32 `json:"outbound_endpoint_count,omitempty"`
}

func (o VpcsData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VpcsData struct{}"
	}

	return strings.Join([]string{"VpcsData", string(data)}, " ")
}
