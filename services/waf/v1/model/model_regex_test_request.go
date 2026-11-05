package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RegexTestRequest struct {

	// 待测试的正则表达式
	Pattern string `json:"pattern"`

	// 待匹配的文本内容
	Text *string `json:"text,omitempty"`
}

func (o RegexTestRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RegexTestRequest struct{}"
	}

	return strings.Join([]string{"RegexTestRequest", string(data)}, " ")
}
