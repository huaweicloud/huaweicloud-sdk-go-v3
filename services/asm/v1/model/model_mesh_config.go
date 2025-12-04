package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MeshConfig 网格配置
type MeshConfig struct {
	ProxyConfig *ProxyConfig `json:"proxyConfig,omitempty"`

	TelemetryConfig *TelemetryConfig `json:"telemetryConfig,omitempty"`
}

func (o MeshConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MeshConfig struct{}"
	}

	return strings.Join([]string{"MeshConfig", string(data)}, " ")
}
