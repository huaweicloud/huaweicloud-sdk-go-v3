package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Dictionary 字典项
type Dictionary struct {

	// **参数解释**： 字典编码 **取值范围**： 字符长度1-256位
	Code *string `json:"code,omitempty"`

	// **参数解释**： 字典值（单个值） **取值范围**： 字符长度1-65535位
	Value *string `json:"value,omitempty"`

	// **参数解释**： 字典值（多个值） **取值范围**：
	Values *[]string `json:"values,omitempty"`
}

func (o Dictionary) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Dictionary struct{}"
	}

	return strings.Join([]string{"Dictionary", string(data)}, " ")
}
