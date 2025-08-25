package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ProtectedBranchBodyApiDto struct {

	// **参数解释：** 保护分支名称。 **取值范围** 字符串长度不少于1，不超过1000。
	Name string `json:"name"`

	// **参数解释：** 事件列表。
	Actions *[]ProtectedActionBasicApiDto `json:"actions,omitempty"`
}

func (o ProtectedBranchBodyApiDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProtectedBranchBodyApiDto struct{}"
	}

	return strings.Join([]string{"ProtectedBranchBodyApiDto", string(data)}, " ")
}
