package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListProjectSetsRequest struct {
	// 查询偏移

	Offset *int32 `json:"offset,omitempty"`
	// 查询数量

	Limit *int32 `json:"limit,omitempty"`
}

func (o ListProjectSetsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListProjectSetsRequest struct{}"
	}

	return strings.Join([]string{"ListProjectSetsRequest", string(data)}, " ")
}
