package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSuggestionsResponse Response Object
type ListSuggestionsResponse struct {

	// 推荐问列表。
	Questions      *[]string `json:"questions,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListSuggestionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSuggestionsResponse struct{}"
	}

	return strings.Join([]string{"ListSuggestionsResponse", string(data)}, " ")
}
