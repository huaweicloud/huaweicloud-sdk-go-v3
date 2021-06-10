package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListFreeResourcesRequest struct {
	// 语言。中文：zh_CN英文：en_US缺省为zh_CN。

	XLanguage *string `json:"X-Language,omitempty"`

	Body *ListFreeResourcesReq `json:"body,omitempty"`
}

func (o ListFreeResourcesRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListFreeResourcesRequest struct{}"
	}

	return strings.Join([]string{"ListFreeResourcesRequest", string(data)}, " ")
}
