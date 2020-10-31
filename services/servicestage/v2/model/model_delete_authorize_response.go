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
type DeleteAuthorizeResponse struct {
}

func (o DeleteAuthorizeResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeleteAuthorizeResponse", string(data)}, " ")
}
