package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInferHraResponse Response Object
type ShowInferHraResponse struct {

	// **参数解释：** hra策略ID。 **取值范围：** 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释：** 工作空间ID。未创建工作空间时默认值为“0”，存在创建并使用的工作空间，以实际取值为准。 **取值范围：** 不涉及。
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// **参数解释：** HRA规则列表。
	HraRules *[]HraRuleResponse `json:"hra_rules,omitempty"`

	// **参数解释：** 用户控制的启用/禁用开关。 **取值范围：** true表示禁用，false表示启用 **约束限制：** 不涉及。 **默认取值：** 不涉及。
	Disable *bool `json:"disable,omitempty"`

	// **参数解释：** 策略状态。 **取值范围：** - CREATING 创建。 - ACTIVE 可用。 - INACTIVE 不可用。 - DELETED 已删除。 **约束限制：** 不涉及。 **默认取值：** 不涉及。
	PolicyStatus *string `json:"policy_status,omitempty"`

	// **参数解释：** 最小副本数，由于当前版本不会进行实质扩缩，因此该配置值无效。 **取值范围：** 1~128。
	MinReplicas *int32 `json:"min_replicas,omitempty"`

	// **参数解释：** 最大副本数，由于当前版本不会进行实质扩缩，因此该配置值无效。 **取值范围：** 1~128。
	MaxReplicas *int32 `json:"max_replicas,omitempty"`

	// **参数解释：** 扩缩容时间窗，由于当前版本不会进行实质扩缩，因此该配置值无效。 **取值范围：** 不涉及。
	ScaleWindow    *int32 `json:"scale_window,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowInferHraResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInferHraResponse struct{}"
	}

	return strings.Join([]string{"ShowInferHraResponse", string(data)}, " ")
}
