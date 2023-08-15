package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LogContents 参数集
type LogContents struct {

	// 日志原数据。
	Content *string `json:"content,omitempty"`

	// 日志单行序列号。
	LineNum *string `json:"line_num,omitempty"`

	// 该条日志包含的 labels。
	Labels map[string]string `json:"labels,omitempty"`
}

func (o LogContents) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LogContents struct{}"
	}

	return strings.Join([]string{"LogContents", string(data)}, " ")
}
