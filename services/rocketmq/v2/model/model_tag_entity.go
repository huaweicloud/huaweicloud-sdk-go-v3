package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TagEntity struct {

	// **参数解释**： 标签键。 **约束限制**： - 不能为空。 - 对于同一个实例，Key值唯一。 - 长度为1~128个字符（中文也可以输入128个字符）。 - 由任意语种字母、数字、空格和字符组成，字符仅支持_ . : = + - @ - 首尾字符不能为空格。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Key string `json:"key"`

	// **参数解释**： 标签值。 **约束限制**： - 长度为0~255个字符（中文也可以输入255个字符）。 - 由任意语种字母、数字、空格和字符组成，字符仅支持_ . : = + - @ - 首尾字符不能为空格。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Value string `json:"value"`
}

func (o TagEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagEntity struct{}"
	}

	return strings.Join([]string{"TagEntity", string(data)}, " ")
}
