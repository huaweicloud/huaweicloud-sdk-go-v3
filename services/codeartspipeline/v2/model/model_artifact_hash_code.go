package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ArtifactHashCode struct {

	// **参数解释**： 哈希算法。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	HashType *string `json:"hash_type,omitempty"`

	// **参数解释**： 哈希值。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	HashValue *string `json:"hash_value,omitempty"`
}

func (o ArtifactHashCode) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ArtifactHashCode struct{}"
	}

	return strings.Join([]string{"ArtifactHashCode", string(data)}, " ")
}
