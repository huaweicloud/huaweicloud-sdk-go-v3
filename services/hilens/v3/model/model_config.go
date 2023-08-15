package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Config struct {

	// 配置项key，最大长度36个字符。不能为空，只能包含大小写字母，数字，中划线“-”，下划线“_”
	Key string `json:"key"`

	// 配置项value，每个值最大长度43个字符，删除时如果value有值按照key/value删除，如果value没值，则按照key删除。不能为空，只能包含大小写字母，数字，中划线“-”，下划线“_”
	Value string `json:"value"`
}

func (o Config) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Config struct{}"
	}

	return strings.Join([]string{"Config", string(data)}, " ")
}
