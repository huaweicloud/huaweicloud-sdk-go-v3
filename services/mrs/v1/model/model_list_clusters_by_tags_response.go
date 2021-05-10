package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListClustersByTagsResponse struct {
	// 资源列表

	Resources *[]MrsResource `json:"resources,omitempty"`
	// 资源总数

	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListClustersByTagsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListClustersByTagsResponse struct{}"
	}

	return strings.Join([]string{"ListClustersByTagsResponse", string(data)}, " ")
}
