package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCustomRegexResponse Response Object
type ListCustomRegexResponse struct {

	// 总数
	Total *int64 `json:"total,omitempty"`

	// 正则列表
	Items          *[]CustomRegexItem `json:"items,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListCustomRegexResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCustomRegexResponse struct{}"
	}

	return strings.Join([]string{"ListCustomRegexResponse", string(data)}, " ")
}
