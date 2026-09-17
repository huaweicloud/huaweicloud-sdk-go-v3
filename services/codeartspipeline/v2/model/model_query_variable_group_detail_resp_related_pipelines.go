package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QueryVariableGroupDetailRespRelatedPipelines struct {

	// **参数解释**： 流水线ID。 **取值范围**： 32位字符，由数字和字母组成。
	Id *string `json:"id,omitempty"`

	// **参数解释**： 流水线名称。 **取值范围**： 不涉及。
	Name *string `json:"name,omitempty"`
}

func (o QueryVariableGroupDetailRespRelatedPipelines) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryVariableGroupDetailRespRelatedPipelines struct{}"
	}

	return strings.Join([]string{"QueryVariableGroupDetailRespRelatedPipelines", string(data)}, " ")
}
