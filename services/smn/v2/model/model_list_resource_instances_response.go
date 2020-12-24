/*
 * SMN
 *
 * SMN Open API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListResourceInstancesResponse struct {
	// 返回的资源列表。
	Resources *[]TagResource `json:"resources,omitempty"`
	// 总记录数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListResourceInstancesResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListResourceInstancesResponse", string(data)}, " ")
}
