package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PublicTags Tags字段数据结构说明
type PublicTags struct {

	// 键。最大长度128个unicode字符。key不能为空。不能包含非打印字符ASCII(0-31)，“=”,“*”,“<”,“>”,“\\”,“,”,“|”,“/”
	Key string `json:"key"`

	// 值。每个值最大长度255个unicode字符，删除时如果value有值按照key/value删除，如果value没值，则按照key删除，可以为空字符串。,不能包含非打印字符ASCII(0-31), “=”,“*”,“<”,“>”,“\\\\”,“,”,“|”,“/”
	Value string `json:"value"`
}

func (o PublicTags) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublicTags struct{}"
	}

	return strings.Join([]string{"PublicTags", string(data)}, " ")
}
