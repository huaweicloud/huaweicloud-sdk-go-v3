package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateEndpointRequestBody struct {

	// 待创建的终端节点名称。 取值范围：1-64个字符，支持数字、字母、中文、_（下划线）、-（中划线）、.（点）。
	Name string `json:"name"`

	// 终端节点方向。 取值： inbound，表示入站终端节点。 outbound，表示出站终端节点。
	Direction string `json:"direction"`

	// 当前子网所在的region。
	Region string `json:"region"`

	// 终端节点的IP地址和子网信息。最少需要添加2个IP地址，最多可以添加6个IP地址。
	Ipaddresses []IpaddressInfo `json:"ipaddresses"`
}

func (o CreateEndpointRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEndpointRequestBody struct{}"
	}

	return strings.Join([]string{"CreateEndpointRequestBody", string(data)}, " ")
}
