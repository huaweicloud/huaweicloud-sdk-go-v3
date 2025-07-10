package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PathRule 路径规则信息,限制对应路径及子路径
type PathRule struct {

	// 完整路径： 1. 名称允许可见字符或空格，不可为全空格。 2. 长度0~512个字符。
	Path *string `json:"path,omitempty"`
}

func (o PathRule) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PathRule struct{}"
	}

	return strings.Join([]string{"PathRule", string(data)}, " ")
}
