package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchTagActionTagOption struct {
	// 标签键。最大长度36个unicode字符，不能为null或者空字符串，不能为空格，校验和使用之前会自动过滤掉前后空格。 字符集：0-9，A-Z，a-z，“_”，“-”，中文。

	Key string `json:"key"`
	// 标签值。最大长度43个unicode字符，可以为空字符串，不能为空格，校验和使用之前会自动过滤掉前后空格。 字符集：0-9，A-Z，a-z，“_”，“.”，“-”，中文。 - “action”值为“create”时，该参数必选。 - “action”值为“delete”时，该参数可选。

	Value *string `json:"value,omitempty"`
}

func (o BatchTagActionTagOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchTagActionTagOption struct{}"
	}

	return strings.Join([]string{"BatchTagActionTagOption", string(data)}, " ")
}
