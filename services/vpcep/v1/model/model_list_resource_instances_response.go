package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListResourceInstancesResponse struct {
	Resources *ResourceInstance `json:"resources,omitempty"`
	// 总记录数。

	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListResourceInstancesResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListResourceInstancesResponse struct{}"
	}

	return strings.Join([]string{"ListResourceInstancesResponse", string(data)}, " ")
}
