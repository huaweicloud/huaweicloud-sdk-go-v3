package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AdvancedConfig 部署高级配置
type AdvancedConfig struct {

	// **参数解释：** 部署超时时间
	DeployTimeoutMinutes *string `json:"deploy_timeout_minutes,omitempty"`

	UpgradeConfig *UpgradeConfig `json:"upgrade_config"`

	ServiceSecret *ServiceSecret `json:"service_secret,omitempty"`

	// **参数解释：** 智能路由开关
	DynamicRoutingEnable *bool `json:"dynamic_routing_enable,omitempty"`

	// **参数解释：** 智能路由策略
	Strategy *string `json:"strategy,omitempty"`

	// **参数解释：** EMS加速开关
	EmsEnable *bool `json:"ems_enable,omitempty"`

	// **参数解释：** 智能路由指标采集scheme
	MetricApiScheme *string `json:"metric_api_scheme,omitempty"`

	// **参数解释：** 智能路由指标采集端口
	MetricApiPort *string `json:"metric_api_port,omitempty"`

	// **参数解释：** 智能路由指标采集地址
	MetricApiPath *string `json:"metric_api_path,omitempty"`

	// **参数解释：** 自定义监控采集指标地址
	CustomMetricsPath *string `json:"custom_metrics_path,omitempty"`

	// **参数解释：** 容器端口
	Port int32 `json:"port"`

	// **参数解释：** 容器请求协议。当选择WSS与WS时，服务接口会升级为WebSocket。开启WebSocket时，不支持同时设置“服务流量限制”。 **取值范围：** - HTTP：HTTP协议。 - HTTPS：HTTPS协议。 - WSS：WebSocket Secure协议。 - WS：WebSocket协议。 - TCP：传输控制协议。 - NA：不使用任何协议。
	Protocol string `json:"protocol"`
}

func (o AdvancedConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AdvancedConfig struct{}"
	}

	return strings.Join([]string{"AdvancedConfig", string(data)}, " ")
}
