package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VersionModelVersionUpdateDto struct {

	// **参数解释：**  唯一标识。  **约束限制：**  不涉及。  **取值范围：**  -9223372036854775808到9223372036854775807的整数。  **默认取值：**  不涉及。
	Id string `json:"id"`

	// **参数解释：**  迭代版本。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Iteration *int32 `json:"iteration,omitempty"`

	// **参数解释：**  版本号。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Version string `json:"version"`
}

func (o VersionModelVersionUpdateDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VersionModelVersionUpdateDto struct{}"
	}

	return strings.Join([]string{"VersionModelVersionUpdateDto", string(data)}, " ")
}
