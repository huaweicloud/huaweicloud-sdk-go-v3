package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ProjectProtectedBranchApiDto struct {

	// **参数解释：** 保护分支唯一标识。
	Id *int32 `json:"id,omitempty"`

	// **参数解释：** 保护分支名称 **取值范围：** 字符串长度不少于1，不超过1000。
	Name *string `json:"name,omitempty"`

	// **参数解释：** 事件列表。
	Actions *[]ProjectProtectedActionResultApiDto `json:"actions,omitempty"`
}

func (o ProjectProtectedBranchApiDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProjectProtectedBranchApiDto struct{}"
	}

	return strings.Join([]string{"ProjectProtectedBranchApiDto", string(data)}, " ")
}
