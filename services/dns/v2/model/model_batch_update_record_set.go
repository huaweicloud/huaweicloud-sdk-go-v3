package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchUpdateRecordSet struct {

	// **参数解释：** 记录集ID。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Id string `json:"id"`

	// **参数解释：** 记录集的描述信息。 **约束限制：** 不涉及。 **取值范围：** 长度不超过255个字符。 **默认取值：** 不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释：** 记录集的有效缓存时间，以秒为单位。 **约束限制：** 不涉及。 **取值范围：** 1~2147483647。 **默认取值：** 300
	Ttl *int32 `json:"ttl,omitempty"`

	// **参数解释：** 解析记录的权重。 **约束限制：** 在相同域名、类型、线路下的解析记录，规则如下： - 全部设置权重，或全部不设置权重。 - 当不设置权重时，只能创建一个解析记录。 - 当设置权重时，最多能创建20个解析记录。  **取值范围：** 0~1000。 - 当weight=null时，表示该解析记录不设置权重。 - 当weight=0，表示备用域名解析记录。 - 当weight>0，表示主用域名解析记录。  **默认取值：** null。
	Weight *int32 `json:"weight,omitempty"`

	// **参数解释：** 解析记录的值。不同类型解析记录对应的值的规则不同。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Records []string `json:"records"`
}

func (o BatchUpdateRecordSet) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateRecordSet struct{}"
	}

	return strings.Join([]string{"BatchUpdateRecordSet", string(data)}, " ")
}
