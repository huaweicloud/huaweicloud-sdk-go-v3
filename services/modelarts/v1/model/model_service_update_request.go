package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ServiceUpdateRequest struct {

	// **参数解释：** 服务ID **约束限制：** 不涉及。 **取值范围：** 不涉及
	Id *string `json:"id,omitempty"`

	// **参数解释：** 服务名称。 **约束限制：** 不涉及。 **取值范围：** 支持1-128个字符，可以包含字母、汉字、数字、连字符和下划线。 **默认取值：** 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释：** 服务部署超时时间，integer类型，取值在1~300（860版本该参数做保留兼容）。 **约束限制：** 不涉及。 **取值范围：** [0, 300]。 **默认取值：** 不涉及。
	DeployTimeoutMinutes *int32 `json:"deploy_timeout_minutes,omitempty"`

	// **参数解释：** 必填，填了之后，数据库中如果存在相同版本号，将会报错（仅修改描述的场景除外）。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Version *string `json:"version,omitempty"`

	// **参数解释：** 非必填，仅更新描述的场景直接修改对应version的数据库字段，不新增版本号。 **约束限制：** 不涉及。 **取值范围：** 不涉及 **默认取值：** 不涉及
	Description *string `json:"description,omitempty"`

	// **参数解释：** 仅修改服务时不需传，兼容部署分离之前版本。 **约束限制：** group_configs的最大元素数量为1。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	GroupConfigs *[]GroupConfig `json:"group_configs,omitempty"`

	RuntimeConfig *RuntimeConfig `json:"runtime_config,omitempty"`

	UpgradeConfig *UpgradeConfig `json:"upgrade_config,omitempty"`

	// **参数解释：** 日志策略。 **约束限制：** 不涉及。 **取值范围：** - POOL：使用资源池日志插件配置的日志流。 - AUTO_CREATE：自动创建日志流。 - DEFAULT: 由系统决定日志策略 **默认取值：** AUTO_CREATE：自动创建日志流。
	LtsStrategy *string `json:"lts_strategy,omitempty"`

	// **参数解释：** 日志配置，当开启LTS日志的时候，STDOUT类型为必填。 **约束限制：** 当开启LTS日志的时候，STDOUT类型为必填。 数量上限为2个。
	LogConfigs *[]LtsConfig `json:"log_configs,omitempty"`

	// **参数解释：** 服务标签,上限20个 **约束限制：** 不涉及。 **取值范围：** 不涉及 **默认取值：** 不涉及
	Tags *string `json:"tags,omitempty"`

	// **参数解释：** 工作空间id，默认是“0” **约束限制：** 不涉及。 **取值范围：** 不涉及 **默认取值：** 不涉及
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// **参数解释：** 定时停止配置。 **约束限制：** 仅当body中另一个参数description为空时，此参数才生效。
	Schedule *[]ScheduleConfig `json:"schedule,omitempty"`

	// **参数解释：** 该参数值由英文逗号隔开的协议、端口号、地址组成，其中地址长度不超过255 ，且需要与镜像给定的协议、地址、端口一致，否则指标无法上报。
	CustomMetricsPath *string `json:"custom_metrics_path,omitempty"`

	// **参数解释：** 模型类型。 **取值范围：** - TEXT_GENERATION：文本生成 - IMAGE_UNDERSTANDING：图像理解 - VIDEO_GENERATION：视频生成 - IMAGE_GENERATION：图像生成 - RERANK：重排序 - VECTOR_MODEL：向量模型 - EMBEDDING：Embedding嵌入
	TaskType *string `json:"task_type,omitempty"`
}

func (o ServiceUpdateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServiceUpdateRequest struct{}"
	}

	return strings.Join([]string{"ServiceUpdateRequest", string(data)}, " ")
}
