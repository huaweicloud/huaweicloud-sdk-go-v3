package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 参数集
type StructLogContents struct {
	// 日志原数据。

	LogContent *string `json:"log_content,omitempty"`
	// 日志单行序列号。

	LineNum *string `json:"line_num,omitempty"`
}

func (o StructLogContents) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StructLogContents struct{}"
	}

	return strings.Join([]string{"StructLogContents", string(data)}, " ")
}
