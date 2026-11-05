package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCustomRegexRequest Request Object
type ListCustomRegexRequest struct {

	// 正则库ID
	RegexSetId string `json:"regex_set_id"`

	// 正则名称（模糊匹配）
	Name *string `json:"name,omitempty"`

	Offset *int32 `json:"offset,omitempty"`

	Limit *int32 `json:"limit,omitempty"`
}

func (o ListCustomRegexRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCustomRegexRequest struct{}"
	}

	return strings.Join([]string{"ListCustomRegexRequest", string(data)}, " ")
}
