/*
 * FunctionGraph
 *
 * API v2
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListVersionAliasesResponse struct {
	// 函数版本别名列表
	Body           *[]ListVersionAliasResult `json:"body,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o ListVersionAliasesResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListVersionAliasesResponse", string(data)}, " ")
}
