package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TagsFilter struct {

	// 标签键，最大长度128个unicode字符。 key不能为空。
	Key string `json:"key"`

	// 标签值列表，每个值最大长度255个unicode字符，校验和使用之前先trim 前后半角空格。value可为空数组但不可缺省。如果values为空列表，则表示any_value（查询任意value）。value之间为或的关系。。
	Values []string `json:"values"`
}

func (o TagsFilter) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagsFilter struct{}"
	}

	return strings.Join([]string{"TagsFilter", string(data)}, " ")
}
