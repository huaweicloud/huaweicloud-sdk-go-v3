package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeResourceTagOption **参数解释**：要添加或删除的资源标签。  **约束限制**：不涉及
type ChangeResourceTagOption struct {

	// **参数解释**：标签的键。  **约束限制**：不能传入空值。  **取值范围**：最大长度36个unicode字符。 不能包含非打印字符\"=\"，“*”，“<”，“>”，“\\”，“,”，\"|\"，“/”。  **默认取值**：不涉及
	Key string `json:"key"`

	// **参数解释**：标签的值，可以传入空字符串，表示任何值。  **约束限制**：注意删除标签也需要同时传入标签键值。  **取值范围**：最大长度255个unicode字符。 不能包含非打印字符\"=\"，“*”，“<”，“>”，“\\”，“,”，\"|\"，“/”。  **默认取值**：不涉及
	Value string `json:"value"`
}

func (o ChangeResourceTagOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeResourceTagOption struct{}"
	}

	return strings.Join([]string{"ChangeResourceTagOption", string(data)}, " ")
}
