package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CustomizeResourceTag struct {

	// **参数解释：** Key值。 **约束限制：** 不涉及 **取值范围：** - 不能为空且首尾不能包含空格，最多支持128个字符 - 可用UTF-8格式表示的任意语种字母、数字和空格 - 支持部分特殊字符：_.:=+-@ - 不能以\"\\_sys\\_\"开头  **默认取值：** 不涉及
	Key *string `json:"key,omitempty"`

	// **参数解释：** Value值列表。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Values *[]string `json:"values,omitempty"`
}

func (o CustomizeResourceTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CustomizeResourceTag struct{}"
	}

	return strings.Join([]string{"CustomizeResourceTag", string(data)}, " ")
}
