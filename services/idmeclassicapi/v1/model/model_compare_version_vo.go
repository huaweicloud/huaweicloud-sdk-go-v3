package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CompareVersionVo struct {

	// **参数解释：**  基础版本号。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	BasicVersion string `json:"basicVersion"`

	// **参数解释：**  对比版本号。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	CorrelationVersion string `json:"correlationVersion"`

	// **参数解释：**  实例ID。  **约束限制：**  不涉及。  **取值范围：**  -9223372036854775808到9223372036854775807的整数。  **默认取值：**  不涉及。
	Id string `json:"id"`
}

func (o CompareVersionVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CompareVersionVo struct{}"
	}

	return strings.Join([]string{"CompareVersionVo", string(data)}, " ")
}
