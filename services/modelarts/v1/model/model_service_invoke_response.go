package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ServiceInvokeResponse **参数解释：**  服务调用时的相关配置。
type ServiceInvokeResponse struct {

	// **参数解释：** 服务端口号。 **取值范围：** [1, 65535]。
	Port int32 `json:"port"`

	// **参数解释：** 服务请求协议。当选择WSS与WS时，服务接口会升级为WebSocket。开启WebSocket时，不支持同时设置“服务流量限制”。 **取值范围：** - HTTP：HTTP协议。 - HTTPS：HTTPS协议。 - WSS：WebSocket Secure协议。 - WS：WebSocket协议。
	Protocol string `json:"protocol"`

	// **参数解释：** 认证类型。 **取值范围：** - TOKEN：IAM Token认证。 - API_KEY：API Key认证。 - NONE：无认证。
	AuthType string `json:"auth_type"`

	// **参数解释：** 外网访问。 **取值范围：** - TRUE：要外网访问。 - FALSE：不要外网访问。
	InternetAccessEnable *bool `json:"internet_access_enable,omitempty"`

	// **参数解释：** 内网审批。 **取值范围：** - TRUE：要内网审批。 - FALSE：不要内网审批。
	IntranetApprovalEnable *bool `json:"intranet_approval_enable,omitempty"`

	// **参数解释：** 动态路由开关。 **取值范围：** - TRUE：开启动态路由。 - FALSE：不开启动态路由。
	DynamicRoutingEnable *bool `json:"dynamic_routing_enable,omitempty"`

	// **参数解释：** 智能路由策略。 **取值范围：** - ROUND_ROBIN：轮询。 - ORIGIN_IP_HASH：源IP哈希。 - MIN_CONN：最小连接数。 - MIN_FIRST_TOKEN_TIME：最小首token时延。 - COMPOSITE：综合负载。 - SLO_BASED：SLO优先级。
	Strategy *string `json:"strategy,omitempty"`

	// **参数解释：** 指标接口服务请求协议。 **取值范围：** - HTTP：HTTP协议。 - HTTPS：HTTPS协议。
	MetricApiScheme *string `json:"metric_api_scheme,omitempty"`

	// **参数解释：** 指标接口端口号。 **取值范围：** [1, 65535]。
	MetricApiPort *string `json:"metric_api_port,omitempty"`

	// **参数解释：** 指标接口path。 **取值范围：** 不涉及。
	MetricApiPath *string `json:"metric_api_path,omitempty"`

	// **参数解释：** 是否开启EMS加速。 **取值范围：** - TRUE：开启EMS加速。 - FALSE：不开启EMS加速。
	EmsEnable *bool `json:"ems_enable,omitempty"`

	// **参数解释：** proxy支持请求重调度开关。 **取值范围：** - true：开启proxy支持请求重调度。 - false：不开启proxy支持请求重调度。
	RequestRetryEnable *bool `json:"request_retry_enable,omitempty"`

	// **参数解释：** proxy支持请求重调度的重试次数 **取值范围：** [1, 10]。
	RequestRetryCntMax *int32 `json:"request_retry_cnt_max,omitempty"`

	// **参数解释：** proxy支持请求重调度的重试间隔，单位ms **取值范围：** [1, 10000]。
	RequestRetryIntervalMs *int32 `json:"request_retry_interval_ms,omitempty"`

	FuseConfigs *FuseConfig `json:"fuse_configs,omitempty"`

	ElbConnection *ElbConnectionResponse `json:"elb_connection,omitempty"`
}

func (o ServiceInvokeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServiceInvokeResponse struct{}"
	}

	return strings.Join([]string{"ServiceInvokeResponse", string(data)}, " ")
}
