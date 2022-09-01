package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 资源标签
type ResourceTag struct {

	// 键。最大长度36个unicode字符。 key不能为空。 搜索时不对此参数做校验 不能包含非打印字符\"=\"，“*”，“<”，“>”，“\\”，“,”，\"|\"，“/”。
	Key string `json:"key" xml:"key"`

	// 值列表。每个值最大长度255个unicode字符，如果values为空列表，则表示any_value。value之间为或的关系。不能包含非打印字符\"=\"，“*”，“<”，“>”，“\\”，“,”，\"|\"，“/”。
	Value string `json:"value" xml:"value"`
}

func (o ResourceTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceTag struct{}"
	}

	return strings.Join([]string{"ResourceTag", string(data)}, " ")
}
