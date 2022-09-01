package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListInstanceTagsResult struct {

	// 标签键。最大长度36个unicode字符，key不能为空。 字符集：0-9，A-Z，a-z，“_”，“-”，中文。
	Key string `json:"key" xml:"key"`

	// 标签值。最大长度43个unicode字符，可以为空字符串。 字符集：0-9，A-Z，a-z，“_”，“.”，“-”，中文。
	Value string `json:"value" xml:"value"`
}

func (o ListInstanceTagsResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceTagsResult struct{}"
	}

	return strings.Join([]string{"ListInstanceTagsResult", string(data)}, " ")
}
