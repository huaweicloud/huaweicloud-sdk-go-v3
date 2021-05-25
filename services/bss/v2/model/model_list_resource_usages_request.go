package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListResourceUsagesRequest struct {
	// 语言。中文：zh_CN英文：en_US缺省为zh_CN。

	XLanguage *string `json:"X-Language,omitempty"`
}

func (o ListResourceUsagesRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListResourceUsagesRequest struct{}"
	}

	return strings.Join([]string{"ListResourceUsagesRequest", string(data)}, " ")
}
