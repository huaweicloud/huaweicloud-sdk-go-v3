package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAgentInstancesRequest Request Object
type ListAgentInstancesRequest struct {

	// Agent 类型，支持多选（OR 逻辑）：OpenClaw / OfficeClaw / HermesAgent
	AiAgentType *[]string `json:"ai_agent_type,omitempty"`

	// 区域 ID，支持多选（OR 逻辑）
	RegionId *[]string `json:"region_id,omitempty"`

	// Agent 运行状态，支持多选（OR 逻辑）： - UNREACHABLE：连续3次心跳丢失（90秒无上报），触发告警 - ERROR：Agent进程健康检查连续失败3次或进程异常退出，尝试自动重启 - OFFLINE：桌面关机或重建中，停止心跳检测 - RUNNING：心跳正常且Agent进程健康检查通过
	AgentStatus *[]string `json:"agent_status,omitempty"`

	// 桌面运行状态，支持多选（OR 逻辑）： - ACTIVE：运行中 - SHUTOFF：已关机 - HIBERNATED：已休眠 - ERROR：故障
	DesktopStatus *[]string `json:"desktop_status,omitempty"`

	// 桌面连接状态，支持多选（OR 逻辑）： - UNREGISTER：桌面未注册（关机后也会出现） - REGISTERED：桌面已注册，等待用户连接 - CONNECTED：用户已连接，正在使用桌面 - DISCONNECTED：桌面与客户端断开会话
	DesktopConnection *[]string `json:"desktop_connection,omitempty"`

	// 已授权的模型分组 ID（单选）
	ModelGroupId *string `json:"model_group_id,omitempty"`

	// 通道配置状态，支持多选（OR 逻辑）： - UNCONFIGURED：未配置 - APPLYING：配置中 - CONFIGURED：已配置 - FAILED：配置失败
	ChannelConfigStatus *[]string `json:"channel_config_status,omitempty"`

	// 实例名称（模糊搜索）
	InstanceName *string `json:"instance_name,omitempty"`

	// 实例 ID（精确搜索）
	InstanceId *string `json:"instance_id,omitempty"`

	// 云桌面 ID（精确搜索）
	DesktopId *string `json:"desktop_id,omitempty"`

	// 创建时间范围-开始
	CreateTimeStart *sdktime.SdkTime `json:"create_time_start,omitempty"`

	// 创建时间范围-结束
	CreateTimeEnd *sdktime.SdkTime `json:"create_time_end,omitempty"`

	// 标签过滤，格式：key1=val1,key2=val2，多个键值对用逗号分隔
	Tags *string `json:"tags,omitempty"`

	// 风险类型过滤，支持多选（OR 逻辑）： - MODEL_CONFIG_INCONSISTENT：模型配置不一致 - IM_CHANNEL_CONFIG_INCONSISTENT：IM通道配置不一致
	RiskType *[]string `json:"risk_type,omitempty"`

	// 模型配置状态，支持多选（OR 逻辑）： - UNCONFIGURED：未配置 - APPLYING：配置中 - CONFIGURED：已配置 - FAILED：配置失败
	ModelConfigStatus *[]string `json:"model_config_status,omitempty"`

	// Agent 版本号（精确搜索）
	AgentVersion *string `json:"agent_version,omitempty"`

	// 排序字段：create_time（默认）/ instance_name / agent_status / heartbeat_time
	SortField *string `json:"sort_field,omitempty"`

	// 排序方向：DESC（默认）/ ASC
	SortOrder *string `json:"sort_order,omitempty"`

	// 偏移量，从0开始
	Offset *int32 `json:"offset,omitempty"`

	// 每页条数
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListAgentInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAgentInstancesRequest struct{}"
	}

	return strings.Join([]string{"ListAgentInstancesRequest", string(data)}, " ")
}
