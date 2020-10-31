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
type ListProjectsResponse struct {
	// 项目列表。
	Projects *[]Project `json:"projects,omitempty"`
}

func (o ListProjectsResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListProjectsResponse", string(data)}, " ")
}
