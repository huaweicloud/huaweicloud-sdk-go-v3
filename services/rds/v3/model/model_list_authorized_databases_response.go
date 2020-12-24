/*
 * RDS
 *
 * API v3
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListAuthorizedDatabasesResponse struct {
	// 数据库及相关权限。
	Databases *[]DatabaseWithPrivilege `json:"databases,omitempty"`
	// 总数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListAuthorizedDatabasesResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListAuthorizedDatabasesResponse", string(data)}, " ")
}
