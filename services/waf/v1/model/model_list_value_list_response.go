package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListValueListResponse struct {
	// 引用表条数

	Total *int32 `json:"total,omitempty"`
	// 引用表列表

	Items          *[]ValueListResponseBody `json:"items,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ListValueListResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListValueListResponse struct{}"
	}

	return strings.Join([]string{"ListValueListResponse", string(data)}, " ")
}
