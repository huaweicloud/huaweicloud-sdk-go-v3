package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateHpaRequestBody 修改定时扩缩容策略请求体
type UpdateHpaRequestBody struct {

	// **参数解释：** 自动扩缩容策略绑定的目标ID **取值范围：** 实例组ID **约束限制：** 不涉及。 **默认取值：** 不涉及。
	Id string `json:"id"`

	// **参数解释：** 自动扩缩容规则。 **约束限制：** 不涉及。
	HpaRules *[]HpaRules `json:"hpa_rules,omitempty"`

	// **参数解释：** 工作空间ID。 **取值范围：** - 0：默认空间ID - 由数字和小写字母组成的32位字符：其他空间ID，可参考[工作空间创建](CreateWorkspace.xml) **约束限制：** 不涉及。 **默认取值：** 不涉及。
	WorkspaceId *string `json:"workspace_id,omitempty"`
}

func (o UpdateHpaRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHpaRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateHpaRequestBody", string(data)}, " ")
}
