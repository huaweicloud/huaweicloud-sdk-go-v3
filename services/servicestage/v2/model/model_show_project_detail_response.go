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
type ShowProjectDetailResponse struct {
	// 命名空间ID。
	NamespaceId *string `json:"namespace_id,omitempty"`
	// 命名空间。
	Namespace *string `json:"namespace,omitempty"`
	// 仓库项目ID。
	ProjectId *string `json:"project_id,omitempty"`
	// 仓库项目。
	Project *string `json:"project,omitempty"`
}

func (o ShowProjectDetailResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowProjectDetailResponse", string(data)}, " ")
}
