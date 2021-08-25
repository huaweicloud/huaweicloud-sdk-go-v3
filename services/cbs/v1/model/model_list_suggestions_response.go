package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListSuggestionsResponse struct {
	// 推荐问列表。

	Questions      *[]string `json:"questions,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListSuggestionsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListSuggestionsResponse struct{}"
	}

	return strings.Join([]string{"ListSuggestionsResponse", string(data)}, " ")
}
