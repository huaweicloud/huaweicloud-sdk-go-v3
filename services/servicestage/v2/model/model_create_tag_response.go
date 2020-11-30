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
type CreateTagResponse struct {
	// tag标签名称。
	Name           *string `json:"name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateTagResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateTagResponse", string(data)}, " ")
}
