package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PublicTag Tag 字段数据结构说明
type PublicTag struct {

	// 键。最大长度128个unicode字符。key不能为空。(搜索时不对此参数做校验) ，key不能为空或者空字符串，不能为空格，校验和使用之前先trim 前后空格。
	Key string `json:"key"`

	// - 值列表。每个值最大长度255个unicode字符，不能为空格，校验和使用之前先trim 前后空格。 - *为系统保留字符，value可为空但不可缺省。 - 如果里面的value是以\\*开头表示按照\\*后面的值全模糊匹配。 - 如果values为空列表，则表示any_value（查询任意value）。value之间为或的关系。
	Values []string `json:"values"`
}

func (o PublicTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublicTag struct{}"
	}

	return strings.Join([]string{"PublicTag", string(data)}, " ")
}
