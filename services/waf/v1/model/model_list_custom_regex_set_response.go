package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCustomRegexSetResponse Response Object
type ListCustomRegexSetResponse struct {

	// 总数
	Total *int64 `json:"total,omitempty"`

	// 正则库列表
	Items          *[]CustomRegexSetItem `json:"items,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ListCustomRegexSetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCustomRegexSetResponse struct{}"
	}

	return strings.Join([]string{"ListCustomRegexSetResponse", string(data)}, " ")
}
