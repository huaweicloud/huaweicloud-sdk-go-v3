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

// Request Object
type DeleteAuthorizeRequest struct {
	Name string `json:"name"`
}

func (o DeleteAuthorizeRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeleteAuthorizeRequest", string(data)}, " ")
}
