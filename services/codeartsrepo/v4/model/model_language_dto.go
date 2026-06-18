package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LanguageDto struct {

	// **参数解释：** 语言名称。 **约束限制：** 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释：** 文件后缀名。 **约束限制：** 不涉及。
	ExtensionList *[]string `json:"extension_list,omitempty"`
}

func (o LanguageDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LanguageDto struct{}"
	}

	return strings.Join([]string{"LanguageDto", string(data)}, " ")
}
