package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QueryParentDto struct {

	// **参数解释：**  子节点实例ID。  **约束限制：**  不涉及。  **取值范围：**  -9223372036854775808到9223372036854775807的整数。  **默认取值：**  不涉及。
	ChildId string `json:"childId"`
}

func (o QueryParentDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryParentDto struct{}"
	}

	return strings.Join([]string{"QueryParentDto", string(data)}, " ")
}
