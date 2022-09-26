package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FuncCode结构返回体。
type FuncCode struct {

	// 函数代码，如果不为空必须进行base64编码，为空时使用默认的代码。
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
