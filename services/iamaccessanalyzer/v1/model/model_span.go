package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Span 光标在策略文本中的范围。范围由开始位置（含）和结束位置（不含）组成。
type Span struct {
	Start *Position `json:"start"`

	End *Position `json:"end"`
}

func (o Span) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Span struct{}"
	}

	return strings.Join([]string{"Span", string(data)}, " ")
}
