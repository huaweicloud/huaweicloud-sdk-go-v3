package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResourceTag CCE资源标签
type ResourceTag struct {

	// **参数解释：** Key值。 **约束限制：** 不涉及 **取值范围：** - 不能为空且首尾不能包含空格，最多支持128个字符 - 可用UTF-8格式表示的任意语种字母、数字和空格 - 支持部分特殊字符：_.:=+-@ - 不能以\"\\_sys\\_\"开头  **默认取值：** 不涉及
	Key *string `json:"key,omitempty"`

	// Value值。 - 可以为空但不能缺省，最多支持255个字符 - 可用UTF-8格式表示的汉字、字母、数字和空格 - 支持部分特殊字符：_.:/=+-@
	Value *string `json:"value,omitempty"`
}

func (o ResourceTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceTag struct{}"
	}

	return strings.Join([]string{"ResourceTag", string(data)}, " ")
}
