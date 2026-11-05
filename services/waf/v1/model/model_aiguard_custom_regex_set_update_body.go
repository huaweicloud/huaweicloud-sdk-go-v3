package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AiguardCustomRegexSetUpdateBody struct {

	// 正则库名称
	Name *string `json:"name,omitempty"`

	// 正则库描述
	Description *string `json:"description,omitempty"`
}

func (o AiguardCustomRegexSetUpdateBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AiguardCustomRegexSetUpdateBody struct{}"
	}

	return strings.Join([]string{"AiguardCustomRegexSetUpdateBody", string(data)}, " ")
}
