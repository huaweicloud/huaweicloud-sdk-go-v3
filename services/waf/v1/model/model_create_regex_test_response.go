package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateRegexTestResponse Response Object
type CreateRegexTestResponse struct {

	// 测试是否成功
	Success *bool `json:"success,omitempty"`

	// 测试结果消息
	Message *string `json:"message,omitempty"`

	// 匹配结果列表
	Matches        *[]string `json:"matches,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o CreateRegexTestResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRegexTestResponse struct{}"
	}

	return strings.Join([]string{"CreateRegexTestResponse", string(data)}, " ")
}
