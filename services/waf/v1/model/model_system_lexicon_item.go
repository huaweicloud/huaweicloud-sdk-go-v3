package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SystemLexiconItem 系统词典详情
type SystemLexiconItem struct {

	// 词典编码
	Label *string `json:"label,omitempty"`

	// 词典名称
	Name *string `json:"name,omitempty"`

	// 词典类型
	Type *string `json:"type,omitempty"`
}

func (o SystemLexiconItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SystemLexiconItem struct{}"
	}

	return strings.Join([]string{"SystemLexiconItem", string(data)}, " ")
}
