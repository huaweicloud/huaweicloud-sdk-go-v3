package model

import (
	"encoding/json"

	"strings"
)

type PreheatingTaskBody struct {
	// 输入示例：abc.com/image/1.png，多个URL之间需要用逗号分隔，目前不支持对目录的预热，单个url的长度限制为10240字符,单次最多输入1000个url。

	Urls []string `json:"urls"`
}

func (o PreheatingTaskBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "PreheatingTaskBody struct{}"
	}

	return strings.Join([]string{"PreheatingTaskBody", string(data)}, " ")
}
