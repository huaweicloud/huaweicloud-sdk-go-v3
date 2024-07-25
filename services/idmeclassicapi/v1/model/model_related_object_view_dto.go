package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RelatedObjectViewDto struct {

	// **参数解释：**  数据实例ID。  **取值范围：**  -9223372036854775808到9223372036854775807的整数。  **默认取值：**  不涉及。
	ObjectId *string `json:"objectId,omitempty"`

	// **参数解释：**  关联的数据传输对象列表。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	RelatedList *[]BasicObjectQueryViewDto `json:"relatedList,omitempty"`

	// **参数解释：**  关联的数据实体列表。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	RelatedEntityList *[]BasicObjectQueryViewDto `json:"relatedEntityList,omitempty"`
}

func (o RelatedObjectViewDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RelatedObjectViewDto struct{}"
	}

	return strings.Join([]string{"RelatedObjectViewDto", string(data)}, " ")
}
