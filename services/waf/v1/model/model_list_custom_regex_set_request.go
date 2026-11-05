package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCustomRegexSetRequest Request Object
type ListCustomRegexSetRequest struct {

	// 正则库名称（模糊匹配）
	Name *string `json:"name,omitempty"`

	Offset *int32 `json:"offset,omitempty"`

	Limit *int32 `json:"limit,omitempty"`
}

func (o ListCustomRegexSetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCustomRegexSetRequest struct{}"
	}

	return strings.Join([]string{"ListCustomRegexSetRequest", string(data)}, " ")
}
