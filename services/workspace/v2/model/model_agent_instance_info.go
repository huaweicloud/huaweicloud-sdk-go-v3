package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AgentInstanceInfo Agent 实例信息
type AgentInstanceInfo struct {

	// 主键 ID
	Id *string `json:"id,omitempty"`

	// Agent 实例 ID
	InstanceId *string `json:"instance_id,omitempty"`

	// 实例名称
	InstanceName *string `json:"instance_name,omitempty"`

	// 关联云桌面 ID
	DesktopId *string `json:"desktop_id,omitempty"`

	// 区域 ID
	RegionId *string `json:"region_id,omitempty"`

	// Agent 类型
	AiAgentType *string `json:"ai_agent_type,omitempty"`

	// Agent 运行状态： - UNREACHABLE：连续3次心跳丢失（90秒无上报），触发告警 - ERROR：Agent进程健康检查连续失败3次或进程异常退出，尝试自动重启 - OFFLINE：桌面关机或重建中，停止心跳检测 - RUNNING：心跳正常且Agent进程健康检查通过
	AgentStatus *string `json:"agent_status,omitempty"`

	// 桌面运行状态： - ACTIVE：运行中 - SHUTOFF：已关机 - HIBERNATED：已休眠 - ERROR：故障
	DesktopStatus *string `json:"desktop_status,omitempty"`

	// 桌面连接状态： - UNREGISTER：桌面未注册（关机后也会出现） - REGISTERED：桌面已注册，等待用户连接 - CONNECTED：用户已连接，正在使用桌面 - DISCONNECTED：桌面与客户端断开会话
	DesktopConnection *string `json:"desktop_connection,omitempty"`

	// 模型配置状态
	ModelConfigStatus *string `json:"model_config_status,omitempty"`

	// 通道配置状态： - UNCONFIGURED：未配置 - APPLYING：配置中 - CONFIGURED：已配置 - FAILED：配置失败
	ChannelConfigStatus *string `json:"channel_config_status,omitempty"`

	// IM 通道配置 ID 列表
	ImChannelConfigs *[]string `json:"im_channel_configs,omitempty"`

	// 创建时间
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 更新时间
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`

	// 产品 ID
	ProductId *string `json:"product_id,omitempty"`

	// 产品名称
	ProductName *string `json:"product_name,omitempty"`

	// 镜像 ID
	ImageId *string `json:"image_id,omitempty"`

	// 镜像名称
	ImageName *string `json:"image_name,omitempty"`

	// 桌面池 ID
	DesktopPoolId *string `json:"desktop_pool_id,omitempty"`

	// 用户名
	UserName *string `json:"user_name,omitempty"`

	// 风险列表
	Risks *[]AgentRisk `json:"risks,omitempty"`

	// Agent 版本号
	AgentVersion *string `json:"agent_version,omitempty"`

	// 企业项目 ID
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o AgentInstanceInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AgentInstanceInfo struct{}"
	}

	return strings.Join([]string{"AgentInstanceInfo", string(data)}, " ")
}
