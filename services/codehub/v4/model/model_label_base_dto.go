package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LabelBaseDto struct {

	// **参数解释：** 标签ID。 **取值范围：** 1-2147483647
	Id *int32 `json:"id,omitempty"`

	// **参数解释：** 标签名。 **取值范围：** 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释：** 标签颜色，以6位十六进制表示法给出，带有前导“#”符号（例如，#FFAABB）。 **取值范围：** 正则`^#([a-fA-F0-9]{6}|[a-fA-F0-9]{3})$`
	Color *string `json:"color,omitempty"`

	// **参数解释：** 描述。 **取值范围：** 不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释：** 字体颜色，以6位十六进制表示法给出，带有前导“#”符号（例如，#FFAABB）。 **取值范围：** 正则`^#([a-fA-F0-9]{6}|[a-fA-F0-9]{3})$`
	TextColor *string `json:"text_color,omitempty"`
}

func (o LabelBaseDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LabelBaseDto struct{}"
	}

	return strings.Join([]string{"LabelBaseDto", string(data)}, " ")
}
