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
type ListBranchesResponse struct {
	// 项目分支列表。
	Branches *[]string `json:"branches,omitempty"`
}

func (o ListBranchesResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListBranchesResponse", string(data)}, " ")
}
