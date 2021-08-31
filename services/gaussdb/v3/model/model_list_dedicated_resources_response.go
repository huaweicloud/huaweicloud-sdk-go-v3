package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListDedicatedResourcesResponse struct {
	// 专属资源池信息

	Resources *[]DedicatedResource `json:"resources,omitempty"`
	// 专属资源池数量

	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListDedicatedResourcesResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListDedicatedResourcesResponse struct{}"
	}

	return strings.Join([]string{"ListDedicatedResourcesResponse", string(data)}, " ")
}
