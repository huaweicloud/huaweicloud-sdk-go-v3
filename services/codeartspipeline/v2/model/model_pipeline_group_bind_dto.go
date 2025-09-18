package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PipelineGroupBindDto **参数解释**： 流水线分组参数详情。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
type PipelineGroupBindDto struct {

	// **参数解释**： 分组ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	GroupId string `json:"group_id"`

	// **参数解释**： 流水线集合。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Pipelines []PipelineGroupBindDtoPipelines `json:"pipelines"`
}

func (o PipelineGroupBindDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PipelineGroupBindDto struct{}"
	}

	return strings.Join([]string{"PipelineGroupBindDto", string(data)}, " ")
}
