package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResourceTag 资源标签
type ResourceTag struct {

	// **参数解释**：标签的键。  **取值范围**：最大长度36个unicode字符。不能包含非打印字符\"=\"，“*”，“<”，“>”，“\\”，“,”，\"|\"，“/”。
	Key string `json:"key"`

	// **参数解释**：标签的值。  **取值范围**：最大长度255个unicode字符。不能包含非打印字符\"=\"，“*”，“<”，“>”，“\\”，“,”，\"|\"，“/”。
	Value string `json:"value"`
}

func (o ResourceTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceTag struct{}"
	}

	return strings.Join([]string{"ResourceTag", string(data)}, " ")
}
