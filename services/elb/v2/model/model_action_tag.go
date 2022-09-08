package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 要搜索的标签值。
type ActionTag struct {

	// 标签的键。最大长度127个unicode字符。
	Key string `json:"key"`

	// 标签的值列表。每个值最大长度255个unicode字符， value之间为或的关系。
	Values *[]string `json:"values,omitempty"`
}

func (o ActionTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ActionTag struct{}"
	}

	return strings.Join([]string{"ActionTag", string(data)}, " ")
}
