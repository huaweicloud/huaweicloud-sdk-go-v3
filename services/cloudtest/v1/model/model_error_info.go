package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ErrorInfo struct {

	// 错误内容的描述
	ErrorContent *string `json:"error_content,omitempty"`

	// 错误索引的描述
	ErrorIndex *string `json:"error_index,omitempty"`

	// 错误点的描述
	ErrorPoint *string `json:"error_point,omitempty"`

	// 是否高亮标识
	HighLight *bool `json:"high_light,omitempty"`
}

func (o ErrorInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ErrorInfo struct{}"
	}

	return strings.Join([]string{"ErrorInfo", string(data)}, " ")
}
