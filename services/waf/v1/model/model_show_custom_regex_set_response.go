package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCustomRegexSetResponse Response Object
type ShowCustomRegexSetResponse struct {

	// 正则库ID
	Id *string `json:"id,omitempty"`

	// 正则库名称
	Name *string `json:"name,omitempty"`

	// 正则库描述
	Description *string `json:"description,omitempty"`

	// 正则条目数量
	RegexCount *int32 `json:"regex_count,omitempty"`

	// 引用此正则库的规则列表
	ReferencedRules *[]ReferencedRule `json:"referenced_rules,omitempty"`

	// 创建时间（毫秒时间戳）
	CreateTime *int64 `json:"create_time,omitempty"`

	// 更新时间（毫秒时间戳）
	UpdateTime     *int64 `json:"update_time,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowCustomRegexSetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCustomRegexSetResponse struct{}"
	}

	return strings.Join([]string{"ShowCustomRegexSetResponse", string(data)}, " ")
}
