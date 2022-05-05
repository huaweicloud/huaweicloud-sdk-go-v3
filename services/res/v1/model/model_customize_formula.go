package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 自定义公式
type CustomizeFormula struct {

	// 别名。
	Alias *string `json:"alias,omitempty"`

	// 公式。
	Formula *string `json:"formula,omitempty"`
}

func (o CustomizeFormula) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CustomizeFormula struct{}"
	}

	return strings.Join([]string{"CustomizeFormula", string(data)}, " ")
}
