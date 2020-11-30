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
type DeleteFileResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteFileResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeleteFileResponse", string(data)}, " ")
}
