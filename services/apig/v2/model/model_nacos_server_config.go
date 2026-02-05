package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NacosServerConfig nacos服务端配置信息。
type NacosServerConfig struct {

	// nacos服务端IP地址。不包含中文字符。
	IpAddress string `json:"ip_address"`

	// nacos服务端端口号。取值范围1 ~ 65535。
	Port int32 `json:"port"`

	// nacos服务端gRPC端口号，默认为port+1000。取值范围1 ~ 65535。
	GrpcPort *int32 `json:"grpc_port,omitempty"`
}

func (o NacosServerConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NacosServerConfig struct{}"
	}

	return strings.Join([]string{"NacosServerConfig", string(data)}, " ")
}
