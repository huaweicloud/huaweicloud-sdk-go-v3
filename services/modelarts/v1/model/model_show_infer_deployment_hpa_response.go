package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInferDeploymentHpaResponse Response Object
type ShowInferDeploymentHpaResponse struct {

	// **参数解释：** 自动扩缩容策略ID **取值范围：** 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释：** 自动扩缩容策略名称 **取值范围：** 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释：** 自动扩缩容策略绑定的目标ID **取值范围：** 实例组ID
	TargetResourceId *string `json:"target_resource_id,omitempty"`

	// **参数解释：** 自动扩缩容策略绑定的目标类型。 **取值范围：** - GROUP：实例组
	TargetResourceType *string `json:"target_resource_type,omitempty"`

	// **参数解释：** 自动扩缩容最小实例数。 **取值范围：** 1-128
	MinReplicas *int32 `json:"min_replicas,omitempty"`

	// **参数解释：** 自动扩缩容最大实例数。 **取值范围：** 1-128
	MaxReplicas *int32 `json:"max_replicas,omitempty"`

	// 参数解释：** 自动扩缩容策略状态。 **取值范围：** - INACTIVE：不启用 - ACTIVE：配置成功 - DELETED：已删除
	Status *string `json:"status,omitempty"`

	// **参数解释：** 工作空间ID。 **取值范围：** - 0：默认空间ID - 由数字和小写字母组成的32位字符：其他空间ID，可参考[工作空间创建](CreateWorkspace.xml)
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// **参数解释：** 自动扩缩容规则列表
	HpaRules       *[]HpaRule `json:"hpa_rules,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ShowInferDeploymentHpaResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInferDeploymentHpaResponse struct{}"
	}

	return strings.Join([]string{"ShowInferDeploymentHpaResponse", string(data)}, " ")
}
