package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TagOperationDto struct {

	// **参数解释：**  数据实例ID。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	ObjectId string `json:"objectId"`

	// **参数解释：**  标签ID。  **约束限制：**  不涉及。  **取值范围：**  -9223372036854775808到9223372036854775807的整数。  **默认取值：**  不涉及。
	TagId string `json:"tagId"`
}

func (o TagOperationDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagOperationDto struct{}"
	}

	return strings.Join([]string{"TagOperationDto", string(data)}, " ")
}
