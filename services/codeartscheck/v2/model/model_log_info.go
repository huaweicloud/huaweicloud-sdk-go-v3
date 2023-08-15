package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LogInfo struct {

	// 日志标题
	DisplayName *string `json:"display_name,omitempty"`

	// 日志内容
	Log *string `json:"log,omitempty"`

	// 日志级别
	Level *string `json:"level,omitempty"`

	// 日志分析
	Analysis *string `json:"analysis,omitempty"`

	// 常见问题解答
	Faq *string `json:"faq,omitempty"`
}

func (o LogInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LogInfo struct{}"
	}

	return strings.Join([]string{"LogInfo", string(data)}, " ")
}
