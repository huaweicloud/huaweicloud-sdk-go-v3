package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResponseTag struct {

	// **参数解释**： 标签键。 **取值范围**：   - 最大长度128个unicode字符， key不能为空。   - 同一资源的key不能重复。   - 可以包含的字符范围：  - 英文字母     - 数字     - 特殊字符：下划线（_）、点（.）、冒号（:）、加号（+）、中划线（-）、等号（=）     - 中文字符
	Key string `json:"key"`

	// **参数解释**： 标签值。 **取值范围**：   - 每个值最大长度255个unicode字符，value可以为空。   - 可以包含的字符范围：  - 英文字母     - 数字     - 特殊字符：下划线（_）、点（.）、冒号（:）、加号（+）、中划线（-）、等号（=）     - 中文字符
	Value string `json:"value"`
}

func (o ResponseTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResponseTag struct{}"
	}

	return strings.Join([]string{"ResponseTag", string(data)}, " ")
}
