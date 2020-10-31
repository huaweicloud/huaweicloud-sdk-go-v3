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
type DeleteEnvironmentResponse struct {
}

func (o DeleteEnvironmentResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeleteEnvironmentResponse", string(data)}, " ")
}
