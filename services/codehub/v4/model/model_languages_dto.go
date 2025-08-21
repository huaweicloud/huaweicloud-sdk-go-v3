package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LanguagesDto struct {

	// **参数解释：** 颜色。
	Color *string `json:"color,omitempty"`

	// **参数解释：** 语言类型。
	Label *string `json:"label,omitempty"`

	// **参数解释：** 比重。
	Value *float64 `json:"value,omitempty"`
}

func (o LanguagesDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LanguagesDto struct{}"
	}

	return strings.Join([]string{"LanguagesDto", string(data)}, " ")
}
