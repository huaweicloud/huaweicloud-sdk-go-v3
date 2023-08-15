package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LtsLogErrorDetail struct {

	// 节点ID。
	NodeId *string `json:"node_id,omitempty"`

	// 执行时间。
	Time *string `json:"time,omitempty"`

	// 日志级别。
	Level *string `json:"level,omitempty"`

	// 错误日志内容。
	Content *string `json:"content,omitempty"`

	// 日志单行序列号。
	LineNum *string `json:"line_num,omitempty"`
}

func (o LtsLogErrorDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LtsLogErrorDetail struct{}"
	}

	return strings.Join([]string{"LtsLogErrorDetail", string(data)}, " ")
}
