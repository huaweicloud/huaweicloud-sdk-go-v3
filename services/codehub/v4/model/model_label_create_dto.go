package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LabelCreateDto struct {

	// **参数解释：** 标签名。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Name string `json:"name"`

	// **参数解释：** 标签颜色，以6位十六进制表示法给出，带有前导“#”符号（例如，#FFAABB）。 **约束限制：** 不涉及。 **取值范围：** 正则`^#([a-fA-F0-9]{6}|[a-fA-F0-9]{3})$` **默认取值：** 不涉及。
	Color *string `json:"color,omitempty"`

	// **参数解释：** 描述。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释：** 失效时间。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	ExpiresAt *string `json:"expires_at,omitempty"`
}

func (o LabelCreateDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LabelCreateDto struct{}"
	}

	return strings.Join([]string{"LabelCreateDto", string(data)}, " ")
}
