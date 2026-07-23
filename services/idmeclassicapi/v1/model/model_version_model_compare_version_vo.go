package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VersionModelCompareVersionVo struct {

	// **参数解释：**  主对象ID，用于定位待对比版本所属的主对象。  **约束限制：**  不涉及。  **取值范围：**  -9223372036854775808到9223372036854775807的整数。  **默认取值：**  不涉及。
	Id string `json:"id"`

	// **参数解释：**  基础版本号，作为对比的基准版本。格式为“大版本.迭代版本”，如“A.1”。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	BasicVersion string `json:"basicVersion"`

	// **参数解释：**  对比版本号，与基础版本进行对比的目标版本。格式为“大版本.迭代版本”，如“B.2”。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	CorrelationVersion string `json:"correlationVersion"`
}

func (o VersionModelCompareVersionVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VersionModelCompareVersionVo struct{}"
	}

	return strings.Join([]string{"VersionModelCompareVersionVo", string(data)}, " ")
}
