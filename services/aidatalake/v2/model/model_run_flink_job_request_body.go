package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunFlinkJobRequestBody **参数解释**：启动Flink作业的请求参数，包含作业名称、端点信息、作业配置、资源配置等信息。 **约束限制**：不涉及。
type RunFlinkJobRequestBody struct {

	// **参数解释**：作业名称，用于标识Flink作业。 **约束限制**：不涉及。 **取值范围**：长度为1~128个字符，以字母或数字开头和结尾，中间可包含字母、数字、下划线和中划线。 **默认取值**：不涉及。
	Name string `json:"name"`

	// **参数解释**：作业运行的端点名称，用于指定作业执行环境。可通过[查询端点列表](ListV2Endpoints.xml)或[查询所有端点列表](ListV2AllEndpoints.xml)接口获取。 **约束限制**：不涉及。 **取值范围**：只能由英文小写字母、数字及中划线组成，以英文小写字母开头，以英文小写字母或数字结尾，且长度为1~63个字符。 **默认取值**：不涉及。
	EndpointName string `json:"endpoint_name"`

	JobConfig *FlinkJobArtifact `json:"job_config"`

	// **参数解释**：已部署作业ID，用于关联聚合Flink实例。 **约束限制**：不涉及。 **取值范围**：长度为1~64个字符，只能由英文字母、数字、下划线和中划线组成。 **默认取值**：不涉及。
	DeploymentId *string `json:"deployment_id,omitempty"`

	// **参数解释**：用户自定义委托名称，用于授权作业访问OBS等外部资源。 **约束限制**：须为当前租户下已存在的委托，系统管理委托不可用于提交作业。 **取值范围**：长度为1~64个字符，字符组成不限制。 **默认取值**：不指定时默认使用端点作业委托。
	JobAgency *string `json:"job_agency,omitempty"`

	// **参数解释**：用户自定义的Flink配置参数，为键值对形式。最多支持100个参数，每个值长度为1~1024个字符。不允许配置execution.checkpointing.dir、execution.checkpointing.savepoint-dir、high-availability.storageDir等系统保留参数，上述系统保留参数由平台根据端点配置自动填充并在查询响应中返回，无需用户配置。 **约束限制**：不涉及。
	FlinkConfig map[string]string `json:"flink_config,omitempty"`

	AutoscalingPolicy *AutoscalingPolicy `json:"autoscaling_policy,omitempty"`

	ImageConfig *FlinkJobImageConfig `json:"image_config,omitempty"`

	ResourceConfig *FlinkResourceConfig `json:"resource_config,omitempty"`

	ExecutionMode *FlinkJobExecutionMode `json:"execution_mode,omitempty"`

	RestoreStrategy *RestoreStrategy `json:"restore_strategy,omitempty"`

	LoggingConfig *FlinkLoggingConfig `json:"logging_config,omitempty"`

	// **参数解释**：标签，用于作业分类和筛选。 **约束限制**：标签数量不能超过16条。
	Labels *[]FlinkLabel `json:"labels,omitempty"`

	// **参数解释**：描述信息，用于说明作业用途。 **约束限制**：不涉及。 **取值范围**：长度为1~512个字符。 **默认取值**：不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释**：是否对该作业开启失败告警通知，用于按作业控制是否向客户发送告警消息。 **约束限制**：仅提交作业时指定，作业运行期间不支持修改。 **取值范围**：   - true：开启。作业运行失败或提交失败时，通过消息中心向客户发送告警通知。   - false：关闭。该作业失败时不发送告警通知。 **默认取值**：false，不指定时默认关闭（不发送该作业的失败告警通知），需要告警时显式指定为 true。
	AlertEnabled *bool `json:"alert_enabled,omitempty"`
}

func (o RunFlinkJobRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunFlinkJobRequestBody struct{}"
	}

	return strings.Join([]string{"RunFlinkJobRequestBody", string(data)}, " ")
}
