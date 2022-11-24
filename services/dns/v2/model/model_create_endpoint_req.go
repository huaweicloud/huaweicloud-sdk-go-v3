package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateEndpointReq struct {

	// 待终端节点名称。 取值范围：1-64个字符，支持数字、字母、中文、_（下划线）、-（中划线）、.（点）。
	Name string `json:"name"`

	// 终端节点方向。 取值： inbound，表示入站规则。 outbound，表示出站规则。
	Direction string `json:"direction"`

	// 当前子网所在的region。
	Region string `json:"region"`

	// 终端节点的ip地址和子网信息。
	Ipaddresses []Ipaddresses `json:"ipaddresses"`
}

func (o CreateEndpointReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEndpointReq struct{}"
	}

	return strings.Join([]string{"CreateEndpointReq", string(data)}, " ")
}
