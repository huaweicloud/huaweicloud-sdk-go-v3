package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TagsResp
type TagsResp struct {

	// 键。  key最大长度36个字符。  key不能为空字符串。  key只能由中文，字母，数字，“-”，“_”组成。
	Key *string `json:"key,omitempty"`

	// 值列表。  value最大长度43个字符。  value可以为空字符串。  key只能由中文，字母，数字，“-”，“_”组成。
	Values *string `json:"values,omitempty"`
}

func (o TagsResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagsResp struct{}"
	}

	return strings.Join([]string{"TagsResp", string(data)}, " ")
}
