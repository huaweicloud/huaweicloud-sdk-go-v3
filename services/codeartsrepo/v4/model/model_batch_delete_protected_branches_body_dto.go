package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchDeleteProtectedBranchesBodyDto struct {

	// **参数解释：** 保护分支名或通配符列表。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Names *[]string `json:"names,omitempty"`
}

func (o BatchDeleteProtectedBranchesBodyDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteProtectedBranchesBodyDto struct{}"
	}

	return strings.Join([]string{"BatchDeleteProtectedBranchesBodyDto", string(data)}, " ")
}
