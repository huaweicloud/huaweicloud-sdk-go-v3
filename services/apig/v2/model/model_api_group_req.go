package model

import (
	"encoding/json"

	"strings"
)

type ApiGroupReq struct {
	// API分组的名称。 由中文、英文字母、数字、“_”组成，且只能以英文或中文开头。 > 中文字符必须为UTF-8或者unicode编码。
	Name string `json:"name"`
	// API分组描述。 > 中文字符必须为UTF-8或者unicode编码。
	Remark *string `json:"remark,omitempty"`
}

func (o ApiGroupReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ApiGroupReq struct{}"
	}

	return strings.Join([]string{"ApiGroupReq", string(data)}, " ")
}
