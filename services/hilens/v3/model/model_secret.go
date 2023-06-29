package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Secret 密钥信息
type Secret struct {

	// 密钥key，最大长度63个字符。不能为空，只能包含大小写字母，数字，中划线“-”，下划线“_”
	Key string `json:"key"`

	// 密钥value，每个值最大长度10000个字符，删除时如果value有值按照key/value删除，如果value没值，则按照key删除。不能为空，只能包含大小写字母，数字，中划线“-”，下划线“_”
	Value string `json:"value"`
}

func (o Secret) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Secret struct{}"
	}

	return strings.Join([]string{"Secret", string(data)}, " ")
}
