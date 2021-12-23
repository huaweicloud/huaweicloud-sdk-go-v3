package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 值为列表的tag结构
type TagParams struct {
	// 键。最大长度127个unicode字符。 key不能为空。key不能为空或者空字符串，不能为空格，使用之前先trim 前后半角空格。

	Key string `json:"key"`
	// 值列表。每个值最大长度255个unicode字符，使用之前先trim 前后半角空格。\\*为系统保留字符，如果value是以\\*开头表示按照\\*后面的值全模糊匹配。不能只传入“\\*”。如果values为空列表，则表示any_value（查询任意value）。value之间为或的关系。

	Values []string `json:"values"`
}

func (o TagParams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagParams struct{}"
	}

	return strings.Join([]string{"TagParams", string(data)}, " ")
}
