package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Tag 标签
type Tag struct {

	// 标签的键，不能为空，最大长度128个unicode字符。大小写字母，数字，可以包含中划线“-”，下划线“_”，不能包含以下字符“=”,“*”,“<”,“>”,“\\”,“,”,“|”,“/”。
	Key string `json:"key"`

	// 标签的值，最大长度43个unicode字符。大小写字母，数字，可以包含中划线“-”，下划线“_”，不能包含以下字符“=”,“*”,“<”,“>”,“\\”,“,”,“|”,“/”。
	Value *string `json:"value,omitempty"`
}

func (o Tag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Tag struct{}"
	}

	return strings.Join([]string{"Tag", string(data)}, " ")
}
