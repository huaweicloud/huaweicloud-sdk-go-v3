package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QueryResourceTagItem struct {

	// 标签键。最大长度36个unicode字符，key不能为空。 字符集：0-9，A-Z，a-z，“_”，“-”，中文。
	Key *string `json:"key,omitempty" xml:"key"`

	// 标签值。最大长度43个unicode字符，可以为空字符串。 字符集：0-9，A-Z，a-z，“_”，“.”，“-”，中文。
	Value *string `json:"value,omitempty" xml:"value"`
}

func (o QueryResourceTagItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryResourceTagItem struct{}"
	}

	return strings.Join([]string{"QueryResourceTagItem", string(data)}, " ")
}
