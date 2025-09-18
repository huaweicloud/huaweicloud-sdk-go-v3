package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PipelineGroupCreateDto PipelineGroupCreateDTO
type PipelineGroupCreateDto struct {

	// **参数解释**： 流水线分组名。 **约束限制**： 不涉及。 **取值范围**： 32位字符，由数字和字母组成。 **默认取值**： 不涉及。
	Name string `json:"name"`

	// **参数解释**： CodeArts项目ID。 **约束限制**： 不涉及。 **取值范围**： 32位字符，由数字和字母组成。 **默认取值**： 不涉及。
	ProjectId string `json:"project_id"`

	// **参数解释**： 父分组ID。 **约束限制**： 不涉及。 **取值范围**： 32位字符，由数字和字母组成。 **默认取值**： 不涉及。
	ParentId *string `json:"parent_id,omitempty"`
}

func (o PipelineGroupCreateDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PipelineGroupCreateDto struct{}"
	}

	return strings.Join([]string{"PipelineGroupCreateDto", string(data)}, " ")
}
