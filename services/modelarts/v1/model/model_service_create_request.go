package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ServiceCreateRequest **参数解释：** 创建服务请求体。 **取值范围：** 不涉及。
type ServiceCreateRequest struct {

	// **参数解释：** 服务名，用户在[创建服务](CreateInferService.xml)时自定义的名称。 **约束限制：** 服务在删除之前名字不能重复。 **取值范围：** 支持1-64位字符，可包含字母、中文、数字、中划线、下划线。 **默认取值：** 不涉及。
	Name string `json:"name"`

	// **参数解释：** 服务版本号，以数字及点号组成，形如1.0.1。 **约束限制：** 不涉及。 **取值范围：** 1.0.0 ~ 99.99.99，长度不超过8位。 **默认取值：** 前端可不传默认设置为1.0.0。
	Version *string `json:"version,omitempty"`

	// **参数解释：** 服务备注。 **约束限制：** 不涉及。 **取值范围：** 长度不可以超过512，不能包含大于号，小于号。 **默认取值：** 默认为空。
	Description *string `json:"description,omitempty"`

	// **参数解释：** 推理服务类型。 **约束限制：** 不涉及。 **取值范围：** - REAL_TIME：在线服务。 - ASYNC_REAL_TIME：异步服务。 **默认取值：** 不涉及。
	Type string `json:"type"`

	// **参数解释：** 部署类型。 **约束限制：** 不涉及。 **取值范围：** - SINGLE：单机单卡。 - MULTI：多机多卡。 **默认取值：** 不涉及。
	DeployType string `json:"deploy_type"`

	// **参数解释：** 服务实例组配置。 **约束限制：** 仅创建服务时group_configs可传[]空数组，group_configs的最大元素数量为1。
	GroupConfigs []GroupConfig `json:"group_configs"`

	RuntimeConfig *RuntimeConfig `json:"runtime_config"`

	UpgradeConfig *UpgradeConfig `json:"upgrade_config,omitempty"`

	// **参数解释：** 日志策略。 **约束限制：** 不涉及。 **取值范围：** - POOL：使用资源池日志插件配置的日志流。 - AUTO_CREATE：自动创建日志流。 - DEFAULT: 由系统决定日志策略 **默认取值：** AUTO_CREATE：自动创建日志流。
	LtsStrategy *string `json:"lts_strategy,omitempty"`

	// **参数解释：** 日志配置。 **约束限制：** 当开启LTS日志的时候，STDOUT类型为必填。 数量上限为2个。
	LogConfigs *[]LtsConfig `json:"log_configs,omitempty"`

	// **参数解释：** 服务标签。 **约束限制：** 上限20个。
	Tags *[]ServiceCreateRequestTags `json:"tags,omitempty"`

	// **参数解释：** 工作空间ID。 **约束限制：** 不涉及。 **取值范围：** - 0：默认空间ID。 - 由数字和小写字母组成的32位字符：其他空间ID，可参考[工作空间创建](CreateWorkspace.xml)。 **默认取值：** 不涉及。
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// **参数解释：** 定时停止配置。 **约束限制：** 最多支持一个定时任务。
	Schedule *[]ScheduleConfig `json:"schedule,omitempty"`

	// **参数解释：** 该参数值由英文逗号隔开的协议、端口号、地址组成，比如：[http,8080,metrics]，其中地址长度不超过255 ，且需要与镜像给定的协议、地址、端口一致，否则指标无法上报。 **约束限制：** 长度不超过255。 **取值范围：** - 协议范围：http/https。 - 端口范围：1-65535。 - 地址范围：仅包含字母、数字、点号（.）、中划线（-)、下划线（_）、斜杠（/）的路径，非斜杠（/）开头。 **默认取值：** 不涉及。
	CustomMetricsPath string `json:"custom_metrics_path"`

	// **参数解释：** 服务部署超时时间，integer类型，取值在1~300 （860版本该参数在服务层级做保留兼容）。 **约束限制：** 不涉及。 **取值范围：** [0, 300]。 **默认取值：** 不涉及。
	DeployTimeoutMinutes *int32 `json:"deploy_timeout_minutes,omitempty"`

	// **参数解释：** 模型类型。 **取值范围：** - TEXT_GENERATION：文本生成 - IMAGE_UNDERSTANDING：图像理解 - VIDEO_GENERATION：视频生成 - IMAGE_GENERATION：图像生成 - RERANK：重排序 - VECTOR_MODEL：向量模型 - EMBEDDING：Embedding嵌入
	TaskType *string `json:"task_type,omitempty"`

	// **参数解释：** 工作负载类型。 **取值范围：** - DEPLOYMENT：DEPLOYMENT类型 - LWS：LWS类型
	WorkloadType *string `json:"workload_type,omitempty"`
}

func (o ServiceCreateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServiceCreateRequest struct{}"
	}

	return strings.Join([]string{"ServiceCreateRequest", string(data)}, " ")
}
