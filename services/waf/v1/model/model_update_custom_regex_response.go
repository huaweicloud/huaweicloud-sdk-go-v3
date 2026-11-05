package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateCustomRegexResponse Response Object
type UpdateCustomRegexResponse struct {

	// 正则ID
	Id *string `json:"id,omitempty"`

	// 正则名称
	Name *string `json:"name,omitempty"`

	// 正则表达式
	RegexExpression *string `json:"regex_expression,omitempty"`

	// 创建时间（毫秒时间戳）
	CreateTime *int64 `json:"create_time,omitempty"`

	// 更新时间（毫秒时间戳）
	UpdateTime     *int64 `json:"update_time,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o UpdateCustomRegexResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCustomRegexResponse struct{}"
	}

	return strings.Join([]string{"UpdateCustomRegexResponse", string(data)}, " ")
}
