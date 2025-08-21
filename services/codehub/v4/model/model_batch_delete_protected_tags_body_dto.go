package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchDeleteProtectedTagsBodyDto struct {

	// **参数解释：** 保护Tag名或通配符列表。 **约束限制：** 必填 **取值范围：** 不涉及 **默认取值：** 不涉及。
	Names *[]string `json:"names,omitempty"`
}

func (o BatchDeleteProtectedTagsBodyDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteProtectedTagsBodyDto struct{}"
	}

	return strings.Join([]string{"BatchDeleteProtectedTagsBodyDto", string(data)}, " ")
}
