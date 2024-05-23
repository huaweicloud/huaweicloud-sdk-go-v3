package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Tag struct {

	// 键。最大长度128个unicode字符。 key不能为空。(搜索时不对此参数做字符集校验)，key不能为空或者空字符串，不能为空格，校验和使用之前先trim 前后半角空格。
	Key string `json:"key"`

	// 值列表。每个值最大长度255个unicode字符，校验和使用之前先trim 前后半角空格。value可为空数组但不可缺省。如果values为空列表，则表示any_value（查询任意value）。value之间为或的关系。(搜索时不对此参数做字符集校验，只做长度校验)
	Values []string `json:"values"`
}

func (o Tag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Tag struct{}"
	}

	return strings.Join([]string{"Tag", string(data)}, " ")
}
