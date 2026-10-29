package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ParamInfo **参数解释：** 任务基本信息。
type ParamInfo struct {

	// **参数解释：** 代码仓地址。 **取值范围:** 不涉及。
	Url *string `json:"url,omitempty"`

	// **参数解释：** 任务本次扫描的代码仓分支。 **取值范围:** 不涉及。
	Branch *string `json:"branch,omitempty"`

	// **参数解释：** 任务本次扫描的规则语言。 **取值范围:** 不涉及。
	Language *string `json:"language,omitempty"`

	// **参数解释：** 任务配置的排除目录信息。 **取值范围:** 不涉及。
	ExcludeDir *string `json:"exclude_dir,omitempty"`

	// **参数解释：** 废弃字段。 **取值范围:** 不涉及。
	Encode *string `json:"encode,omitempty"`

	// **参数解释：** 任务的部分检查参数json格式信息。 **取值范围:** 不涉及。
	CompileConfig *string `json:"compile_config,omitempty"`

	// **参数解释：** 废弃字段。 **取值范围:** 不涉及。
	RuleTemplate *string `json:"rule_template,omitempty"`
}

func (o ParamInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ParamInfo struct{}"
	}

	return strings.Join([]string{"ParamInfo", string(data)}, " ")
}
