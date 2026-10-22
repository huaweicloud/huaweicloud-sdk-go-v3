package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BriefWorkspace struct {

	// **参数解释**：工作空间ID。 **取值范围**：长度为1~36的英文字母、数字和中划线的组合。
	Id *string `json:"id,omitempty"`

	// **参数解释**：工作空间名称。 **约束限制**：不涉及。 **取值范围**：长度为4~32的英文字母、数字、下划线和中划线的组合。 **默认取值**：不涉及。
	Name *string `json:"name,omitempty"`
}

func (o BriefWorkspace) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BriefWorkspace struct{}"
	}

	return strings.Join([]string{"BriefWorkspace", string(data)}, " ")
}
