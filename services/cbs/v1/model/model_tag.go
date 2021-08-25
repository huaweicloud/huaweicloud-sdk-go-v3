package model

import (
	"encoding/json"

	"strings"
)

//
type Tag struct {
	// 必须要包含其中之一的答案标签id列表

	Should *[]string `json:"should,omitempty"`
}

func (o Tag) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "Tag struct{}"
	}

	return strings.Join([]string{"Tag", string(data)}, " ")
}
