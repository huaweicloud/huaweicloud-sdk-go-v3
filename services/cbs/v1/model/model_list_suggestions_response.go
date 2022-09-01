package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListSuggestionsResponse struct {

	// 推荐问列表。
	Questions      *[]string `json:"questions,omitempty" xml:"questions"`
	HttpStatusCode int       `json:"-"`
}

func (o ListSuggestionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSuggestionsResponse struct{}"
	}

	return strings.Join([]string{"ListSuggestionsResponse", string(data)}, " ")
}
