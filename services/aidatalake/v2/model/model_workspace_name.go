package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkspaceName **参数解释**：工作空间名称。 **约束限制**：不涉及。 **取值范围**：长度为4~32的英文字母、数字、下划线和中划线的组合。 **默认取值**：不涉及。
type WorkspaceName struct {
}

func (o WorkspaceName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkspaceName struct{}"
	}

	return strings.Join([]string{"WorkspaceName", string(data)}, " ")
}
