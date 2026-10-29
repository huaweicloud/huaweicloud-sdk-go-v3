package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CustomParameter **参数解释：** 自定义参数。 **取值范围：** 不涉及。
type CustomParameter struct {

	// **参数解释：** 是否是私密参数。 **取值范围：** - true：私密参数。 - false：非私密参数。
	IsPrivate *bool `json:"isPrivate,omitempty"`

	// **参数解释：** 自定义参数的key。 **取值范围：** 长度为[1~128]个字符。
	Key *string `json:"key,omitempty"`

	// **参数解释：** 自定义参数的Value。 **取值范围：** 长度为[1~8192]个字符。
	Value *string `json:"value,omitempty"`
}

func (o CustomParameter) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CustomParameter struct{}"
	}

	return strings.Join([]string{"CustomParameter", string(data)}, " ")
}
