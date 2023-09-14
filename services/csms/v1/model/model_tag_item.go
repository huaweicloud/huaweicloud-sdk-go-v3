package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TagItem struct {

	// 标签的名称。 同一个凭据，一个标签键只能对应一个标签值；不同的凭据可以使用相同的标签键。 用户最多可以给单个凭据添加20个标签。  约束：取值范围为1到128个字符，满足正则匹配\"^((?!\\\\s)(?!_sys_)[\\\\p{L}\\\\p{Z}\\\\p{N}_.:=+\\\\-@]*)(?<!\\\\s)$\"
	Key string `json:"key"`

	// 标签的值。  约束：取值范围不超过255个字符，满足正则匹配\"^([\\\\p{L}\\\\p{Z}\\\\p{N}_.:\\/=+\\\\-@]*)$\"
	Value *string `json:"value,omitempty"`
}

func (o TagItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagItem struct{}"
	}

	return strings.Join([]string{"TagItem", string(data)}, " ")
}
