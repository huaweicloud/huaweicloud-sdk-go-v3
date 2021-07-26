package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListValueListRequest struct {
	// 页码

	Page *int32 `json:"page,omitempty"`
	// 每页的条数

	Pagesize *int32 `json:"pagesize,omitempty"`
}

func (o ListValueListRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListValueListRequest struct{}"
	}

	return strings.Join([]string{"ListValueListRequest", string(data)}, " ")
}
