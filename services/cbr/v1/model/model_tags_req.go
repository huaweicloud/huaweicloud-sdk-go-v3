package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type TagsReq struct {
	// 键。  最大长度127个unicode字符。  不允许为空字符串。  前后的空格会被丢弃。

	Key string `json:"key"`
	// 值列表。  values中最多包含10个value。  每个value最大长度255个unicode字符。前后的空格会被丢弃。  values中value不允许重复。  values中多个value之间是\"或\"的关系。  values允许为空列表，value允许为空字符串。  values如果为空列表，表示任意值。  \\*为系统保留字符，如果value是以\\*开头表示按照\\*后面的值全模糊匹配，不能只传入“\\*”。

	Values []string `json:"values"`
}

func (o TagsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagsReq struct{}"
	}

	return strings.Join([]string{"TagsReq", string(data)}, " ")
}
