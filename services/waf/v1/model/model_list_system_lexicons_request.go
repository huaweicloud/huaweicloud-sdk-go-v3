package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSystemLexiconsRequest Request Object
type ListSystemLexiconsRequest struct {

	// **参数解释：** 语言 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** en-us
	XLanguage *string `json:"X-Language,omitempty"`

	// **参数解释：** 词典类型 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Type *string `json:"type,omitempty"`
}

func (o ListSystemLexiconsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSystemLexiconsRequest struct{}"
	}

	return strings.Join([]string{"ListSystemLexiconsRequest", string(data)}, " ")
}
