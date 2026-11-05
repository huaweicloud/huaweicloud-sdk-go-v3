package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AiguardCustomRegexSetBody struct {

	// 正则库名称，同project唯一
	Name string `json:"name"`

	// 正则库描述
	Description *string `json:"description,omitempty"`
}

func (o AiguardCustomRegexSetBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AiguardCustomRegexSetBody struct{}"
	}

	return strings.Join([]string{"AiguardCustomRegexSetBody", string(data)}, " ")
}
