package model

import (
	"encoding/json"

	"strings"
)

type EnvReq struct {
	// 环境的名称，支持英文，数字，下划线，且只能以英文字母开头。

	Name string `json:"name"`
	// 描述信息 > 中文字符必须为UTF-8或者unicode编码。

	Remark *string `json:"remark,omitempty"`
}

func (o EnvReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "EnvReq struct{}"
	}

	return strings.Join([]string{"EnvReq", string(data)}, " ")
}
