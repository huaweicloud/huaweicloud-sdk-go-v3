package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TmsTagForDeletion 删除标签的标签结构体，value可以不填，当value不填，表示删除匹配到key的标签。
type TmsTagForDeletion struct {

	// **参数解释：** Tms标签的key。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Key string `json:"key"`

	// **参数解释：** Tms标签的value，非必填。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Value *string `json:"value,omitempty"`
}

func (o TmsTagForDeletion) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TmsTagForDeletion struct{}"
	}

	return strings.Join([]string{"TmsTagForDeletion", string(data)}, " ")
}
