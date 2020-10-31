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
type ListNamespacesRequest struct {
	XRepoAuth string `json:"X-Repo-Auth"`
}

func (o ListNamespacesRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListNamespacesRequest", string(data)}, " ")
}
