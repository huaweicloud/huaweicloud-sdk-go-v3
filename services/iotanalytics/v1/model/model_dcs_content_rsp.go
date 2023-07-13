package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DcsContentRsp redis数据源配置内容
type DcsContentRsp struct {

	// VPC-EP服务端id
	EndpointServiceId *string `json:"endpointServiceId,omitempty"`

	// VPC-EP服务端名称
	EndpointServiceName *string `json:"endpointServiceName,omitempty"`

	// VPC-EP客户端IP
	EndpointIp *string `json:"endpointIp,omitempty"`

	// VPC-EP客户端Port
	Port *int32 `json:"port,omitempty"`

	// redis实例类型
	DcsType *string `json:"dcsType,omitempty"`

	// redis访问密码
	Password *string `json:"password,omitempty"`
}

func (o DcsContentRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DcsContentRsp struct{}"
	}

	return strings.Join([]string{"DcsContentRsp", string(data)}, " ")
}
