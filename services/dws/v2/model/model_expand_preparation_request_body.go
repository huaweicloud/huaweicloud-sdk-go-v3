package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExpandPreparationRequestBody **参数解释**： 扩容准备操作请求体。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
type ExpandPreparationRequestBody struct {

	// **参数解释**： 扩容节点数。 **取值范围：** 大于等于3。
	NumNode int32 `json:"num_node"`

	// **参数解释**： 是否是扩容准备重试。 **约束限制**： 不涉及。 **取值范围**： true：扩容准备重试； false：首次进行扩容准备； **默认取值**： false
	IsRetry bool `json:"is_retry"`
}

func (o ExpandPreparationRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExpandPreparationRequestBody struct{}"
	}

	return strings.Join([]string{"ExpandPreparationRequestBody", string(data)}, " ")
}
