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
type ListBranchesRequest struct {
	XRepoAuth string `json:"X-Repo-Auth"`
	Namespace string `json:"namespace"`
	Project   string `json:"project"`
}

func (o ListBranchesRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListBranchesRequest", string(data)}, " ")
}
