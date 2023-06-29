package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResourceTag 标签属性
type ResourceTag struct {

	// 键。最大长度36个unicode字符。 不能为空，只能包含大小写字母，数字，中划线“-”，下划线“_”。
	Key *string `json:"key,omitempty"`

	// 值。每个值最大长度43个unicode字符，删除时如果value有值按照key/value删除，如果value没值，则按照key删除。 不能为空，只能包含大小写字母，数字，中划线“-”，下划线“_”。
	Value *string `json:"value,omitempty"`
}

func (o ResourceTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceTag struct{}"
	}

	return strings.Join([]string{"ResourceTag", string(data)}, " ")
}
