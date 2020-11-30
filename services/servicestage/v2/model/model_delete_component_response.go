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
type DeleteComponentResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteComponentResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeleteComponentResponse", string(data)}, " ")
}
