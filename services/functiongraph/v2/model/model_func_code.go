package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FuncCode结构返回体。
type FuncCode struct {
	// 函数代码，当CodeTye为inline/zip/jar时必选，且代码必须要进行base64编码。

	File *string `json:"file,omitempty"`
	// 函数代码链接。

	Link *string `json:"link,omitempty"`
}

func (o FuncCode) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FuncCode struct{}"
	}

	return strings.Join([]string{"FuncCode", string(data)}, " ")
}
