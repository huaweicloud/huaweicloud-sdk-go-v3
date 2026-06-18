package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SymbolNodeDto **参数解释：** 代码导航符号节点信息
type SymbolNodeDto struct {
	Def *DefEntryDto `json:"def,omitempty"`

	// **参数解释：** 子节点信息
	Children *[]SymbolNodeDto `json:"children,omitempty"`
}

func (o SymbolNodeDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SymbolNodeDto struct{}"
	}

	return strings.Join([]string{"SymbolNodeDto", string(data)}, " ")
}
