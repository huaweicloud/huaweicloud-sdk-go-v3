package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 项目下的某资源的标签列表。
type ListTag struct {
	// 键。最大长度36个unicode字符。 key不能为空。不能包含非打印字符\"=\"，“*”，“<”，“>”，“\\”，“,”，\"|\"，“/”。

	Key string `json:"key"`
	// 值列表。每个值最大长度43个unicode字符，可以为空字符串。 如果values为空列表，则表示any_value。value之间为或的关系。

	Values []string `json:"values"`
}

func (o ListTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTag struct{}"
	}

	return strings.Join([]string{"ListTag", string(data)}, " ")
}
