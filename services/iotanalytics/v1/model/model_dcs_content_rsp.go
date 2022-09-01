package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// redis数据源配置内容
type DcsContentRsp struct {

	// VPC-EP服务端id
	EndpointServiceId *string `json:"endpointServiceId,omitempty" xml:"endpointServiceId"`

	// VPC-EP服务端名称
	EndpointServiceName *string `json:"endpointServiceName,omitempty" xml:"endpointServiceName"`

	// VPC-EP客户端IP
	EndpointIp *string `json:"endpointIp,omitempty" xml:"endpointIp"`

	// VPC-EP客户端Port
	Port *int32 `json:"port,omitempty" xml:"port"`

	// redis实例类型
	DcsType *string `json:"dcsType,omitempty" xml:"dcsType"`

	// redis访问密码
	Password *string `json:"password,omitempty" xml:"password"`
}

func (o DcsContentRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DcsContentRsp struct{}"
	}

	return strings.Join([]string{"DcsContentRsp", string(data)}, " ")
}
