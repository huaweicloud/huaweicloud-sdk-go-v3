package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateHpaRequestBody 创建定时扩缩容策略请求体
type CreateHpaRequestBody struct {

	// **参数解释：** 自动扩缩容类型。 **取值范围：** - CRON_HPA：定时扩缩容策略 - METRIC_HPA：指标扩缩容策略 **约束限制：** 不涉及。 **默认取值：** 不涉及。
	Type string `json:"type"`

	// **参数解释：** 自动扩缩容策略绑定的目标ID **取值范围：** 实例组ID **约束限制：** 不涉及。 **默认取值：** 不涉及。
	TargetResourceId string `json:"target_resource_id"`

	// **参数解释：** 自动扩缩容策略绑定的目标类型。 **取值范围：** - GROUP：实例组 **约束限制：** 不涉及。 **默认取值：** 不涉及。
	TargetResourceType string `json:"target_resource_type"`

	// **参数解释：** 自动扩缩容规则。 **约束限制：** 不涉及。
	HpaRules []HpaRules `json:"hpa_rules"`

	// **参数解释：** 工作空间ID。 **取值范围：** - 0：默认空间ID - 由数字和小写字母组成的32位字符：其他空间ID，可参考[工作空间创建](CreateWorkspace.xml) **约束限制：** 不涉及。 **默认取值：** 不涉及。
	WorkspaceId *string `json:"workspace_id,omitempty"`
}

func (o CreateHpaRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateHpaRequestBody struct{}"
	}

	return strings.Join([]string{"CreateHpaRequestBody", string(data)}, " ")
}
