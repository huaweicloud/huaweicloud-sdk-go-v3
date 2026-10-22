package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkspaceIdAdmin **参数解释**：工作空间ID。 **取值范围**：长度为1~36的英文字母、数字和中划线的组合。
type WorkspaceIdAdmin struct {
}

func (o WorkspaceIdAdmin) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkspaceIdAdmin struct{}"
	}

	return strings.Join([]string{"WorkspaceIdAdmin", string(data)}, " ")
}
