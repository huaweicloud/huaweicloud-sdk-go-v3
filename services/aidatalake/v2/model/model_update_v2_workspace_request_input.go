package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateV2WorkspaceRequestInput **参数解释**：更新工作空间的请求体。
type UpdateV2WorkspaceRequestInput struct {

	// **参数解释**：工作空间名称。 **约束限制**：不涉及。 **取值范围**：长度为4~32的英文字母、数字、下划线和中划线的组合。 **默认取值**：不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**：用户输入的描述。 **约束限制**：不涉及。 **取值范围**：不允许输入感叹号、小于号、大于号、等号、与号、双引号、单引号、艾特符、井号、百分号、分号、左括号、右括号、竖线、星号、反斜杠。长度为0~255个字符。 **默认取值**：不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释**：标签列表，每个元素为一个V2ResourceTag对象，包含key（标签键）和value（标签值）字段。 **约束限制**：最多20个标签。
	Tags *[]V2ResourceTag `json:"tags,omitempty"`
}

func (o UpdateV2WorkspaceRequestInput) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateV2WorkspaceRequestInput struct{}"
	}

	return strings.Join([]string{"UpdateV2WorkspaceRequestInput", string(data)}, " ")
}
