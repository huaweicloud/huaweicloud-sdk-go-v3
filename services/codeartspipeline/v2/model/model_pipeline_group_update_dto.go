package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PipelineGroupUpdateDto PipelineGroupUpdateDTO
type PipelineGroupUpdateDto struct {

	// **参数解释**： 流水线分组名。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Name string `json:"name"`

	// **参数解释**： 流水线分组ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Id string `json:"id"`
}

func (o PipelineGroupUpdateDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PipelineGroupUpdateDto struct{}"
	}

	return strings.Join([]string{"PipelineGroupUpdateDto", string(data)}, " ")
}
