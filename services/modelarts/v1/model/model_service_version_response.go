package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ServiceVersionResponse **参数解释：** 当前服务版本信息。
type ServiceVersionResponse struct {

	// **参数解释：** 版本id。 **取值范围：** 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释：** 版本号。 **取值范围：** 不涉及。
	Version *string `json:"version,omitempty"`

	// **参数解释：** 描述。 **取值范围：** 不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释：** 推理请求的访问地址。 **取值范围：** 不涉及。
	PredictUrl *string `json:"predict_url,omitempty"`

	RuntimeConfig *RuntimeConfigResponse `json:"runtime_config,omitempty"`

	UpgradeConfig *UpgradeConfigResponse `json:"upgrade_config,omitempty"`

	// **参数解释：** 服务部署信息。
	InstanceGroups *[]GroupConfigResponse `json:"instance_groups,omitempty"`

	// **参数解释：** 日志策略。 **取值范围：** - POOL：使用资源池日志插件配置的日志流。 - AUTO_CREATE：自动创建日志流。 - DEFAULT: 由系统决定日志策略
	LtsStrategy *string `json:"lts_strategy,omitempty"`

	// **参数解释：** 服务容器标准输出对接lts开关状态。 **取值范围：** - ON：开启。 - OFF：关闭。
	LtsStatus *string `json:"lts_status,omitempty"`

	// **参数解释：** 服务对接lts k8s事件开关状态。 **取值范围：** - ON：开启。 - OFF：关闭。
	LtsEventStatus *string `json:"lts_event_status,omitempty"`

	// **参数解释：** 服务容器日志文件对接lts开关状态。 **取值范围：** - ON：开启。 - OFF：关闭。
	LtsFileStatus *string `json:"lts_file_status,omitempty"`

	// **参数解释：** 服务日志配置信息。
	LogConfigs *[]LogConfigResponse `json:"log_configs,omitempty"`

	// **参数解释：** 部署超时时间。 **取值范围：** 不涉及。
	DeployTimeoutMinutes *int32 `json:"deploy_timeout_minutes,omitempty"`
}

func (o ServiceVersionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServiceVersionResponse struct{}"
	}

	return strings.Join([]string{"ServiceVersionResponse", string(data)}, " ")
}
