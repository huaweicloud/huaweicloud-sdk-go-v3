/*
 * ServiceStage
 *
 * ServiceStage的API,包括应用管理和仓库授权管理
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListEnvironmentsResponse struct {
	// 环境总数。
	Count *int32 `json:"count,omitempty"`
	// 环境列表。
	Environments   *[]Environment `json:"environments,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListEnvironmentsResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListEnvironmentsResponse", string(data)}, " ")
}
