package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MultiValueTagInfo struct {

	// 键。最大长度128个unicode字符。 key不能为空
	Key *string `json:"key,omitempty"`

	// 值列表。每个值最大长度255个unicode字符
	Values *[]string `json:"values,omitempty"`
}

func (o MultiValueTagInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MultiValueTagInfo struct{}"
	}

	return strings.Join([]string{"MultiValueTagInfo", string(data)}, " ")
}
