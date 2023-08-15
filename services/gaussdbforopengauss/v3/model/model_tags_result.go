package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TagsResult struct {

	// 标签键。最大长度36个unicode字符，key不能为空。 字符集：0-9，A-Z，a-z，“_”，“-”，中文。
	Key *string `json:"key,omitempty"`

	// 标签值。最大长度43个unicode字符，可以为空字符串。 字符集：0-9，A-Z，a-z，“_”，“.”，“-”，中文。
	Value *[]string `json:"value,omitempty"`
}

func (o TagsResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagsResult struct{}"
	}

	return strings.Join([]string{"TagsResult", string(data)}, " ")
}
