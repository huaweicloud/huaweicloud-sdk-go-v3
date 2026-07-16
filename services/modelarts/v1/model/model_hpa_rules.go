package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HpaRules 自动扩缩容规则
type HpaRules struct {

	// **参数解释：** 自动扩缩容规则名。 **取值范围：** 支持4-64个字符，可以包含小写字母、数字和中划线，必须以小写字母开头，不能以中划线结尾。 **约束限制：** 不涉及。 **默认取值：** 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释：** 自动扩缩容id **取值范围：** 支持4-64个字符，可以包含字母、汉字、数字、中划线和下划线。 **约束限制：** 不涉及。 **默认取值：** 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释：** 自动扩缩容规则是否启用。 **取值范围：** - false：启动 - true：不启动 **约束限制：** 不涉及。 **默认取值：** 不涉及。
	Disable *bool `json:"disable,omitempty"`

	// **参数解释：** 自动扩缩容类型。 **取值范围：** - CRON_HPA：定时扩缩容策略 - METRIC_HPA：指标扩缩容策略 **约束限制：** 不涉及。 **默认取值：** 不涉及。
	Type string `json:"type"`

	// **参数解释：** 自动扩缩容类型。 **取值范围：** - CREATING：创建扩缩容策略 - CONFIG_SUCCESS：配置成功扩缩容策略 - EXECUTE_SUCCESS：执行成功扩缩容策略 - DELETED：删除扩缩容策略 - FAILED：失败扩缩容策略 **约束限制：** 不涉及。 **默认取值：** 不涉及。
	Status *string `json:"status,omitempty"`

	// **参数解释：** 自动扩缩容类型。 **取值范围：** - ADD：添加扩缩容策略规则 - UPDATE：修改扩缩容策略规则 - DELETE：删除扩缩容策略规则 **约束限制：** 不涉及。 **默认取值：** 不涉及。
	Operate *string `json:"operate,omitempty"`

	// **参数解释：** 定时自动扩缩容执行的cron表达式，不支持秒，从分钟开始设定 **取值范围：** 不涉及。 **约束限制：** 不涉及。 **默认取值：** 不涉及。
	Schedule string `json:"schedule"`

	// **参数解释：** 自动扩缩容目标实例数。 **取值范围：** 1-128，接口能接受浮点类型，后端会自动向下取整 **约束限制：** 不涉及。 **默认取值：** 不涉及。
	TargetReplicas int32 `json:"target_replicas"`

	// **参数解释：** 自动扩缩容最小实例数。 **取值范围：** 1-128，接口能接受浮点类型，后端会自动向下取整 **约束限制：** 不涉及。 **默认取值：** 不涉及。
	MinReplicas *int32 `json:"min_replicas,omitempty"`

	// **参数解释：** 自动扩缩容最大实例数。 **取值范围：** 1-128，接口能接受浮点类型，后端会自动向下取整 **约束限制：** 不涉及。 **默认取值：** 不涉及。
	MaxReplicas *int32 `json:"max_replicas,omitempty"`

	// **参数解释：** 自动扩缩容扩容冷却时间。 **取值范围：** 1-128，接口能接受浮点类型，后端会自动向下取整 **约束限制：** 不涉及。 **默认取值：** 不涉及。
	DownscaleWindow *int32 `json:"downscale_window,omitempty"`

	// **参数解释：** 自动扩缩容缩容冷却时间。 **取值范围：** 1-128，接口能接受浮点类型，后端会自动向下取整 **约束限制：** 不涉及。 **默认取值：** 不涉及。
	UpscaleWindow *int32 `json:"upscale_window,omitempty"`
}

func (o HpaRules) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HpaRules struct{}"
	}

	return strings.Join([]string{"HpaRules", string(data)}, " ")
}
