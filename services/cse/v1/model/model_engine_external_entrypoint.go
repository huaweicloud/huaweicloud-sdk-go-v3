package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EngineExternalEntrypoint struct {

	// 微服务引擎专享版暴露的IP地址。
	ExternalAddress *string `json:"external_address,omitempty"`

	// 微服务引擎专享版的公网地址。
	PublicAddress *string `json:"public_address,omitempty"`

	// 微服务引擎专享版组件的访问地址。
	ServiceEndpoint map[string]EntrypointItem `json:"service_endpoint,omitempty"`

	// 微服务引擎专享版组件的公网地址。
	PublicServiceEndpoint map[string]EntrypointItem `json:"public_service_endpoint,omitempty"`
}

func (o EngineExternalEntrypoint) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EngineExternalEntrypoint struct{}"
	}

	return strings.Join([]string{"EngineExternalEntrypoint", string(data)}, " ")
}
