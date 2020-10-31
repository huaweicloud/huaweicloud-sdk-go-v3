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
type ListTemplatesRequest struct {
}

func (o ListTemplatesRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListTemplatesRequest", string(data)}, " ")
}
