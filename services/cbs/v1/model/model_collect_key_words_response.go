package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CollectKeyWordsResponse struct {
	// 指定时间范围内，用户问关键词列表。

	Keywords       *[]KeyWordsStat `json:"keywords,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o CollectKeyWordsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CollectKeyWordsResponse struct{}"
	}

	return strings.Join([]string{"CollectKeyWordsResponse", string(data)}, " ")
}