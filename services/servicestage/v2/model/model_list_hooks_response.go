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
type ListHooksResponse struct {
	// hook列表。
	Hooks *[]Hook `json:"hooks,omitempty"`
}

func (o ListHooksResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListHooksResponse", string(data)}, " ")
}
