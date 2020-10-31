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
type ListRuntimesRequest struct {
}

func (o ListRuntimesRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListRuntimesRequest", string(data)}, " ")
}
