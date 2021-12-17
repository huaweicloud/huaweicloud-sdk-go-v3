package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DCS数据源配置内容
type DcsContentReq struct {
	// VPC-EP服务端id

	EndpointServiceId string `json:"endpoint_service_id"`
	// VPC-EP服务端名称

	EndpointServiceName string `json:"endpoint_service_name"`
	// VPC-EP客户端Port

	Port int32 `json:"port"`
	// redis实例类型

	DcsType string `json:"dcs_type"`
	// redis访问密码

	Password string `json:"password"`
}

func (o DcsContentReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DcsContentReq struct{}"
	}

	return strings.Join([]string{"DcsContentReq", string(data)}, " ")
}
