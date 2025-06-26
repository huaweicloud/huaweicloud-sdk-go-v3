package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteResourceTags **参数解释**： 标签列表信息。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
type BatchDeleteResourceTags struct {

	// **参数解释**： 标签列表。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Tags []BatchDeleteResourceTag `json:"tags"`
}

func (o BatchDeleteResourceTags) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteResourceTags struct{}"
	}

	return strings.Join([]string{"BatchDeleteResourceTags", string(data)}, " ")
}
