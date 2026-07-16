package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HpaRule 自动扩缩容规则返回体
type HpaRule struct {

	// **参数解释：** 自动扩缩容规则ID **取值范围：** 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释：** 自动扩缩容规则关联的策略ID **取值范围：** 不涉及。
	HpaId *string `json:"hpa_id,omitempty"`

	// **参数解释：** 自动扩缩容规则名 **取值范围：** 支持4-64个字符，可以包含字母、汉字、数字、中划线和下划线。
	Name string `json:"name"`

	// **参数解释：** 自动扩缩容类型。 **取值范围：** - CRON_HPA：定时扩缩容策略 - METRIC_HPA：指标扩缩容策略
	Type *string `json:"type,omitempty"`

	// **参数解释：** 定时自动扩缩容执行的cron表达式，不支持秒，从分钟开始设定 **取值范围：** 不涉及。
	Schedule *string `json:"schedule,omitempty"`

	// **参数解释：** 自动扩缩容目标实例数。 **取值范围：** 1-128
	TargetReplicas *int32 `json:"target_replicas,omitempty"`

	// **参数解释：** 自动扩缩容规则是否启用。 **取值范围：** - false - true
	Disable *string `json:"disable,omitempty"`

	// **参数解释：** 自动扩缩容规则状态。 **取值范围：** - CREATING：创建中 - CONFIG_SUCCESS：配置成功 - EXECUTE_SUCCESS：执行成功 - DELETED：已删除 - FAILED: 执行失败
	Status *string `json:"status,omitempty"`
}

func (o HpaRule) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HpaRule struct{}"
	}

	return strings.Join([]string{"HpaRule", string(data)}, " ")
}
