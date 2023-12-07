package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TagsMultiValue struct {

	// 资源标签键。最大长度36个unicode字符。key不能为空。（搜索时不对此参数做校验）。最多为10个，不能为空或者空字符串。且不能重复。
	Key string `json:"key"`

	// 资源标签值列表，每个值最大长度43个unicode字符，每个key下最多为10个，同一个key中values不能重复。  “*”为系统保留字符，如果value是以“*”开头表示按照“*”后面的值全模糊匹配。不能只传入“*”。 如果values为空列表但不可缺省，则表示any_value（查询任意value）。value之间为或的关系
	Values []string `json:"values"`
}

func (o TagsMultiValue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagsMultiValue struct{}"
	}

	return strings.Join([]string{"TagsMultiValue", string(data)}, " ")
}
